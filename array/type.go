package array

type Type int

const (
	IntArray Type = iota + 1
	Int32Array
	Int64Array
	UintArray
	Uint32Array
	Uint64Array
	Float32Array
	Float64Array
	StringArray
	InterfaceArray
)
