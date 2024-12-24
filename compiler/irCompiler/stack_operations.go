package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
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

func DefineTypeofFunction(program *lang.Program) *ir.Func {

	typeofFn := program.Module.NewFunc("typeof", types.Void, ir.NewParam("a", types.I64), ir.NewParam("b", types.I64))
	typeofFnBody := typeofFn.NewBlock("entry")

	op1 := typeofFn.Params[1]

	typeofFnBody.NewCall(program.Funcs[lang.PushT].IrFunc, op1, constant.NewInt(types.I64, int64(lang.Type_T)))
	typeofFnBody.NewRet(nil)

	return typeofFn
}
func GetPreventValues(program *lang.Program, argc int, block *ir.Block) []value.Value {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	args := make([]value.Value, argc)
	//TO-DO make for types !!!
	for i := 0; i < argc; i++ {
		currentTop := block.NewLoad(types.I64, program.Globals["top"])
		newTop := block.NewSub(currentTop, constant.NewInt(types.I64, 1+int64(i)))
		stackPtr := block.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
		val := block.NewLoad(types.I64, stackPtr)
		args[i] = val
	}
	return args
}
func GetValues(program *lang.Program, argc int, block *ir.Block) []value.Value {

	args := make([]value.Value, argc*2)
	for i := 0; i < argc; i += 1 {
		args[i] = block.NewCall(program.Funcs[lang.PopT].IrFunc)
	}
	for i := argc; i < argc*2; i += 1 {
		args[i] = block.NewCall(program.Funcs[lang.PopTypeT].IrFunc)
	}
	return args
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
	load := dupFnBody.NewLoad(types.I64, stackPtr)

	typeStackPtr := dupFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), newTop)
	typeVal := dupFnBody.NewLoad(types.I64, typeStackPtr)

	dupFnBody.NewCall(program.Funcs[lang.PushT].IrFunc, load, typeVal)
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

	typePtr1 := swapFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), newTop)
	type1 := swapFnBody.NewLoad(types.I64, typePtr1)

	newTop = swapFnBody.NewSub(currentTop, constant.NewInt(types.I64, 2))

	stackPtr2 := swapFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value2 := swapFnBody.NewLoad(types.I64, stackPtr2)

	typePtr2 := swapFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), newTop)
	type2 := swapFnBody.NewLoad(types.I64, typePtr2)

	swapFnBody.NewStore(value1, stackPtr2)
	swapFnBody.NewStore(value2, stackPtr1)

	swapFnBody.NewStore(type1, typePtr2)
	swapFnBody.NewStore(type2, typePtr1)

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

	typePtr1 := overFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), newTop)
	typeVal := overFnBody.NewLoad(types.I64, typePtr1)

	overFnBody.NewCall(program.Funcs[lang.PushT].IrFunc, value, typeVal)
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

	typePtr1 := rotFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), newTop)
	type1 := rotFnBody.NewLoad(types.I64, typePtr1)

	newTop = rotFnBody.NewSub(currentTop, constant.NewInt(types.I64, 2))

	stackPtr2 := rotFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value2 := rotFnBody.NewLoad(types.I64, stackPtr2)
	typePtr2 := rotFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), newTop)
	type2 := rotFnBody.NewLoad(types.I64, typePtr2)

	newTop = rotFnBody.NewSub(currentTop, constant.NewInt(types.I64, 3))

	stackPtr3 := rotFnBody.NewGetElementPtr(stackType, program.Globals["stack"], constant.NewInt(types.I64, 0), newTop)
	value3 := rotFnBody.NewLoad(types.I64, stackPtr3)
	typePtr3 := rotFnBody.NewGetElementPtr(stackType, program.Globals["type_stack"], constant.NewInt(types.I64, 0), newTop)
	type3 := rotFnBody.NewLoad(types.I64, typePtr3)

	rotFnBody.NewStore(value1, stackPtr2)
	rotFnBody.NewStore(value2, stackPtr3)
	rotFnBody.NewStore(value3, stackPtr1)

	rotFnBody.NewStore(type1, typePtr2)
	rotFnBody.NewStore(type2, typePtr3)
	rotFnBody.NewStore(type3, typePtr1)

	rotFnBody.NewRet(nil)
	return rotFn
}
