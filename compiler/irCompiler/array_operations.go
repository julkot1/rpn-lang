package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
)

func DefineAtFunction(program *lang.Program) *ir.Func {
	printFn := program.Module.NewFunc("print", types.Void, ir.NewParam("a", types.I64), ir.NewParam("b", types.I64))
	printFnBody := printFn.NewBlock("entry")

	//op1 := printFn.Params[0]

	printFnBody.NewRet(nil)
	return printFn
}
