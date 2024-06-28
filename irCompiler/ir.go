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

func DefineTypes() {
	types.I8Ptr = types.NewPointer(types.I8)
}

func NewProgram() *Program {
	program := &Program{}
	program.Module = ir.NewModule()

	DefineTypes()

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

	program.Funcs["dup"] = &Func{IrFunc: DefineDupFunc(program)}
	program.Funcs["swap"] = &Func{IrFunc: DefineSwapFunction(program)}
	program.Funcs["rot"] = &Func{IrFunc: DefineRotFunction(program)}
	program.Funcs["over"] = &Func{IrFunc: DefineOverFunction(program)}
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

func CallFunc(function *ir.Func, block *ir.Block, program *Program, tok lexer.Token) {
	if tok.StackPrevent == true {
		argc := tok.ArgsC()
		if argc == 0 {
			block.NewCall(function)
		} else {
			args := GetPreventValues(program, argc, block)
			if argc == 1 {
				block.NewCall(function, args[0])
			}
			if argc == 2 {
				block.NewCall(function, args[0], args[1])
			}
		}
	} else {
		argc := tok.ArgsC()
		if argc == 0 {
			block.NewCall(function)
		} else {
			args := GetValues(program, argc, block)
			if argc == 1 {
				block.NewCall(function, args[0])
			}
			if argc == 2 {
				block.NewCall(function, args[0], args[1])
			}
		}
	}

}

func LoadProgram(program *Program, arr []lexer.Token) {
	mainFn := program.Module.NewFunc("main", types.I32)
	mainFnBody := mainFn.NewBlock("entry")

	for i := 0; i < len(arr); i++ {
		switch arr[i].TokenType {
		case lexer.PushT:
			mainFnBody.NewCall(program.Funcs["push"].IrFunc, constant.NewInt(types.I32, int64(arr[i].Value.(int))))
			break
		case lexer.PopT:
			CallFunc(program.Funcs["pop"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.AddT:
			CallFunc(program.Funcs["add"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.SubT:
			CallFunc(program.Funcs["sub"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.MulT:
			CallFunc(program.Funcs["mul"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.PrintT:
			CallFunc(program.Funcs["print"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.PrintI8T:
			CallFunc(program.Funcs["printI8"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.InputT:
			CallFunc(program.Funcs["input"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.DupT:
			CallFunc(program.Funcs["dup"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.SwapT:
			CallFunc(program.Funcs["swap"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.RotT:
			CallFunc(program.Funcs["rot"].IrFunc, mainFnBody, program, arr[i])
			break
		case lexer.OverT:
			CallFunc(program.Funcs["over"].IrFunc, mainFnBody, program, arr[i])
			break
		default:
			panic("unhandled default case")
		}
	}
	mainFnBody.NewRet(constant.NewInt(types.I32, 0))
}
