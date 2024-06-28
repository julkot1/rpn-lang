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

func GetPreventValues(program *Program, argc int, block *ir.Block) []*ir.InstLoad {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	args := make([]*ir.InstLoad, argc)
	for i := 0; i < argc; i++ {
		currentTop := block.NewLoad(types.I32, program.Globals["top"])
		newTop := block.NewSub(currentTop, constant.NewInt(types.I32, 1+int64(i)))
		stackPtr := block.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), newTop)
		val := block.NewLoad(types.I32, stackPtr)
		args[i] = val
	}
	return args
}
func GetValues(program *Program, argc int, block *ir.Block) []*ir.InstCall {

	args := make([]*ir.InstCall, argc)
	for i := 0; i < argc; i++ {
		args[i] = block.NewCall(program.Funcs["pop"].IrFunc)
	}
	return args
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

func DefineBinaryFunction(program *Program, name string) (*ir.Block, *ir.Func, *ir.Param, *ir.Param) {

	binFn := program.Module.NewFunc(name, types.Void, ir.NewParam("a", types.I32), ir.NewParam("b", types.I32))
	binFnBody := binFn.NewBlock("entry")

	op1 := binFn.Params[0]
	op2 := binFn.Params[1]

	return binFnBody, binFn, op1, op2
}

func DefineDupFunc(program *Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	dupFn := program.Module.NewFunc("dup", types.Void)
	dupFnBody := dupFn.NewBlock("entry")

	currentTop := dupFnBody.NewLoad(types.I32, program.Globals["top"])
	newTop := dupFnBody.NewSub(currentTop, constant.NewInt(types.I32, 1))

	stackPtr := dupFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), newTop)
	value := dupFnBody.NewLoad(types.I32, stackPtr)

	dupFnBody.NewCall(program.Funcs["push"].IrFunc, value)
	dupFnBody.NewRet(nil)
	return dupFn
}
func DefineSwapFunction(program *Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	swapFn := program.Module.NewFunc("swap", types.Void)
	swapFnBody := swapFn.NewBlock("entry")

	currentTop := swapFnBody.NewLoad(types.I32, program.Globals["top"])
	newTop := swapFnBody.NewSub(currentTop, constant.NewInt(types.I32, 1))

	stackPtr1 := swapFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), newTop)
	value1 := swapFnBody.NewLoad(types.I32, stackPtr1)

	newTop = swapFnBody.NewSub(currentTop, constant.NewInt(types.I32, 2))

	stackPtr2 := swapFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), newTop)
	value2 := swapFnBody.NewLoad(types.I32, stackPtr2)

	swapFnBody.NewStore(value1, stackPtr2)
	swapFnBody.NewStore(value2, stackPtr1)

	swapFnBody.NewRet(nil)
	return swapFn
}
func DefineOverFunction(program *Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	overFn := program.Module.NewFunc("over", types.Void)
	overFnBody := overFn.NewBlock("entry")

	currentTop := overFnBody.NewLoad(types.I32, program.Globals["top"])
	newTop := overFnBody.NewSub(currentTop, constant.NewInt(types.I32, 2))

	stackPtr := overFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), newTop)
	value := overFnBody.NewLoad(types.I32, stackPtr)

	overFnBody.NewCall(program.Funcs["push"].IrFunc, value)
	overFnBody.NewRet(nil)
	return overFn
}
func DefineRotFunction(program *Program) *ir.Func {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	rotFn := program.Module.NewFunc("rot", types.Void)
	rotFnBody := rotFn.NewBlock("entry")

	currentTop := rotFnBody.NewLoad(types.I32, program.Globals["top"])
	newTop := rotFnBody.NewSub(currentTop, constant.NewInt(types.I32, 1))

	stackPtr1 := rotFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), newTop)
	value1 := rotFnBody.NewLoad(types.I32, stackPtr1)

	newTop = rotFnBody.NewSub(currentTop, constant.NewInt(types.I32, 2))

	stackPtr2 := rotFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), newTop)
	value2 := rotFnBody.NewLoad(types.I32, stackPtr2)

	newTop = rotFnBody.NewSub(currentTop, constant.NewInt(types.I32, 3))

	stackPtr3 := rotFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I32, 0), newTop)
	value3 := rotFnBody.NewLoad(types.I32, stackPtr3)

	rotFnBody.NewStore(value1, stackPtr2)
	rotFnBody.NewStore(value2, stackPtr3)
	rotFnBody.NewStore(value3, stackPtr1)

	rotFnBody.NewRet(nil)
	return rotFn
}
