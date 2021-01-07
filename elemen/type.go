package elemen

type Type int

const (
	IntElement Type = iota + 1
	Int32Element
	Int64Element
	UintElement
	Uint32Element
	Uint64Element
	Float32Element
	Float64Element
	StringElement
	InterfaceElement
)
