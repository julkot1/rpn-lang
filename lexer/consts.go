package lexer

const (
	PopToken          = `pop`
	AddToken          = `\+`
	SubToken          = `\-`
	MulToken          = `\*`
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
	IdentifierToken   = `([a-zA-Z_][a-zA-Z0-9_]*)`
	BlockOpenToken    = `{`
	BlockCloseToken   = `}`
)
