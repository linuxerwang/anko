package ast

// Token is used in the lexer to split characters into a string called a token
type Token struct {
	PosImpl
	Tok int
	Lit string
}

// TypeKind blah
type TypeKind int

const (
	// TypeDefault default type
	TypeDefault TypeKind = iota
	// TypePtr ptr type
	TypePtr
	// TypeSlice slice type
	TypeSlice
	// TypeMap map type
	TypeMap
	// TypeChan chan type
	TypeChan
)

// TypeStruct blah
type TypeStruct struct {
	Kind       TypeKind
	Env        []string
	Name       string
	Dimensions int
	SubType    *TypeStruct
	Key        *TypeStruct
}
