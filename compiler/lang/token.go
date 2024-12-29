package lang

import (
	"fmt"
	"github.com/timtadh/lexmachine/machines"
)

const (
	PopToken         = `pop`
	AddToken         = `+`
	SubToken         = `-`
	MulToken         = `*`
	ModToken         = `%`
	PrintToken       = `print`
	DivToken         = `/`
	PrintI8Token     = `puts`
	InputToken       = `input`
	DupToken         = `dup`
	SwapToken        = `swap`
	OverToken        = `over`
	RotToken         = `rot`
	FunctionDefToken = `fun`
	IfToken          = `if`
	BlockOpenToken   = `{`
	BlockCloseToken  = `}`
	EqualsToken      = `==`
	NotToken         = `not`
	GreaterToken     = `>`
	LessToken        = `<`
	GreaterOrEqToken = `>=`
	LessOrEqToken    = "<="
	NotEqualToken    = "!="
	OrToken          = `or`
	AndToken         = `and`
	RepeatToken      = `repeat`
	TypeofToken      = `typeof`
	AtToken          = `at`
	AssignToken      = `:=`
	LenToken         = `len`
	SetToken         = `set`
)

type TokenType int

const (
	PushT TokenType = iota
	PopT
	AddT
	SubT
	MulT
	PrintT
	DivT
	ModT
	PrintI8T
	InputT
	DupT
	SwapT
	OverT
	RotT
	StackPreventT
	IdentifierT
	FunctionDefT
	FunctionT
	BlockOpenT
	BlockCloseT
	BlockT
	IfT
	EqualsT
	NotT
	GreaterT
	LessT
	GreaterOrEqT
	LessOrEqT
	NotEqualT
	OrT
	AndT
	RepeatT
	PopTypeT
	AtT
	LenT
	TypeofT
	AssignT
	SetT
)

type Token struct {
	TokenType    TokenType
	Value        interface{}
	StackPrevent bool
	Match        *machines.Match
}

func (t Token) String() string {
	return fmt.Sprintf("{tokenID: %d, value: %d, prevent: %t, token: %s }", t.TokenType, t.Value.(int), t.StackPrevent, t.Match.Bytes)
}

func StrToTokenType(str string) TokenType {
	switch str {
	case PopToken:
		return PopT
	case AddToken:
		return AddT
	case SubToken:
		return SubT
	case MulToken:
		return MulT
	case ModToken:
		return ModT
	case PrintToken:
		return PrintT
	case DivToken:
		return DivT
	case PrintI8Token:
		return PrintI8T
	case InputToken:
		return InputT
	case DupToken:
		return DupT
	case SwapToken:
		return SwapT
	case OverToken:
		return OverT
	case RotToken:
		return RotT
	case FunctionDefToken:
		return FunctionDefT
	case IfToken:
		return IfT
	case BlockOpenToken:
		return BlockOpenT
	case BlockCloseToken:
		return BlockCloseT
	case EqualsToken:
		return EqualsT
	case NotToken:
		return NotT
	case GreaterToken:
		return GreaterT
	case LessToken:
		return LessT
	case GreaterOrEqToken:
		return GreaterOrEqT
	case LessOrEqToken:
		return LessOrEqT
	case NotEqualToken:
		return NotEqualT
	case OrToken:
		return OrT
	case AndToken:
		return AndT
	case RepeatToken:
		return RepeatT
	case TypeofToken:
		return TypeofT
	case AtToken:
		return AtT
	case AssignToken:
		return AssignT
	case LenToken:
		return LenT
	case SetToken:
		return SetT
	}

	return AddT
}
