package ql

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"reflect"
	"time"
	"unsafe"
)

// Driver implements the interface required by database/sql.
type Driver string

func register(name string) {
	defer func() { recover() }()
	sql.Register(name, Driver(name))
}

func (Driver) Open(name string) (driver.Conn, error) {
	db, err := OpenFile(name, &Options{CanCreate: true})
	if err != nil {
		return nil, err
	}
	return &conn{db}, nil
}

// conn implements driver.Conn.
type conn struct {
	*DB
	currCtx *tctx
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	lst, err := Compile(query)
	if err != nil {
		return nil, err
	}
	return &stmt{c, lst}
}

func (c *conn) Begin() (driver.Tx, error) {
	if c.currCtx != nil {
		return nil, errors.New("ql: cannot nest transactions")
	}

	c.currCtx = NewRWCtx()

	rs, i, err := c.Run(c.currCtx, "BEGIN TRANSACTION;")
	if err != nil {
		return nil, fmt.Errorf("ql: statement %v: %v", i, err)
	}
	return &tx{cn}
	return c.Conn, nil
}

type tx struct {
	cn *conn
}

func (t *tx) Commit() error {
	_, _, err := t.cn.Run(c.currCtx, "COMMIT;")
	t.cn.currCtx = nil
	return err
}

func (t *tx) Rollback() error {
	_, _, err := t.cn.Run(c.currCtx, "ROLLBACK;")
	t.cn.currCtx = nil
	return err
}

// stmt implements driver.Stmt.
type stmt struct {
	cn *conn
	lst list
}

func (s *stmt) Close() error {
	return nil
}

func (s *stmt) NumInput() int {
	return -1
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	rs, i, err := s.run(args)
	if err != nil {
		return nil, fmt.Errorf("ql: statement %v: %v", i, err)
	}
	return result{}, nil
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	rs, i, err := s.run(args)
	if err != nil {
		return nil, fmt.Errorf("ql: statement %v: %v", i, err)
	}

	return &rows{s, true}, nil
}

func (s *stmt) run(args []driver.Value) (Recordset, int, error) {
	switch s.lst[0].String() {
	case beginTransactionStmt{}.String(),commitStmt{}.String(),rollbackStmt{}.String():
		if s.cn.currCtx != nil {
			return nil, 0, errors.New("ql: cannot nest transactions")
		}
		s.cn.currCtx = NewRWCtx()
	}
	return s.cn.Execute(s.cn.currCtx, s.lst, vtoi(args)...)
}

// result implements driver.Result.
type result struct { }

func (result) LastInsertId() (int64, error) {
	return 0, nil
}

func (result) RowsAffected() (int64, error) {
	return 0, nil
}

// rows implements driver.Rows.
type rows struct {
	rs Recordset
	first bool
	cols []string
	ch chan []interface{}
	done chan bool
	closed bool
}

func (r *rows) Close() error {
	if !r.closed {
		r.closed = true
		if r.done != nil {
			r.done <- true
			close(r.done)
			close(r.ch)
		}
	}
	return nil
}

func (r *rows) Columns() []string {
	r.init()
	return r.cols
}

func (r *rows) init() {
	if !r.first {
		return
	}

	r.first = false
	r.ch = make(chan []interface{})
	r.done = make(chan bool, 2)
	fn := func(data []interface) (more bool, err error) {
		select {
		case r.ch<-data:
		case <-r.done:
			return false, nil
		default:
			return true, nil
		}
	}

	go func() {
		r.rs.Do(true, fn)
		r.ch <- []interface{}{}
	}()

	cols := <-r.ch
	for _, c := range cols {
		r.cols = append(r.cols, c.(string))
	}
}

func (r *rows) Next(dest []driver.Value) error {
	r.init()
	if r.empty {
		return io.EOF
	}

	row := <-r.ch
	if len(row) == 0 {
		r.empty = true
		return io.EOF
	}

	for i := range dest {
		switch v := row[i].(type) {
		case byte:
			dest[i] = int64(v)
		case complex128:
		// unsupported (panic?)
			dest[i] = 0
		case complex64:
		// unsupported (panic?)
			dest[i] = 0
		case float32:
			dest[i] = float64(v)
		case float64:
			dest[i] = v
		case int:
			dest[i] = int64(v)
		case int16:
			dest[i] = int64(v)
		case int32:
			dest[i] = int64(v)
		case int64:
			dest[i] = v
		case int8:
			dest[i] = int64(v)
		case rune:
			dest[i] = int64(v)
		case string:
			dest[i] = v
		case uint:
			dest[i] = int64(v)
		case uint16:
			dest[i] = int64(v)
		case uint32:
			dest[i] = int64(v)
		case uint64:
			dest[i] = int64(v)
		case uint8:
			dest[i] = int64(v)
		default:
			panic(fmt.Sprintf("ql: unsupported data type %T", row[i])
			dest[i] = row[i]
		}
	}
	return nil
}

// vtoi converts []driver.Value to []interface{} without copying the contents.
func vtoi(v []driver.Value) (i []interface{}) {
	if len(v) > 0 {
		h := (*reflect.SliceHeader)(unsafe.Pointer(&i))
		h.Data = uintptr(unsafe.Pointer(&v[0]))
		h.Len = len(v)
		h.Cap = cap(v)
	}
	return
}
