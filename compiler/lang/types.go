package lang

type Type int64

const (
	INT_T Type = iota
	BOOL_T
	CHAR_T
	FLOAT_T
	STRING_T
	ARRAY_T
	REF_T
	Type_T
	ANY_T
)

type PushableToken struct {
	Value interface{}
	Typ   Type
}
