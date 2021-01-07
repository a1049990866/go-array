package array

import (
	"github.com/a1049990866/go-array/elemen"
)

type Array struct {
	Data          interface{}
	Type          Type
	Error         error
	deduplication bool
	dict          interface{}
}

func NewArray(val interface{}, typ Type) *Array {
	return &Array{
		Data: val,
		Type: typ,
	}
}

func (t *Array) Deduplication(v bool) *Array {
	t.deduplication = v
	return t
}

func (t *Array) base(f func(v interface{}, typ elemen.Type)) {
	switch t.Type {
	case IntArray:
		for _, v := range t.Data.([]int) {
			f(v, elemen.IntElement)
		}
	case Int32Array:
		for _, v := range t.Data.([]int32) {
			f(v, elemen.Int32Element)

		}
	case Int64Array:
		for _, v := range t.Data.([]int64) {
			f(v, elemen.Int64Element)
		}
	case UintArray:
		for _, v := range t.Data.([]uint) {
			f(v, elemen.UintElement)
		}
	case Uint32Array:
		for _, v := range t.Data.([]uint32) {
			f(v, elemen.Uint32Element)
		}
	case Uint64Array:
		for _, v := range t.Data.([]uint64) {
			f(v, elemen.Uint64Element)
		}
	case Float32Array:
		for _, v := range t.Data.([]float32) {
			f(v, elemen.Float32Element)
		}
	case Float64Array:
		for _, v := range t.Data.([]float64) {
			f(v, elemen.Float64Element)
		}
	case StringArray:
		for _, v := range t.Data.([]string) {
			f(v, elemen.StringElement)
		}
	case InterfaceArray:
		for _, v := range t.Data.([]interface{}) {
			f(v, elemen.InterfaceElement)
		}
	}
}

func (t *Array) ToIntArray() (res []int) {
	m := make(map[int]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToInt()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToInt32Array() (res []int32) {
	m := make(map[int32]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToInt32()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToInt64Array() (res []int64) {
	m := make(map[int64]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToInt64()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToUintArray() (res []uint) {
	m := make(map[uint]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToUint()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToUint32Array() (res []uint32) {
	m := make(map[uint32]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToUint32()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToUint64Array() (res []uint64) {
	m := make(map[uint64]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToUint64()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToFloat32Array() (res []float32) {
	m := make(map[float32]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToFloat32()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToFloat64Array() (res []float64) {
	m := make(map[float64]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToFloat64()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToStringArray() (res []string) {
	m := make(map[string]int)
	f := func(v interface{}, typ elemen.Type) {
		any := elemen.NewElement(v, typ)
		vv := any.ToString()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToInterfaceArray() (res []interface{}) {
	m := make(map[interface{}]int)
	f := func(v interface{}, typ elemen.Type) {
		if !t.deduplication || m[v] == 0 {
			res = append(res, v)
		}
		m[v] += 1
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToIntMap() (res map[int]int) {
	switch t.dict.(type) {
	case map[int]int:
		return t.dict.(map[int]int)
	default:
		t.ToIntArray()
		return t.ToIntMap()
	}
}

func (t *Array) ToInt32Map() (res map[int32]int) {
	switch t.dict.(type) {
	case map[int32]int:
		return t.dict.(map[int32]int)
	default:
		t.ToInt32Array()
		return t.ToInt32Map()
	}
}

func (t *Array) ToInt64Map() (res map[int64]int) {
	switch t.dict.(type) {
	case map[int64]int:
		return t.dict.(map[int64]int)
	default:
		t.ToInt64Array()
		return t.ToInt64Map()
	}
}

func (t *Array) ToUintMap() (res map[uint]int) {
	switch t.dict.(type) {
	case map[uint]int:
		return t.dict.(map[uint]int)
	default:
		t.ToUintArray()
		return t.ToUintMap()
	}
}

func (t *Array) ToUint32Map() (res map[uint32]int) {
	switch t.dict.(type) {
	case map[uint32]int:
		return t.dict.(map[uint32]int)
	default:
		t.ToUint32Array()
		return t.ToUint32Map()
	}
}

func (t *Array) ToUint64Map() (res map[uint64]int) {
	switch t.dict.(type) {
	case map[uint64]int:
		return t.dict.(map[uint64]int)
	default:
		t.ToUint64Array()
		return t.ToUint64Map()
	}
}

func (t *Array) ToFloat32Map() (res map[float32]int) {
	switch t.dict.(type) {
	case map[float32]int:
		return t.dict.(map[float32]int)
	default:
		t.ToFloat32Array()
		return t.ToFloat32Map()
	}
}

func (t *Array) ToFloat64Map() (res map[float64]int) {
	switch t.dict.(type) {
	case map[float64]int:
		return t.dict.(map[float64]int)
	default:
		t.ToFloat64Array()
		return t.ToFloat64Map()
	}
}

func (t *Array) ToStringMap() (res map[string]int) {
	switch t.dict.(type) {
	case map[string]int:
		return t.dict.(map[string]int)
	default:
		t.ToStringArray()
		return t.ToStringMap()
	}
}

func (t *Array) ToInterfaceMap() (res map[interface{}]int) {
	switch t.dict.(type) {
	case map[interface{}]int:
		return t.dict.(map[interface{}]int)
	default:
		t.ToInterfaceArray()
		return t.ToInterfaceMap()
	}
}
