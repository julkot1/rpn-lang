package irCompiler

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"math"
	"rpn/lang"
	"rpn/lexer"
)

func DefineFuncs(program *lang.Program) {

	program.Funcs[lang.PrintFT] = DefinePrintFnFunction(program)
	program.Funcs[lang.ScanFT] = DefineScanFnFunction(program)

	program.Funcs[lang.PushT] = &lang.DefaultFunc{IrFunc: DefinePushFunction(program)}
	program.Funcs[lang.PopT] = &lang.DefaultFunc{IrFunc: DefinePopFunction(program)}
	program.Funcs[lang.PopTypeT] = &lang.DefaultFunc{IrFunc: DefineTypePopFunction(program)}

	program.Funcs[lang.DupT] = &lang.DefaultFunc{IrFunc: DefineDupFunc(program)}
	program.Funcs[lang.SwapT] = &lang.DefaultFunc{IrFunc: DefineSwapFunction(program)}
	program.Funcs[lang.RotT] = &lang.DefaultFunc{IrFunc: DefineRotFunction(program)}
	program.Funcs[lang.OverT] = &lang.DefaultFunc{IrFunc: DefineOverFunction(program)}

	program.Funcs[lang.InputT] = &lang.DefaultFunc{IrFunc: DefineInputFunction(program)}

	program.Funcs[lang.TypeofT] = &lang.DefaultFunc{IrFunc: DefineTypeofFunction(program)}

}

func DefineGlobals(m *ir.Module) map[string]*ir.Global {
	stackSize := 100
	stackType := types.NewArray(uint64(stackSize), types.I64)

	globals := make(map[string]*ir.Global)

	globals["stack"] = m.NewGlobalDef("vstack", constant.NewZeroInitializer(stackType))
	globals["type_stack"] = m.NewGlobalDef("tstack", constant.NewZeroInitializer(stackType))
	globals["top"] = m.NewGlobalDef("top", constant.NewInt(types.I64, 0))

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
		argc := len(function.Sig.Params) / 2
		if argc == 0 {
			block.NewCall(function)
		} else {
			args := GetPreventValues(program, argc, block)
			block.NewCall(function, args...)
		}
		return
	}
	argc := len(function.Sig.Params) / 2
	if argc == 0 {
		block.NewCall(function)
	} else {
		args := GetValues(program, argc, block)
		block.NewCall(function, args...)
	}

}

func LoadBlock(program *lang.Program, block *lang.Block, fun *lang.Function, idx int) {
	irBlock := block.Ir

	for i := 0; i < len(block.Tokens); i++ {
		tok := block.Tokens[i]
		if tok.TokenType == lang.PushT {
			pushableTok := tok.Value.(lang.PushableToken)
			pushToken(irBlock, pushableTok, program)
		} else if lexer.ConstTokens[tok.TokenType].DefaultFunction {
			CallFunc(program.Funcs[tok.TokenType].IrFunc, irBlock, program, tok)
		} else if tok.TokenType == lang.IdentifierT {
			if program.GlobalTokenTable[tok.Value.(string)] == lang.PStaticFunc {
				CallFunc(program.StaticFunctions[tok.Value.(string)].IrFunc, irBlock, program, tok)
			}
		} else {
			switch tok.TokenType {

			case lang.IfT:
				LoadIf(program, irBlock, block.NextFreeBlock(block), fun.Blocks[idx+1])
				break
			case lang.RepeatT:
				LoadRepeat(program, fun, irBlock, block.NextFreeBlock(block, tok.Value.(*lang.RepeatStatement).LoopBlock), fun.Blocks[idx+1], tok.Value.(*lang.RepeatStatement).LoopBlock)
				break
			default:
				fmt.Println(tok)
				panic("unhandled default case")
			}

		}
	}
}

func pushToken(block *ir.Block, tok lang.PushableToken, program *lang.Program) {
	var val value.Value
	switch tok.Typ {
	case lang.INT_T:
		val = constant.NewInt(types.I64, tok.Value.(int64))
		break
	case lang.CHAR_T:
		val = constant.NewInt(types.I64, tok.Value.(int64))
		break
	case lang.FLOAT_T:
		f := math.Float64bits(tok.Value.(float64))
		val = constant.NewInt(types.I64, int64(f))
		break

	}
	block.NewCall(program.Funcs[lang.PushT].IrFunc,
		val,
		constant.NewInt(types.I64, int64(tok.Typ))) // TO-DO type
}

func CreateIrBlocks(fun *lang.Function) {

	for _, block := range fun.Blocks {
		block.Ir = fun.Ir.NewBlock(block.Name)
		for _, tok := range block.Tokens {
			if tok.TokenType == lang.RepeatT {
				tok.Value.(*lang.RepeatStatement).LoopBlock.Ir = fun.Ir.NewBlock(tok.Value.(*lang.RepeatStatement).LoopBlock.Name)
			}
		}
	}
}

func LoadFunction(program *lang.Program, fun *lang.Function) {
	funFn := program.Module.NewFunc(fun.Name, types.I64)
	fun.Ir = funFn
	CreateIrBlocks(fun)

	for idx, block := range fun.Blocks {
		LoadBlock(program, block, fun, idx)
	}
	for idx, block := range fun.Blocks {
		if block.Ir.Term == nil {
			if idx+1 < len(fun.Blocks) {
				block.Ir.NewBr(block.NextFreeBlock(block).Ir)
			} else {
				block.Ir.NewRet(constant.NewInt(types.I64, 0))
			}
		}
	}
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
