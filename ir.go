package main

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type Func struct {
	IrFunc *ir.Func
	Type   *types.FuncType
}

type Program struct {
	Module  *ir.Module
	Globals map[string]*ir.Global
	Funcs   map[string]*Func
}

func DefineTypes(m *ir.Module) {
	types.I8Ptr = types.NewPointer(types.I8)
}

func NewProgram() *Program {
	program := &Program{}
	program.Module = ir.NewModule()

	DefineTypes(program.Module)

	program.Globals = DefineGlobals(program.Module)

	DefineFuncs(program)

	return program
}

func DefineFuncs(program *Program) {
	program.Funcs = make(map[string]*Func)
	program.Funcs["printf"] = DefinePrintFnFunction(program)
	program.Funcs["push"] = &Func{IrFunc: DefinePushFunction(program)}
	program.Funcs["pop"] = &Func{IrFunc: DefinePopFunction(program)}
	program.Funcs["add"] = &Func{IrFunc: DefineAddFunc(program)}
	program.Funcs["sub"] = &Func{IrFunc: DefineSubFunc(program)}
	program.Funcs["mul"] = &Func{IrFunc: DefineMulFunc(program)}
	program.Funcs["print"] = &Func{IrFunc: DefinePrintFunction(program)}

}
func DefinePrintFnFunction(program *Program) *Func {
	printfType := types.NewFunc(types.I32, types.NewPointer(types.I8))
	printf := program.Module.NewFunc("printf", printfType)
	printfType.Variadic = true
	return &Func{printf, printfType}
}

func DefineGlobals(m *ir.Module) map[string]*ir.Global {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	globals := make(map[string]*ir.Global)

	globals["stack"] = m.NewGlobalDef("vstack", constant.NewZeroInitializer(stackType))
	globals["top"] = m.NewGlobalDef("top", constant.NewInt(types.I32, 0))

	formatStr := constant.NewCharArrayFromString("%d\n")
	globals["format"] = m.NewGlobalDef("formatStr", formatStr)

	return globals
}

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

func DefinePrintFunction(program *Program) *ir.Func {

	printFn := program.Module.NewFunc("print", types.Void)
	printFnBody := printFn.NewBlock("entry")

	op1 := printFnBody.NewCall(program.Funcs["pop"].IrFunc)

	formatPtr := printFnBody.NewGetElementPtr(types.NewArray(3, types.I8), program.Globals["format"], constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	printFnBody.NewCall(program.Funcs["printf"].IrFunc, formatPtr, op1)
	printFnBody.NewRet(nil)

	return printFn
}

func LoadProgram(program *Program, arr []Token) {
	mainFn := program.Module.NewFunc("main", types.I32)
	mainFnBody := mainFn.NewBlock("entry")

	for i := 0; i < len(arr); i++ {
		switch arr[i].TokenType {
		case PushT:
			mainFnBody.NewCall(program.Funcs["push"].IrFunc, constant.NewInt(types.I32, int64(arr[i].Value)))
			break
		case PopT:
			mainFnBody.NewCall(program.Funcs["pop"].IrFunc)
			break
		case AddT:
			mainFnBody.NewCall(program.Funcs["add"].IrFunc)
			break
		case SubT:
			mainFnBody.NewCall(program.Funcs["sub"].IrFunc)
			break
		case MulT:
			mainFnBody.NewCall(program.Funcs["mul"].IrFunc)
			break
		case PrintT:
			mainFnBody.NewCall(program.Funcs["print"].IrFunc)
			break
		}
	}
	mainFnBody.NewRet(constant.NewInt(types.I32, 0))
}
