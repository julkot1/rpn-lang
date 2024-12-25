package irCompiler

import (
	"fmt"
	"os"
	"rpn/lang"
	"rpn/parser"
	"rpn/util"
)

func ParseToken(text string, block *lang.Block, fun *lang.Function, program *lang.Program, scope util.Stack, ctx *parser.IdentifierContext) {
	tok, ok := program.GlobalTokenTable[text]
	if !ok {
		if TokenExists(text, scope) {
			token := GetToken(text, scope)
			PushToken(token, block, program)
			return
		}
		line := ctx.GetStart().GetLine()
		fmt.Printf("variable '%s' does not exist:\n \tline %v\n", text, line)
		os.Exit(1)
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
		break
	}
}
