package irCompiler

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
	"strconv"
)

func DefineFuncs(program *lang.Program) {
	program.Funcs = make(map[string]*lang.DefaultFunc)
	program.Funcs["printf"] = DefinePrintFnFunction(program)
	program.Funcs["scanf"] = DefineScanFnFunction(program)
	program.Funcs["push"] = &lang.DefaultFunc{IrFunc: DefinePushFunction(program)}
	program.Funcs["pop"] = &lang.DefaultFunc{IrFunc: DefinePopFunction(program)}
	program.Funcs["add"] = &lang.DefaultFunc{IrFunc: DefineAddFunc(program)}
	program.Funcs["sub"] = &lang.DefaultFunc{IrFunc: DefineSubFunc(program)}
	program.Funcs["mul"] = &lang.DefaultFunc{IrFunc: DefineMulFunc(program)}
	program.Funcs["print"] = &lang.DefaultFunc{IrFunc: DefinePrintFunction(program)}
	program.Funcs["printI8"] = &lang.DefaultFunc{IrFunc: DefinePrintCharFunction(program)}
	program.Funcs["input"] = &lang.DefaultFunc{IrFunc: DefineInputFunction(program)}

	program.Funcs["dup"] = &lang.DefaultFunc{IrFunc: DefineDupFunc(program)}
	program.Funcs["swap"] = &lang.DefaultFunc{IrFunc: DefineSwapFunction(program)}
	program.Funcs["rot"] = &lang.DefaultFunc{IrFunc: DefineRotFunction(program)}
	program.Funcs["over"] = &lang.DefaultFunc{IrFunc: DefineOverFunction(program)}
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

func CallFunc(function *ir.Func, block *ir.Block, program *lang.Program, tok lang.Token) {
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

func LoadBlock(program *lang.Program, block *lang.Block) {
	irBlock := block.Func.Ir.NewBlock("block" + strconv.Itoa(block.Id))
	block.Ir = irBlock
	for i := 0; i < len(block.Tokens); i++ {
		tok := block.Tokens[i]
		switch tok.TokenType {
		case lang.PushT:
			irBlock.NewCall(program.Funcs["push"].IrFunc, constant.NewInt(types.I32, int64(tok.Value.(int))))
			break
		case lang.PopT:
			CallFunc(program.Funcs["pop"].IrFunc, irBlock, program, tok)
			break
		case lang.AddT:
			CallFunc(program.Funcs["add"].IrFunc, irBlock, program, tok)
			break
		case lang.SubT:
			CallFunc(program.Funcs["sub"].IrFunc, irBlock, program, tok)
			break
		case lang.MulT:
			CallFunc(program.Funcs["mul"].IrFunc, irBlock, program, tok)
			break
		case lang.PrintT:
			CallFunc(program.Funcs["print"].IrFunc, irBlock, program, tok)
			break
		case lang.PrintI8T:
			CallFunc(program.Funcs["printI8"].IrFunc, irBlock, program, tok)
			break
		case lang.InputT:
			CallFunc(program.Funcs["input"].IrFunc, irBlock, program, tok)
			break
		case lang.DupT:
			CallFunc(program.Funcs["dup"].IrFunc, irBlock, program, tok)
			break
		case lang.SwapT:
			CallFunc(program.Funcs["swap"].IrFunc, irBlock, program, tok)
			break
		case lang.RotT:
			CallFunc(program.Funcs["rot"].IrFunc, irBlock, program, tok)
			break
		case lang.OverT:
			CallFunc(program.Funcs["over"].IrFunc, irBlock, program, tok)
			break
		case lang.BlockT:
			LoadBlock(program, tok.Value.(*lang.Block))
		default:
			fmt.Println(tok)
			panic("unhandled default case")
		}
	}

}

func LoadFunction(program *lang.Program, fun *lang.Function) {
	funFn := program.Module.NewFunc(fun.Name, types.I32)
	funFunctionBlock := program.MainFunction.Blocks[0]
	funFunctionBlock.Func.Ir = funFn
	LoadBlock(program, funFunctionBlock)
	funFunctionBlock.Ir.NewRet(constant.NewInt(types.I32, 0))
}

func LoadProgram(program *lang.Program) {
	program.Globals = DefineGlobals(program.Module)
	DefineFuncs(program)

	LoadFunction(program, program.MainFunction)
	for _, fun := range program.Functions {
		if fun.Name != "main" {
			LoadFunction(program, fun)
		}
	}
}
