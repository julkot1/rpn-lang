package lexer

import (
	"github.com/llir/llvm/ir"
	"rpn/lang"
)

const (
	PopToken          = `pop`
	AddToken          = `\+`
	SubToken          = `\-`
	MulToken          = `\*`
	ModToken          = `\%`
	PrintToken        = `print`
	DivToken          = `\/`
	PrintI8Token      = `puts`
	InputToken        = `input`
	DupToken          = `\|\>`
	SwapToken         = `\|s`
	OverToken         = `\|o`
	RotToken          = `\|r`
	StackPreventToken = `!`
	FunctionDefToken  = `fun`
	IfToken           = `if`
	IdentifierToken   = `([a-zA-Z_][a-zA-Z0-9_]*)`
	BlockOpenToken    = `{`
	BlockCloseToken   = `}`
	EqualsToken       = `==`
	NotToken          = `not`
	GreaterToken      = `>`
	LessToken         = `<`
	GreaterOrEqToken  = `>=`
	LessOrEqToken     = "<="
	NotEqualToken     = "!="
	OrToken           = `or`
	AndToken          = `and`
	RepeatToken       = `repeat`
)

var (
	ConstTokens = CreateConstTokens()
)

type ConstToken struct {
	Str             string
	Argc            int
	CanPrevent      bool
	DefaultFunction bool
	Lexable         bool
	Ir              *ir.Func
}

func CreateConstTokens() map[lang.TokenType]*ConstToken {
	consts := map[lang.TokenType]*ConstToken{}
	consts[lang.PushT] = &ConstToken{Str: "", Argc: 1, CanPrevent: true, DefaultFunction: true, Lexable: false}
	consts[lang.PopT] = &ConstToken{Str: PopToken, Argc: 0, CanPrevent: false, DefaultFunction: true, Lexable: true}
	consts[lang.AddT] = &ConstToken{Str: AddToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.ModT] = &ConstToken{Str: ModToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}

	consts[lang.SubT] = &ConstToken{Str: SubToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.MulT] = &ConstToken{Str: MulToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.DivT] = &ConstToken{Str: DivToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}

	consts[lang.PrintT] = &ConstToken{Str: PrintToken, Argc: 1, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.PrintI8T] = &ConstToken{Str: PrintI8Token, Argc: 1, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.InputT] = &ConstToken{Str: InputToken, Argc: 0, CanPrevent: false, DefaultFunction: true, Lexable: true}

	consts[lang.DupT] = &ConstToken{Str: DupToken, Argc: 0, CanPrevent: false, DefaultFunction: true, Lexable: true}
	consts[lang.SwapT] = &ConstToken{Str: SwapToken, Argc: 0, CanPrevent: false, DefaultFunction: true, Lexable: true}
	consts[lang.OverT] = &ConstToken{Str: OverToken, Argc: 0, CanPrevent: false, DefaultFunction: true, Lexable: true}
	consts[lang.RotT] = &ConstToken{Str: RotToken, Argc: 0, CanPrevent: false, DefaultFunction: true, Lexable: true}

	consts[lang.EqualsT] = &ConstToken{Str: EqualsToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.GreaterT] = &ConstToken{Str: GreaterToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.LessT] = &ConstToken{Str: LessToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.GreaterOrEqT] = &ConstToken{Str: GreaterOrEqToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.LessOrEqT] = &ConstToken{Str: LessOrEqToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.NotEqualT] = &ConstToken{Str: NotEqualToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}

	consts[lang.NotT] = &ConstToken{Str: NotToken, Argc: 1, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.AndT] = &ConstToken{Str: AndToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}
	consts[lang.OrT] = &ConstToken{Str: OrToken, Argc: 2, CanPrevent: true, DefaultFunction: true, Lexable: true}

	consts[lang.StackPreventT] = &ConstToken{Str: StackPreventToken, Argc: 1, CanPrevent: true, DefaultFunction: false, Lexable: true}
	consts[lang.IdentifierT] = &ConstToken{Str: IdentifierToken, Argc: 0, CanPrevent: true, DefaultFunction: false, Lexable: false}
	consts[lang.FunctionDefT] = &ConstToken{Str: FunctionDefToken, Argc: 0, CanPrevent: false, DefaultFunction: false, Lexable: true}
	consts[lang.FunctionT] = &ConstToken{Str: "", Argc: 0, CanPrevent: false, DefaultFunction: false, Lexable: false}
	consts[lang.BlockOpenT] = &ConstToken{Str: BlockOpenToken, Argc: 0, CanPrevent: false, DefaultFunction: false, Lexable: true}
	consts[lang.BlockCloseT] = &ConstToken{Str: BlockCloseToken, Argc: 0, CanPrevent: false, DefaultFunction: false, Lexable: true}
	consts[lang.BlockT] = &ConstToken{Str: "", Argc: 0, CanPrevent: false, DefaultFunction: false, Lexable: false}
	consts[lang.IfT] = &ConstToken{Str: IfToken, Argc: 1, CanPrevent: true, DefaultFunction: false, Lexable: true}
	consts[lang.RepeatT] = &ConstToken{Str: RepeatToken, Argc: 1, CanPrevent: true, DefaultFunction: false, Lexable: true}
	return consts
}
