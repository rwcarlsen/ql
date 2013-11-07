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

func (s *stmt) prepareCtx() error {
	switch s.lst[0].String() {
	case "BEGIN TRANSACTION;","COMMIT;","ROLLBACK;":
		if s.cn.currCtx != nil {
			return errors.New("ql: cannot nest transactions")
		}
		s.cn.currCtx = NewRWCtx()
	}
	return nil
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	if err := s.prepareCtx(); err != nil {
		return nil, err
	}

	if err := s.cn.Execute(vtoi(args)...); err != nil {
		return nil, err
	}
	return result{s.Stmt.Conn()}, nil
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	if err := s.Stmt.Query(vtoi(args)...); err != nil && err != io.EOF {
		return nil, err
	}
	return &rows{s, true}, nil
}

// result implements driver.Result.
type result struct {
	*Conn
}

func (r result) LastInsertId() (int64, error) {
	return int64(r.Conn.LastInsertId()), nil
}

func (r result) RowsAffected() (int64, error) {
	return int64(r.Conn.RowsAffected()), nil
}

// rows implements driver.Rows.
type rows struct {
	*stmt
	first bool
}

func (r *rows) Close() error {
	if r.stmt.closed {
		return r.stmt.Stmt.Close()
	}
	r.stmt.Stmt.Reset()
	return nil
}

func (r *rows) Next(dest []driver.Value) error {
	if r.first {
		r.first = false
		if !r.stmt.Stmt.Busy() {
			return io.EOF
		}
	} else if err := r.stmt.Stmt.Next(); err != nil {
		return err
	}
	for i := range dest {
		v := (*interface{})(&dest[i])
		err := r.stmt.Stmt.scanDynamic(C.int(i), v, true)
		if err != nil {
			return err
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
