package irCompiler

import (
	"rpn/lang"
)

func ParseToken(text string, block *lang.Block, fun *lang.Function, program *lang.Program) {
	tok, ok := program.GlobalTokenTable[text]
	if !ok {
		return
	}
	switch tok {
	case lang.PStaticFunc:
		CallFunc(program.StaticFunctions[text].IrFunc, block.Ir, program, false)
		break
	case lang.PStcFunction:
		CallFunc(program.GetFunction(text, fun).Ir, block.Ir, program, false)
		break
	default:
		panic("unhandled default case")

	}
}
