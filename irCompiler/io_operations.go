package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
)

func DefinePrintFnFunction(program *lang.Program) *lang.DefaultFunc {
	printfType := types.NewFunc(types.I32, types.NewPointer(types.I8))
	printf := program.Module.NewFunc("printf", printfType)
	printfType.Variadic = true
	return &lang.DefaultFunc{IrFunc: printf, Type: printfType}
}

func DefineScanFnFunction(program *lang.Program) *lang.DefaultFunc {
	scanfType := types.NewFunc(types.I32, types.NewPointer(types.I8))
	scanf := program.Module.NewFunc("scanf", scanfType)
	scanfType.Variadic = true
	return &lang.DefaultFunc{IrFunc: scanf, Type: scanfType}
}

func DefinePrintFunction(program *lang.Program) *ir.Func {

	printFn := program.Module.NewFunc("print", types.Void, ir.NewParam("a", types.I32))
	printFnBody := printFn.NewBlock("entry")

	op1 := printFn.Params[0]

	formatPtr := printFnBody.NewGetElementPtr(types.NewArray(4, types.I8), program.Globals["format"], constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	printFnBody.NewCall(program.Funcs["printf"].IrFunc, formatPtr, op1)
	printFnBody.NewRet(nil)

	return printFn
}

func DefinePrintCharFunction(program *lang.Program) *ir.Func {
	printFn := program.Module.NewFunc("printI8", types.Void, ir.NewParam("a", types.I32))
	printFnBody := printFn.NewBlock("entry")

	op1 := printFn.Params[0]
	truncated := printFnBody.NewTrunc(op1, types.I8)

	formatPtr := printFnBody.NewGetElementPtr(types.NewArray(4, types.I8), program.Globals["formatChar"], constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	printFnBody.NewCall(program.Funcs["printf"].IrFunc, formatPtr, truncated)
	printFnBody.NewRet(nil)

	return printFn
}

func DefineInputFunction(program *lang.Program) *ir.Func {
	inputFn := program.Module.NewFunc("input", types.Void)
	inputFnBody := inputFn.NewBlock("entry")
	inputPtr := inputFnBody.NewAlloca(types.I32)

	formatPtr := inputFnBody.NewGetElementPtr(types.NewArray(3, types.I8), program.Globals["formatIn"], constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	inputFnBody.NewCall(program.Funcs["scanf"].IrFunc, formatPtr, inputPtr)

	input := inputFnBody.NewLoad(types.I32, inputPtr)
	inputFnBody.NewCall(program.Funcs["push"].IrFunc, input)
	inputFnBody.NewRet(nil)
	return inputFn

}
