package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
)

func DefinePushFunction(program *lang.Program) *ir.Func {
	pushFn := program.Module.NewFunc("push", types.Void, ir.NewParam("value", types.I64), ir.NewParam("type", types.I64))
	pushFnBody := pushFn.NewBlock("entry")

	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	currentTop := pushFnBody.NewLoad(types.I64, program.Globals["top"])
	stackPtr := pushFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), currentTop)
	typeStackPtr := pushFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), currentTop)
	pushFnBody.NewStore(pushFn.Params[0], stackPtr)
	pushFnBody.NewStore(pushFn.Params[1], typeStackPtr)
	newTop := pushFnBody.NewAdd(currentTop, constant.NewInt(types.I64, 1))
	pushFnBody.NewStore(newTop, program.Globals["top"])

	pushFnBody.NewRet(nil)

	return pushFn
}

func DefinePopFunction(program *lang.Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	popFn := program.Module.NewFunc("pop", types.I64)
	popFnBody := popFn.NewBlock("entry")

	currentTop := popFnBody.NewLoad(types.I64, program.Globals["top"])
	newTop := popFnBody.NewSub(currentTop, constant.NewInt(types.I64, 1))
	popFnBody.NewStore(newTop, program.Globals["top"])

	currentTop = popFnBody.NewLoad(types.I64, program.Globals["top"])
	stackPtr := popFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), currentTop)
	value := popFnBody.NewLoad(types.I64, stackPtr)

	popFnBody.NewRet(value)

	return popFn
}

func DefineTypePopFunction(program *lang.Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	popFn := program.Module.NewFunc("pop_type", types.I64)
	popFnBody := popFn.NewBlock("entry")

	currentTop := popFnBody.NewLoad(types.I64, program.Globals["top"])
	stackPtr := popFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), currentTop)
	value := popFnBody.NewLoad(types.I64, stackPtr)

	popFnBody.NewRet(value)

	return popFn
}

func GetPreventValues(program *lang.Program, argc int, block *ir.Block) []*ir.InstLoad {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	args := make([]*ir.InstLoad, argc)
	for i := 0; i < argc; i++ {
		currentTop := block.NewLoad(types.I64, program.Globals["top"])
		newTop := block.NewSub(currentTop, constant.NewInt(types.I64, 1+int64(i)))
		stackPtr := block.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
		val := block.NewLoad(types.I64, stackPtr)
		args[i] = val
	}
	return args
}
func GetValues(program *lang.Program, argc int, block *ir.Block) []*ir.InstCall {

	args := make([]*ir.InstCall, argc*2)
	for i := 0; i < argc*2; i += 2 {
		args[i] = block.NewCall(program.Funcs["pop"].IrFunc)
		args[i+1] = block.NewCall(program.Funcs["pop_type"].IrFunc)
	}
	return args
}

func DefineAddFunc(program *lang.Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "add")
	result := binFnBody.NewAdd(a, b)

	binFnBody.NewCall(program.Funcs["push"].IrFunc, result, constant.NewInt(types.I64, int64(0)))
	binFnBody.NewRet(nil)
	return binFn
}
func DefineModFunc(program *lang.Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "mod")
	result := binFnBody.NewSRem(a, b)

	binFnBody.NewCall(program.Funcs["push"].IrFunc, result, constant.NewInt(types.I64, int64(0)))
	binFnBody.NewRet(nil)
	return binFn
}
func DefineSubFunc(program *lang.Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "sub")
	result := binFnBody.NewSub(a, b)

	binFnBody.NewCall(program.Funcs["push"].IrFunc, result, constant.NewInt(types.I64, int64(0)))
	binFnBody.NewRet(nil)
	return binFn
}
func DefineDivFunc(program *lang.Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "div")
	result := binFnBody.NewSDiv(a, b)

	binFnBody.NewCall(program.Funcs["push"].IrFunc, result, constant.NewInt(types.I64, int64(0)))
	binFnBody.NewRet(nil)
	return binFn
}
func DefineMulFunc(program *lang.Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "mul")
	result := binFnBody.NewMul(a, b)

	binFnBody.NewCall(program.Funcs["push"].IrFunc, result, constant.NewInt(types.I64, int64(0)))
	binFnBody.NewRet(nil)
	return binFn
}

