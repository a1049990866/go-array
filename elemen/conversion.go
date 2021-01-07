package elemen

import (
	"errors"
	"strconv"
)

type Element struct {
	Val   interface{}
	Type  Type
	Error error
}

func NewElement(val interface{}, typ ...Type) *Element {
	if len(typ) == 0 {
		typ = append(typ, InterfaceElement)
	}
	return &Element{
		Val:  val,
		Type: typ[0],
	}
}

func (t *Element) ToInt() (res int) {
	switch t.Type {
	case IntElement:
		return t.Val.(int)
	case Int32Element:
		return int(t.Val.(int32))
	case Int64Element:
		return int(t.Val.(int64))
	case UintElement:
		return int(t.Val.(uint))
	case Uint32Element:
		return int(t.Val.(uint32))
	case Uint64Element:
		return int(t.Val.(uint64))
	case Float32Element:
		return int(t.Val.(float32))
	case Float64Element:
		return int(t.Val.(float64))
	case StringElement:
		res, t.Error = strconv.Atoi(t.Val.(string))
	case InterfaceElement:
		return t.interfaceToInt()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToInt() (res int) {
	switch t.Val.(type) {
	case int:
		return t.Val.(int)
	case int32:
		return int(t.Val.(int32))
	case int64:
		return int(t.Val.(int64))
	case uint:
		return int(t.Val.(uint))
	case uint32:
		return int(t.Val.(uint32))
	case uint64:
		return int(t.Val.(uint64))
	case float32:
		return int(t.Val.(float32))
	case float64:
		return int(t.Val.(float64))
	case string:
		res, t.Error = strconv.Atoi(t.Val.(string))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToInt32() (res int32) {
	switch t.Type {
	case IntElement:
		return int32(t.Val.(int))
	case Int32Element:
		return t.Val.(int32)
	case Int64Element:
		return int32(t.Val.(int64))
	case UintElement:
		return int32(t.Val.(uint))
	case Uint32Element:
		return int32(t.Val.(uint32))
	case Uint64Element:
		return int32(t.Val.(uint64))
	case Float32Element:
		return int32(t.Val.(float32))
	case Float64Element:
		return int32(t.Val.(float64))
	case StringElement:
		t.Val, t.Error = strconv.ParseInt(t.Val.(string), 10, 32)
		return int32(t.Val.(int64))
	case InterfaceElement:
		return t.interfaceToInt32()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToInt32() (res int32) {
	switch t.Val.(type) {
	case int:
		return int32(t.Val.(int))
	case int32:
		return t.Val.(int32)
	case int64:
		return int32(t.Val.(int64))
	case uint:
		return int32(t.Val.(uint))
	case uint32:
		return int32(t.Val.(uint32))
	case uint64:
		return int32(t.Val.(uint64))
	case float32:
		return int32(t.Val.(float32))
	case float64:
		return int32(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseInt(t.Val.(string), 10, 32)
		return int32(t.Val.(int64))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToInt64() (res int64) {
	switch t.Type {
	case IntElement:
		return int64(t.Val.(int))
	case Int32Element:
		return int64(t.Val.(int32))
	case Int64Element:
		return t.Val.(int64)
	case UintElement:
		return int64(t.Val.(uint))
	case Uint32Element:
		return int64(t.Val.(uint32))
	case Uint64Element:
		return int64(t.Val.(uint64))
	case Float32Element:
		return int64(t.Val.(float32))
	case Float64Element:
		return int64(t.Val.(float64))
	case StringElement:
		t.Val, t.Error = strconv.ParseInt(t.Val.(string), 10, 64)
		return t.Val.(int64)
	case InterfaceElement:
		return t.interfaceToInt64()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToInt64() (res int64) {
	switch t.Val.(type) {
	case int:
		return int64(t.Val.(int))
	case int32:
		return int64(t.Val.(int32))
	case int64:
		return t.Val.(int64)
	case uint:
		return int64(t.Val.(uint))
	case uint32:
		return int64(t.Val.(uint32))
	case uint64:
		return int64(t.Val.(uint64))
	case float32:
		return int64(t.Val.(float32))
	case float64:
		return int64(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseInt(t.Val.(string), 10, 64)
		return t.Val.(int64)
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToUint() (res uint) {
	switch t.Type {
	case IntElement:
		return uint(t.Val.(int))
	case Int32Element:
		return uint(t.Val.(int32))
	case Int64Element:
		return uint(t.Val.(int64))
	case UintElement:
		return t.Val.(uint)
	case Uint32Element:
		return uint(t.Val.(uint32))
	case Uint64Element:
		return uint(t.Val.(uint64))
	case Float32Element:
		return uint(t.Val.(float32))
	case Float64Element:
		return uint(t.Val.(float64))
	case StringElement:
		t.Val, t.Error = strconv.ParseUint(t.Val.(string), 10, 32)
		return uint(t.Val.(uint64))
	case InterfaceElement:
		return t.interfaceToUint()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToUint() (res uint) {
	switch t.Val.(type) {
	case int:
		return uint(t.Val.(int))
	case int32:
		return uint(t.Val.(int32))
	case int64:
		return uint(t.Val.(int64))
	case uint:
		return t.Val.(uint)
	case uint32:
		return uint(t.Val.(uint32))
	case uint64:
		return uint(t.Val.(uint64))
	case float32:
		return uint(t.Val.(float32))
	case float64:
		return uint(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseUint(t.Val.(string), 10, 32)
		return uint(t.Val.(uint64))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToUint32() (res uint32) {
	switch t.Type {
	case IntElement:
		return uint32(t.Val.(int))
	case Int32Element:
		return uint32(t.Val.(int32))
	case Int64Element:
		return uint32(t.Val.(int64))
	case UintElement:
		return uint32(t.Val.(uint))
	case Uint32Element:
		return t.Val.(uint32)
	case Uint64Element:
		return uint32(t.Val.(uint64))
	case Float32Element:
		return uint32(t.Val.(float32))
	case Float64Element:
		return uint32(t.Val.(float64))
	case StringElement:
		t.Val, t.Error = strconv.ParseUint(t.Val.(string), 10, 32)
		return uint32(t.Val.(uint64))
	case InterfaceElement:
		return t.interfaceToUint32()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToUint32() (res uint32) {
	switch t.Val.(type) {
	case int:
		return uint32(t.Val.(int))
	case int32:
		return uint32(t.Val.(int32))
	case int64:
		return uint32(t.Val.(int64))
	case uint:
		return uint32(t.Val.(uint))
	case uint32:
		return t.Val.(uint32)
	case uint64:
		return uint32(t.Val.(uint64))
	case float32:
		return uint32(t.Val.(float32))
	case float64:
		return uint32(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseUint(t.Val.(string), 10, 32)
		return uint32(t.Val.(uint64))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToUint64() (res uint64) {
	switch t.Type {
	case IntElement:
		return uint64(t.Val.(int))
	case Int32Element:
		return uint64(t.Val.(int32))
	case Int64Element:
		return uint64(t.Val.(int64))
	case UintElement:
		return uint64(t.Val.(uint))
	case Uint32Element:
		return uint64(t.Val.(uint32))
	case Uint64Element:
		return t.Val.(uint64)
	case Float32Element:
		return uint64(t.Val.(float32))
	case Float64Element:
		return uint64(t.Val.(float64))
	case StringElement:
		res, t.Error = strconv.ParseUint(t.Val.(string), 10, 64)
	case InterfaceElement:
		return t.interfaceToUint64()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToUint64() (res uint64) {
	switch t.Val.(type) {
	case int:
		return uint64(t.Val.(int))
	case int32:
		return uint64(t.Val.(int32))
	case int64:
		return uint64(t.Val.(int64))
	case uint:
		return uint64(t.Val.(uint))
	case uint32:
		return uint64(t.Val.(uint32))
	case uint64:
		return t.Val.(uint64)
	case float32:
		return uint64(t.Val.(float32))
	case float64:
		return uint64(t.Val.(float64))
	case string:
		res, t.Error = strconv.ParseUint(t.Val.(string), 10, 64)
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToFloat32() (res float32) {
	switch t.Type {
	case IntElement:
		return float32(t.Val.(int))
	case Int32Element:
		return float32(t.Val.(int32))
	case Int64Element:
		return float32(t.Val.(int64))
	case UintElement:
		return float32(t.Val.(uint))
	case Uint32Element:
		return float32(t.Val.(uint32))
	case Uint64Element:
		return float32(t.Val.(uint64))
	case Float32Element:
		return t.Val.(float32)
	case Float64Element:
		return float32(t.Val.(float64))
	case StringElement:
		t.Val, t.Error = strconv.ParseFloat(t.Val.(string), 32)
		return float32(t.Val.(float64))
	case InterfaceElement:
		return t.interfaceToFloat32()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToFloat32() (res float32) {
	switch t.Val.(type) {
	case int:
		return float32(t.Val.(int))
	case int32:
		return float32(t.Val.(int32))
	case int64:
		return float32(t.Val.(int64))
	case uint:
		return float32(t.Val.(uint))
	case uint32:
		return float32(t.Val.(uint32))
	case uint64:
		return float32(t.Val.(uint64))
	case float32:
		return t.Val.(float32)
	case float64:
		return float32(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseFloat(t.Val.(string), 32)
		return float32(t.Val.(float64))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToFloat64() (res float64) {
	switch t.Type {
	case IntElement:
		return float64(t.Val.(int))
	case Int32Element:
		return float64(t.Val.(int32))
	case Int64Element:
		return float64(t.Val.(int64))
	case UintElement:
		return float64(t.Val.(uint))
	case Uint32Element:
		return float64(t.Val.(uint32))
	case Uint64Element:
		return float64(t.Val.(uint64))
	case Float32Element:
		return float64(t.Val.(float32))
	case Float64Element:
		return t.Val.(float64)
	case StringElement:
		res, t.Error = strconv.ParseFloat(t.Val.(string), 64)
	case InterfaceElement:
		return t.interfaceToFloat64()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToFloat64() (res float64) {
	switch t.Val.(type) {
	case int:
		return float64(t.Val.(int))
	case int32:
		return float64(t.Val.(int32))
	case int64:
		return float64(t.Val.(int64))
	case uint:
		return float64(t.Val.(uint))
	case uint32:
		return float64(t.Val.(uint32))
	case uint64:
		return float64(t.Val.(uint64))
	case float32:
		return float64(t.Val.(float32))
	case float64:
		return t.Val.(float64)
	case string:
		res, t.Error = strconv.ParseFloat(t.Val.(string), 64)
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToString() (res string) {
	switch t.Type {
	case IntElement:
		return strconv.Itoa(t.Val.(int))
	case Int32Element:
		return strconv.FormatInt(int64(t.Val.(int32)), 10)
	case Int64Element:
		return strconv.FormatInt(t.Val.(int64), 10)
	case UintElement:
		return strconv.FormatUint(uint64(t.Val.(uint)), 10)
	case Uint32Element:
		return strconv.FormatUint(uint64(t.Val.(uint32)), 10)
	case Uint64Element:
		return strconv.FormatUint(t.Val.(uint64), 10)
	case Float32Element:
		return strconv.FormatFloat(float64(t.Val.(float32)), 'f', -1, 32)
	case Float64Element:
		return strconv.FormatFloat(t.Val.(float64), 'f', -1, 64)
	case StringElement:
		return t.Val.(string)
	case InterfaceElement:
		return t.interfaceToString()
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) interfaceToString() (res string) {
	switch t.Val.(type) {
	case int:
		return strconv.Itoa(t.Val.(int))
	case int32:
		return strconv.FormatInt(int64(t.Val.(int32)), 10)
	case int64:
		return strconv.FormatInt(t.Val.(int64), 10)
	case uint:
		return strconv.FormatUint(uint64(t.Val.(uint)), 10)
	case uint32:
		return strconv.FormatUint(uint64(t.Val.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(t.Val.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(t.Val.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t.Val.(float64), 'f', -1, 64)
	case string:
		return t.Val.(string)
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}
