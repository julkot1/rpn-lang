package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func DefinePushFunction(program *Program) *ir.Func {
	pushFn := program.Module.NewFunc("push", types.Void, ir.NewParam("value", types.I32))
	pushFnBody := pushFn.NewBlock("entry")

	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	currentTop := pushFnBody.NewLoad(types.I32, program.Globals["top"])
	stackPtr := pushFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), currentTop)
	pushFnBody.NewStore(pushFn.Params[0], stackPtr)
	newTop := pushFnBody.NewAdd(currentTop, constant.NewInt(types.I32, 1))
	pushFnBody.NewStore(newTop, program.Globals["top"])
	pushFnBody.NewRet(nil)

	return pushFn
}

func DefinePopFunction(program *Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	popFn := program.Module.NewFunc("pop", types.I32)
	popFnBody := popFn.NewBlock("entry")

	currentTop := popFnBody.NewLoad(types.I32, program.Globals["top"])
	newTop := popFnBody.NewSub(currentTop, constant.NewInt(types.I32, 1))
	popFnBody.NewStore(newTop, program.Globals["top"])

	currentTop = popFnBody.NewLoad(types.I32, program.Globals["top"])
	stackPtr := popFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), currentTop)
	value := popFnBody.NewLoad(types.I32, stackPtr)

	popFnBody.NewRet(value)

	return popFn
}

func DefineAddFunc(program *Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "add")
	result := binFnBody.NewAdd(a, b)

	binFnBody.NewCall(program.Funcs["push"].IrFunc, result)
	binFnBody.NewRet(nil)
	return binFn
}
func DefineSubFunc(program *Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "div")
	result := binFnBody.NewSub(a, b)

	binFnBody.NewCall(program.Funcs["push"].IrFunc, result)
	binFnBody.NewRet(nil)
	return binFn
}
func DefineMulFunc(program *Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "mul")
	result := binFnBody.NewMul(a, b)

	binFnBody.NewCall(program.Funcs["push"].IrFunc, result)
	binFnBody.NewRet(nil)
	return binFn
}

func DefineBinaryFunction(program *Program, name string) (*ir.Block, *ir.Func, *ir.InstCall, *ir.InstCall) {

	binFn := program.Module.NewFunc(name, types.Void)
	binFnBody := binFn.NewBlock("entry")

	op1 := binFnBody.NewCall(program.Funcs["pop"].IrFunc)
	op2 := binFnBody.NewCall(program.Funcs["pop"].IrFunc)

	return binFnBody, binFn, op1, op2
}
