package lexer

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
	PrintI8T
	InputT
	DupT
	SwapT
	OverT
	RotT
	StackPreventT
)

type Token struct {
	TokenType    TokenType
	Value        interface{}
	StackPrevent bool
	Match        *machines.Match
}

func (t Token) ArgsC() int {
	if t.TokenType == PushT || t.TokenType == PrintT || t.TokenType == PrintI8T {
		return 1
	}
	if t.TokenType == AddT ||
		t.TokenType == SubT ||
		t.TokenType == MulT ||
		t.TokenType == DivT {
		return 2
	}
	return 0
}

func (t Token) String() string {
	return fmt.Sprintf("{tokenID: %d, value: %d, prevent: %t, token: %s }", t.TokenType, t.Value.(int), t.StackPrevent, t.Match.Bytes)
}

func (t Token) CanPrevent() bool {
	return !(t.TokenType == PopT ||
		t.TokenType == PushT ||
		t.TokenType == DupT ||
		t.TokenType == OverT ||
		t.TokenType == RotT ||
		t.TokenType == SubT)

}
