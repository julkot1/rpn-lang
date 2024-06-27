package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lexer"
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
	program.Funcs["scanf"] = DefineScanFnFunction(program)
	program.Funcs["push"] = &Func{IrFunc: DefinePushFunction(program)}
	program.Funcs["pop"] = &Func{IrFunc: DefinePopFunction(program)}
	program.Funcs["add"] = &Func{IrFunc: DefineAddFunc(program)}
	program.Funcs["sub"] = &Func{IrFunc: DefineSubFunc(program)}
	program.Funcs["mul"] = &Func{IrFunc: DefineMulFunc(program)}
	program.Funcs["print"] = &Func{IrFunc: DefinePrintFunction(program)}
	program.Funcs["printI8"] = &Func{IrFunc: DefinePrintCharFunction(program)}
	program.Funcs["input"] = &Func{IrFunc: DefineInputFunction(program)}

}

func DefineGlobals(m *ir.Module) map[string]*ir.Global {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I32)

	globals := make(map[string]*ir.Global)

	globals["stack"] = m.NewGlobalDef("vstack", constant.NewZeroInitializer(stackType))
	globals["top"] = m.NewGlobalDef("top", constant.NewInt(types.I32, 0))

	formatStr := constant.NewCharArrayFromString("%d\n\x00")
	globals["format"] = m.NewGlobalDef("format", formatStr)
	globals["format"].Immutable = true
	globals["format"].Linkage = enum.LinkagePrivate

	formatChar := constant.NewCharArrayFromString("%c\n\x00")
	globals["formatChar"] = m.NewGlobalDef("formatChar", formatChar)
	globals["formatChar"].Immutable = true
	globals["formatChar"].Linkage = enum.LinkagePrivate

	formatIn := constant.NewCharArrayFromString("%d\x00")
	globals["formatIn"] = m.NewGlobalDef("formatIn", formatIn)
	globals["formatIn"].Immutable = true
	globals["formatIn"].Linkage = enum.LinkagePrivate

	return globals
}

func LoadProgram(program *Program, arr []lexer.Token) {
	mainFn := program.Module.NewFunc("main", types.I32)
	mainFnBody := mainFn.NewBlock("entry")

	for i := 0; i < len(arr); i++ {
		switch arr[i].TokenType {
		case lexer.PushT:
			mainFnBody.NewCall(program.Funcs["push"].IrFunc, constant.NewInt(types.I32, int64(arr[i].Value)))
			break
		case lexer.PopT:
			mainFnBody.NewCall(program.Funcs["pop"].IrFunc)
			break
		case lexer.AddT:
			mainFnBody.NewCall(program.Funcs["add"].IrFunc)
			break
		case lexer.SubT:
			mainFnBody.NewCall(program.Funcs["sub"].IrFunc)
			break
		case lexer.MulT:
			mainFnBody.NewCall(program.Funcs["mul"].IrFunc)
			break
		case lexer.PrintT:
			mainFnBody.NewCall(program.Funcs["print"].IrFunc)
			break
		case lexer.PrintI8T:
			mainFnBody.NewCall(program.Funcs["printI8"].IrFunc)
			break
		case lexer.InputT:
			mainFnBody.NewCall(program.Funcs["input"].IrFunc)
			break
		}
	}
	mainFnBody.NewRet(constant.NewInt(types.I32, 0))
}
