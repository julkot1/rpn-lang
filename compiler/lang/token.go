package lang

import (
	"fmt"
	"github.com/timtadh/lexmachine/machines"
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
	PrintFT
	ScanFT
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