func DefineBinaryFunction(program *lang.Program, name string) (*ir.Block, *ir.Func, *ir.Param, *ir.Param) {

	binFn := program.Module.NewFunc(name, types.Void, ir.NewParam("a", types.I64), ir.NewParam("b", types.I64))
	binFnBody := binFn.NewBlock("entry")

	op1 := binFn.Params[0]
	op2 := binFn.Params[1]

	return binFnBody, binFn, op1, op2
}

func DefineDupFunc(program *lang.Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	dupFn := program.Module.NewFunc("dup", types.Void)
	dupFnBody := dupFn.NewBlock("entry")

	currentTop := dupFnBody.NewLoad(types.I64, program.Globals["top"])
	newTop := dupFnBody.NewSub(currentTop, constant.NewInt(types.I64, 1))

	stackPtr := dupFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value := dupFnBody.NewLoad(types.I64, stackPtr)

	dupFnBody.NewCall(program.Funcs["push"].IrFunc, value, constant.NewInt(types.I64, int64(0)))
	dupFnBody.NewRet(nil)
	return dupFn
}
func DefineSwapFunction(program *lang.Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	swapFn := program.Module.NewFunc("swap", types.Void)
	swapFnBody := swapFn.NewBlock("entry")

	currentTop := swapFnBody.NewLoad(types.I64, program.Globals["top"])
	newTop := swapFnBody.NewSub(currentTop, constant.NewInt(types.I64, 1))

	stackPtr1 := swapFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value1 := swapFnBody.NewLoad(types.I64, stackPtr1)

	newTop = swapFnBody.NewSub(currentTop, constant.NewInt(types.I64, 2))

	stackPtr2 := swapFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value2 := swapFnBody.NewLoad(types.I64, stackPtr2)

	swapFnBody.NewStore(value1, stackPtr2)
	swapFnBody.NewStore(value2, stackPtr1)

	swapFnBody.NewRet(nil)
	return swapFn
}
func DefineOverFunction(program *lang.Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	overFn := program.Module.NewFunc("over", types.Void)
	overFnBody := overFn.NewBlock("entry")

	currentTop := overFnBody.NewLoad(types.I64, program.Globals["top"])
	newTop := overFnBody.NewSub(currentTop, constant.NewInt(types.I64, 2))

	stackPtr := overFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value := overFnBody.NewLoad(types.I64, stackPtr)

	overFnBody.NewCall(program.Funcs["push"].IrFunc, value, constant.NewInt(types.I64, int64(0)))
	overFnBody.NewRet(nil)
	return overFn
}
func DefineRotFunction(program *lang.Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	rotFn := program.Module.NewFunc("rot", types.Void)
	rotFnBody := rotFn.NewBlock("entry")

	currentTop := rotFnBody.NewLoad(types.I64, program.Globals["top"])
	newTop := rotFnBody.NewSub(currentTop, constant.NewInt(types.I64, 1))

	stackPtr1 := rotFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value1 := rotFnBody.NewLoad(types.I64, stackPtr1)

	newTop = rotFnBody.NewSub(currentTop, constant.NewInt(types.I64, 2))

	stackPtr2 := rotFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value2 := rotFnBody.NewLoad(types.I64, stackPtr2)

	newTop = rotFnBody.NewSub(currentTop, constant.NewInt(types.I64, 3))

	stackPtr3 := rotFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value3 := rotFnBody.NewLoad(types.I64, stackPtr3)

	rotFnBody.NewStore(value1, stackPtr2)
	rotFnBody.NewStore(value2, stackPtr3)
	rotFnBody.NewStore(value3, stackPtr1)

	rotFnBody.NewRet(nil)
	return rotFn
}
