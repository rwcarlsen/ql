// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ql

import (
	"fmt"
	"io"
	"log"
	"math"
)

// QL types.
const (
	qBool       = 0x62 // 'b'
	qComplex64  = 0x63 // 'c'
	qComplex128 = 0x64 // 'd'
	qFloat32    = 0x66 // 'f'
	qFloat64    = 0x67 // 'g', alias float
	qInt8       = 0x69 // 'i'
	qInt16      = 0x6a // 'j'
	qInt32      = 0x6b // 'k'
	qInt64      = 0x6c // 'l', alias int
	qString     = 0x73 // 's'
	qUint8      = 0x75 // 'u', alias byte
	qUint16     = 0x76 // 'v'
	qUint32     = 0x77 // 'w'
	qUint64     = 0x78 // 'x', alias uint

	//LATER Gob = 'o'
)

var (
	type2Str = map[int]string{
		qBool:       "bool",
		qComplex64:  "complex64",
		qComplex128: "complex128",
		qFloat32:    "float32",
		qFloat64:    "float64",
		qInt8:       "int8",
		qInt16:      "int16",
		qInt32:      "int32",
		qInt64:      "int64",
		qString:     "string",
		qUint8:      "uint8",
		qUint16:     "uint16",
		qUint32:     "uint32",
		qUint64:     "uint64",
	}
)

func typeStr(typ int) (r string) {
	return type2Str[typ]
}

func noEOF(err error) error {
	if err == io.EOF {
		err = nil
	}
	return err
}

func runErr(err error) error { return fmt.Errorf("run time error: %s", err) }

func invXOp(s, x interface{}) error {
	return fmt.Errorf("invalid operation: %v[%v] (index of type %T)", s, x, x)
}

func invSOp(s interface{}) error {
	return fmt.Errorf("cannot slice %s (type %T)", s, s)
}

func invNegX(x interface{}) error {
	return fmt.Errorf("invalid string index %v (index must be non-negative)", x)
}

func invSliceNegX(x interface{}) error {
	return fmt.Errorf("invalid slice index %v (index must be non-negative)", x)
}

func invBoundX(s string, x uint64) error {
	return fmt.Errorf("invalid string index %d (out of bounds for %d-byte string)", x, len(s))
}

func invSliceBoundX(s string, x uint64) error {
	return fmt.Errorf("invalid slice index %d (out of bounds for %d-byte string)", x, len(s))
}

func indexExpr(s *string, x interface{}) (i uint64, err error) {
	switch x := x.(type) {
	case idealFloat:
		if x < 0 {
			return 0, invNegX(x)
		}

		if s != nil && int(x) >= len(*s) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case idealInt:
		if x < 0 {
			return 0, invNegX(x)
		}

		if s != nil && int64(x) >= int64(len(*s)) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case idealRune:
		if x < 0 {
			return 0, invNegX(x)
		}

		if s != nil && int32(x) >= int32(len(*s)) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case idealUint:
		if x < 0 {
			return 0, invNegX(x)
		}

		if s != nil && uint64(x) >= uint64(len(*s)) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case int8:
		if x < 0 {
			return 0, invNegX(x)
		}

		if s != nil && int(x) >= len(*s) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case int16:
		if x < 0 {
			return 0, invNegX(x)
		}

		if s != nil && int(x) >= len(*s) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case int32:
		if x < 0 {
			return 0, invNegX(x)
		}

		if s != nil && int(x) >= len(*s) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case int64:
		if x < 0 {
			return 0, invNegX(x)
		}

		if s != nil && x >= int64(len(*s)) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case uint8:
		if s != nil && int(x) >= len(*s) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case uint16:
		if s != nil && int(x) >= len(*s) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case uint32:
		if s != nil && x >= uint32(len(*s)) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case uint64:
		if s != nil && x >= uint64(len(*s)) {
			return 0, invBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	default:
		return 0, fmt.Errorf("non-integer string index %v", x)
	}
}

func sliceExpr(s *string, x interface{}, mod int) (i uint64, err error) {
	switch x := x.(type) {
	case idealFloat:
		if x < 0 {
			return 0, invSliceNegX(x)
		}

		if s != nil && int(x) >= len(*s)+mod {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case idealInt:
		if x < 0 {
			return 0, invSliceNegX(x)
		}

		if s != nil && int64(x) >= int64(len(*s)+mod) {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case idealRune:
		if x < 0 {
			return 0, invSliceNegX(x)
		}

		if s != nil && int32(x) >= int32(len(*s)+mod) {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case idealUint:
		if x < 0 {
			return 0, invSliceNegX(x)
		}

		if s != nil && uint64(x) >= uint64(len(*s)+mod) {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case int8:
		if x < 0 {
			return 0, invSliceNegX(x)
		}

		if s != nil && int(x) >= len(*s)+mod {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case int16:
		if x < 0 {
			return 0, invSliceNegX(x)
		}

		if s != nil && int(x) >= len(*s)+mod {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case int32:
		if x < 0 {
			return 0, invSliceNegX(x)
		}

		if s != nil && int(x) >= len(*s)+mod {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case int64:
		if x < 0 {
			return 0, invSliceNegX(x)
		}

		if s != nil && x >= int64(len(*s)+mod) {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case uint8:
		if s != nil && int(x) >= len(*s)+mod {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case uint16:
		if s != nil && int(x) >= len(*s)+mod {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case uint32:
		if s != nil && x >= uint32(len(*s)+mod) {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	case uint64:
		if s != nil && x >= uint64(len(*s)+mod) {
			return 0, invSliceBoundX(*s, uint64(x))
		}

		return uint64(x), nil
	default:
		return 0, fmt.Errorf("invalid slice index %s (type %T)", x, x)
	}
}

type iop int

func (o iop) String() string {
	switch i := int(o); i {
	case andand:
		return "&&"
	case andnot:
		return "&^"
	case lsh:
		return "<<"
	case le:
		return "<="
	case eq:
		return "=="
	case ge:
		return ">="
	case neq:
		return "!="
	case oror:
		return "||"
	case rsh:
		return ">>"
	default:
		return string(i)
	}
}

func ideal(v interface{}) interface{} {
	switch x := v.(type) {
	case idealComplex:
		return complex128(x)
	case idealFloat:
		return float64(x)
	case idealInt:
		return int64(x)
	case idealRune:
		return int64(x)
	case idealUint:
		return uint64(x)
	default:
		return v
	}
}

func eval(v expression, ctx map[interface{}]interface{}, arg []interface{}) (y interface{}) {
	y, err := v.eval(ctx, arg)
	if err != nil {
		panic(err) // panic ok here
	}
	return
}

func eval2(a, b expression, ctx map[interface{}]interface{}, arg []interface{}) (x, y interface{}) {
	return eval(a, ctx, arg), eval(b, ctx, arg)
}

func invOp2(x, y interface{}, o int) (interface{}, error) {
	return nil, fmt.Errorf("invalid operation: %v %v %v (mismatched types %T and %T)", x, iop(o), y, ideal(x), ideal(y))
}

func undOp(x interface{}, o int) (interface{}, error) {
	return nil, fmt.Errorf("invalid operation: %v%v (operator %v not defined on %T)", iop(o), x, iop(o), x)
}

func undOp2(x, y interface{}, o int) (interface{}, error) {
	return nil, fmt.Errorf("invalid operation: %v %v %v (operator %v not defined on %T)", x, iop(o), y, iop(o), x)
}

func invConv(val interface{}, typ int) (interface{}, error) {
	return nil, fmt.Errorf("cannot convert %v (type %T) to type %s", val, val, typeStr(typ))
}

func truncConv(val interface{}) (interface{}, error) {
	return nil, fmt.Errorf("constant %v truncated to integer", val)
}

func convert(val interface{}, typ int) (v interface{}, err error) {
	if val == nil {
		return nil, nil
	}

	switch typ {
	case qBool:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		//case idealFloat:
		//case idealInt:
		//case idealRune:
		//case idealUint:
		case bool:
			return bool(x), nil
		//case complex64:
		//case complex128:
		//case float32:
		//case float64:
		//case int8:
		//case int16:
		//case int32:
		//case int64:
		//case string:
		//case uint8:
		//case uint16:
		//case uint32:
		//case uint64:
		default:
			return invConv(val, typ)
		}
	case qComplex64:
		switch x := val.(type) {
		//case nil:
		case idealComplex:
			return complex64(x), nil
		case idealFloat:
			return complex(float32(x), 0), nil
		case idealInt:
			return complex(float32(x), 0), nil
		case idealRune:
			return complex(float32(x), 0), nil
		case idealUint:
			return complex(float32(x), 0), nil
		//case bool:
		case complex64:
			return complex64(x), nil
		case complex128:
			return complex64(x), nil
		//case float32:
		//case float64:
		//case int8:
		//case int16:
		//case int32:
		//case int64:
		//case string:
		//case uint8:
		//case uint16:
		//case uint32:
		//case uint64:
		default:
			return invConv(val, typ)
		}
	case qComplex128:
		switch x := val.(type) {
		//case nil:
		case idealComplex:
			return complex128(x), nil
		case idealFloat:
			return complex(float64(x), 0), nil
		case idealInt:
			return complex(float64(x), 0), nil
		case idealRune:
			return complex(float64(x), 0), nil
		case idealUint:
			return complex(float64(x), 0), nil
		//case bool:
		case complex64:
			return complex128(x), nil
		case complex128:
			return complex128(x), nil
		//case float32:
		//case float64:
		//case int8:
		//case int16:
		//case int32:
		//case int64:
		//case string:
		//case uint8:
		//case uint16:
		//case uint32:
		//case uint64:
		default:
			return invConv(val, typ)
		}
	case qFloat32:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			return float32(x), nil
		case idealInt:
			return float32(x), nil
		case idealRune:
			return float32(x), nil
		case idealUint:
			return float32(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return float32(x), nil
		case float64:
			return float32(x), nil
		case int8:
			return float32(x), nil
		case int16:
			return float32(x), nil
		case int32:
			return float32(x), nil
		case int64:
			return float32(x), nil
		//case string:
		case uint8:
			return float32(x), nil
		case uint16:
			return float32(x), nil
		case uint32:
			return float32(x), nil
		case uint64:
			return float32(x), nil
		default:
			return invConv(val, typ)
		}
	case qFloat64:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			return float64(x), nil
		case idealInt:
			return float64(x), nil
		case idealRune:
			return float64(x), nil
		case idealUint:
			return float64(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return float64(x), nil
		case float64:
			return float64(x), nil
		case int8:
			return float64(x), nil
		case int16:
			return float64(x), nil
		case int32:
			return float64(x), nil
		case int64:
			return float64(x), nil
		//case string:
		case uint8:
			return float64(x), nil
		case uint16:
			return float64(x), nil
		case uint32:
			return float64(x), nil
		case uint64:
			return float64(x), nil
		default:
			return invConv(val, typ)
		}
	case qInt8:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			if _, frac := math.Modf(float64(x)); frac != 0 {
				return truncConv(x)
			}

			return int8(x), nil
		case idealInt:
			return int8(x), nil
		case idealRune:
			return int8(x), nil
		case idealUint:
			return int8(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return int8(x), nil
		case float64:
			return int8(x), nil
		case int8:
			return int8(x), nil
		case int16:
			return int8(x), nil
		case int32:
			return int8(x), nil
		case int64:
			return int8(x), nil
		//case string:
		case uint8:
			return int8(x), nil
		case uint16:
			return int8(x), nil
		case uint32:
			return int8(x), nil
		case uint64:
			return int8(x), nil
		default:
			return invConv(val, typ)
		}
	case qInt16:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			if _, frac := math.Modf(float64(x)); frac != 0 {
				return truncConv(x)
			}

			return int16(x), nil
		case idealInt:
			return int16(x), nil
		case idealRune:
			return int16(x), nil
		case idealUint:
			return int16(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return int16(x), nil
		case float64:
			return int16(x), nil
		case int8:
			return int16(x), nil
		case int16:
			return int16(x), nil
		case int32:
			return int16(x), nil
		case int64:
			return int16(x), nil
		//case string:
		case uint8:
			return int16(x), nil
		case uint16:
			return int16(x), nil
		case uint32:
			return int16(x), nil
		case uint64:
			return int16(x), nil
		default:
			return invConv(val, typ)
		}
	case qInt32:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			if _, frac := math.Modf(float64(x)); frac != 0 {
				return truncConv(x)
			}

			return int32(x), nil
		case idealInt:
			return int32(x), nil
		case idealRune:
			return int32(x), nil
		case idealUint:
			return int32(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return int32(x), nil
		case float64:
			return int32(x), nil
		case int8:
			return int32(x), nil
		case int16:
			return int32(x), nil
		case int32:
			return int32(x), nil
		case int64:
			return int32(x), nil
		//case string:
		case uint8:
			return int32(x), nil
		case uint16:
			return int32(x), nil
		case uint32:
			return int32(x), nil
		case uint64:
			return int32(x), nil
		default:
			return invConv(val, typ)
		}
	case qInt64:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			if _, frac := math.Modf(float64(x)); frac != 0 {
				return truncConv(x)
			}

			return int64(x), nil
		case idealInt:
			return int64(x), nil
		case idealRune:
			return int64(x), nil
		case idealUint:
			return int64(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return int64(x), nil
		case float64:
			return int64(x), nil
		case int8:
			return int64(x), nil
		case int16:
			return int64(x), nil
		case int32:
			return int64(x), nil
		case int64:
			return int64(x), nil
		//case string:
		case uint8:
			return int64(x), nil
		case uint16:
			return int64(x), nil
		case uint32:
			return int64(x), nil
		case uint64:
			return int64(x), nil
		default:
			return invConv(val, typ)
		}
	case qString:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		//case idealFloat:
		case idealInt:
			return string(x), nil
		case idealRune:
			return string(x), nil
		case idealUint:
			return string(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		//case float32:
		//case float64:
		case int8:
			return string(x), nil
		case int16:
			return string(x), nil
		case int32:
			return string(x), nil
		case int64:
			return string(x), nil
		case string:
			return string(x), nil
		case uint8:
			return string(x), nil
		case uint16:
			return string(x), nil
		case uint32:
			return string(x), nil
		case uint64:
			return string(x), nil
		default:
			return invConv(val, typ)
		}
	case qUint8:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			if _, frac := math.Modf(float64(x)); frac != 0 {
				return truncConv(x)
			}

			return uint8(x), nil
		case idealInt:
			return uint8(x), nil
		case idealRune:
			return uint8(x), nil
		case idealUint:
			return uint8(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return uint8(x), nil
		case float64:
			return uint8(x), nil
		case int8:
			return uint8(x), nil
		case int16:
			return uint8(x), nil
		case int32:
			return uint8(x), nil
		case int64:
			return uint8(x), nil
		//case string:
		case uint8:
			return uint8(x), nil
		case uint16:
			return uint8(x), nil
		case uint32:
			return uint8(x), nil
		case uint64:
			return uint8(x), nil
		default:
			return invConv(val, typ)
		}
	case qUint16:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			if _, frac := math.Modf(float64(x)); frac != 0 {
				return truncConv(x)
			}

			return uint16(x), nil
		case idealInt:
			return uint16(x), nil
		case idealRune:
			return uint16(x), nil
		case idealUint:
			return uint16(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return uint16(x), nil
		case float64:
			return uint16(x), nil
		case int8:
			return uint16(x), nil
		case int16:
			return uint16(x), nil
		case int32:
			return uint16(x), nil
		case int64:
			return uint16(x), nil
		//case string:
		case uint8:
			return uint16(x), nil
		case uint16:
			return uint16(x), nil
		case uint32:
			return uint16(x), nil
		case uint64:
			return uint16(x), nil
		default:
			return invConv(val, typ)
		}
	case qUint32:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			if _, frac := math.Modf(float64(x)); frac != 0 {
				return truncConv(x)
			}

			return uint32(x), nil
		case idealInt:
			return uint32(x), nil
		case idealRune:
			return uint32(x), nil
		case idealUint:
			return uint32(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return uint32(x), nil
		case float64:
			return uint32(x), nil
		case int8:
			return uint32(x), nil
		case int16:
			return uint32(x), nil
		case int32:
			return uint32(x), nil
		case int64:
			return uint32(x), nil
		//case string:
		case uint8:
			return uint32(x), nil
		case uint16:
			return uint32(x), nil
		case uint32:
			return uint32(x), nil
		case uint64:
			return uint32(x), nil
		default:
			return invConv(val, typ)
		}
	case qUint64:
		switch x := val.(type) {
		//case nil:
		//case idealComplex:
		case idealFloat:
			if _, frac := math.Modf(float64(x)); frac != 0 {
				return truncConv(x)
			}

			return uint64(x), nil
		case idealInt:
			return uint64(x), nil
		case idealRune:
			return uint64(x), nil
		case idealUint:
			return uint64(x), nil
		//case bool:
		//case complex64:
		//case complex128:
		case float32:
			return uint64(x), nil
		case float64:
			return uint64(x), nil
		case int8:
			return uint64(x), nil
		case int16:
			return uint64(x), nil
		case int32:
			return uint64(x), nil
		case int64:
			return uint64(x), nil
		//case string:
		case uint8:
			return uint64(x), nil
		case uint16:
			return uint64(x), nil
		case uint32:
			return uint64(x), nil
		case uint64:
			return uint64(x), nil
		default:
			return invConv(val, typ)
		}
	default:
		log.Panic("internal error")
	}
	panic("unreachable")
}

func invShiftRHS(lhs, rhs interface{}) (interface{}, error) {
	return nil, fmt.Errorf("invalid operation: %v << %v (shift count type %T, must be unsigned integer)", lhs, rhs, rhs)
}

func typeCheck(rec []interface{}, cols []*col) (err error) {
	for _, c := range cols {
		i := c.index
		if v := rec[i]; !c.typeCheck(v) {
			switch v.(type) {
			case idealComplex:
				y := complex128(v.(idealComplex))
				switch c.typ {
				case qBool:
				case qComplex64:
					rec[i] = complex64(y)
					continue
				case qComplex128:
					rec[i] = complex128(y)
					continue
				case qFloat32, qFloat64, qInt8, qInt16, qInt32, qInt64, qUint8, qUint16, qUint32, qUint64:
					return fmt.Errorf("constant %v truncated to real", y)
				case qString:
				default:
					log.Panic("internal error")
				}
			case idealFloat:
				y := float64(v.(idealFloat))
				switch c.typ {
				case qBool:
				case qComplex64:
					rec[i] = complex(float32(y), 0)
					continue
				case qComplex128:
					rec[i] = complex(float64(y), 0)
					continue
				case qFloat32:
					rec[i] = float32(y)
					continue
				case qFloat64:
					rec[i] = float64(y)
					continue
				case qInt8:
					rec[i] = int8(y)
					continue
				case qInt16:
					rec[i] = int16(y)
					continue
				case qInt32:
					rec[i] = int32(y)
					continue
				case qInt64:
					rec[i] = int64(y)
					continue
				case qString:
				case qUint8:
					rec[i] = uint8(y)
					continue
				case qUint16:
					rec[i] = uint16(y)
					continue
				case qUint32:
					rec[i] = uint32(y)
					continue
				case qUint64:
					rec[i] = uint64(y)
					continue
				default:
					log.Panic("internal error")
				}
			case idealInt:
				y := int64(v.(idealInt))
				switch c.typ {
				case qBool:
				case qComplex64:
					rec[i] = complex(float32(y), 0)
					continue
				case qComplex128:
					rec[i] = complex(float64(y), 0)
					continue
				case qFloat32:
					rec[i] = float32(y)
					continue
				case qFloat64:
					rec[i] = float64(y)
					continue
				case qInt8:
					rec[i] = int8(y)
					continue
				case qInt16:
					rec[i] = int16(y)
					continue
				case qInt32:
					rec[i] = int32(y)
					continue
				case qInt64:
					rec[i] = int64(y)
					continue
				case qString:
				case qUint8:
					rec[i] = uint8(y)
					continue
				case qUint16:
					rec[i] = uint16(y)
					continue
				case qUint32:
					rec[i] = uint32(y)
					continue
				case qUint64:
					rec[i] = uint64(y)
					continue
				default:
					log.Panic("internal error")
				}
			case idealRune:
				y := int64(v.(idealRune))
				switch c.typ {
				case qBool:
				case qComplex64:
					rec[i] = complex(float32(y), 0)
					continue
				case qComplex128:
					rec[i] = complex(float64(y), 0)
					continue
				case qFloat32:
					rec[i] = float32(y)
					continue
				case qFloat64:
					rec[i] = float64(y)
					continue
				case qInt8:
					rec[i] = int8(y)
					continue
				case qInt16:
					rec[i] = int16(y)
					continue
				case qInt32:
					rec[i] = int32(y)
					continue
				case qInt64:
					rec[i] = int64(y)
					continue
				case qString:
				case qUint8:
					rec[i] = uint8(y)
					continue
				case qUint16:
					rec[i] = uint16(y)
					continue
				case qUint32:
					rec[i] = uint32(y)
					continue
				case qUint64:
					rec[i] = uint64(y)
					continue
				default:
					log.Panic("internal error")
				}
			case idealUint:
				y := uint64(v.(idealUint))
				switch c.typ {
				case qBool:
				case qComplex64:
					rec[i] = complex(float32(y), 0)
					continue
				case qComplex128:
					rec[i] = complex(float64(y), 0)
					continue
				case qFloat32:
					rec[i] = float32(y)
					continue
				case qFloat64:
					rec[i] = float64(y)
					continue
				case qInt8:
					rec[i] = int8(y)
					continue
				case qInt16:
					rec[i] = int16(y)
					continue
				case qInt32:
					rec[i] = int32(y)
					continue
				case qInt64:
					rec[i] = int64(y)
					continue
				case qString:
					rec[i] = string(y)
					continue
				case qUint8:
					rec[i] = uint8(y)
					continue
				case qUint16:
					rec[i] = uint16(y)
					continue
				case qUint32:
					rec[i] = uint32(y)
					continue
				case qUint64:
					rec[i] = uint64(y)
					continue
				default:
					log.Panic("internal error")
				}
			}
			return fmt.Errorf("cannot use %v (type %T) as %s in assignment to column %s", v, ideal(v), typeStr(c.typ), c.name)
		}
	}
	return
}
