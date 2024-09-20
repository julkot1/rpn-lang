package irCompiler

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
	"rpn/lexer"
)

func DefineFuncs(program *lang.Program) {
	program.Funcs = make(map[string]*lang.DefaultFunc)
	program.Funcs["printf"] = DefinePrintFnFunction(program)
	program.Funcs["scanf"] = DefineScanFnFunction(program)

	program.Funcs["push"] = &lang.DefaultFunc{IrFunc: DefinePushFunction(program)}
	program.Funcs["pop"] = &lang.DefaultFunc{IrFunc: DefinePopFunction(program)}
	program.Funcs["add"] = &lang.DefaultFunc{IrFunc: DefineAddFunc(program)}
	program.Funcs["mod"] = &lang.DefaultFunc{IrFunc: DefineModFunc(program)}
	program.Funcs["sub"] = &lang.DefaultFunc{IrFunc: DefineSubFunc(program)}
	program.Funcs["mul"] = &lang.DefaultFunc{IrFunc: DefineMulFunc(program)}

	program.Funcs["print"] = &lang.DefaultFunc{IrFunc: DefinePrintFunction(program)}
	program.Funcs["printI8"] = &lang.DefaultFunc{IrFunc: DefinePrintCharFunction(program)}
	program.Funcs["input"] = &lang.DefaultFunc{IrFunc: DefineInputFunction(program)}

	program.Funcs["dup"] = &lang.DefaultFunc{IrFunc: DefineDupFunc(program)}
	program.Funcs["swap"] = &lang.DefaultFunc{IrFunc: DefineSwapFunction(program)}
	program.Funcs["rot"] = &lang.DefaultFunc{IrFunc: DefineRotFunction(program)}
	program.Funcs["over"] = &lang.DefaultFunc{IrFunc: DefineOverFunction(program)}

	program.Funcs["equals"] = &lang.DefaultFunc{IrFunc: DefineEqualsFunction(program)}
	program.Funcs["not_equals"] = &lang.DefaultFunc{IrFunc: DefineNotEqualsFunction(program)}
	program.Funcs["greater"] = &lang.DefaultFunc{IrFunc: DefineGreaterFunction(program)}
	program.Funcs["less"] = &lang.DefaultFunc{IrFunc: DefineLessFunction(program)}
	program.Funcs["greater_eq"] = &lang.DefaultFunc{IrFunc: DefineGreaterOrEqFunction(program)}
	program.Funcs["less_eq"] = &lang.DefaultFunc{IrFunc: DefineLessOrEqFunction(program)}

	program.Funcs["not"] = &lang.DefaultFunc{IrFunc: DefineNotOperation(program)}
	program.Funcs["and"] = &lang.DefaultFunc{IrFunc: DefineAndOperation(program)}
	program.Funcs["or"] = &lang.DefaultFunc{IrFunc: DefineOrOperation(program)}

	AssignToConstTokens(program.Funcs)
}

func AssignToConstTokens(funcs map[string]*lang.DefaultFunc) {
	lexer.ConstTokens[lang.PushT].Ir = funcs["push"].IrFunc
	lexer.ConstTokens[lang.PopT].Ir = funcs["pop"].IrFunc
	lexer.ConstTokens[lang.AddT].Ir = funcs["add"].IrFunc
	lexer.ConstTokens[lang.SubT].Ir = funcs["sub"].IrFunc
	lexer.ConstTokens[lang.MulT].Ir = funcs["mul"].IrFunc
	lexer.ConstTokens[lang.ModT].Ir = funcs["mod"].IrFunc

	lexer.ConstTokens[lang.PrintT].Ir = funcs["print"].IrFunc
	lexer.ConstTokens[lang.PrintI8T].Ir = funcs["printI8"].IrFunc
	lexer.ConstTokens[lang.InputT].Ir = funcs["input"].IrFunc

	lexer.ConstTokens[lang.DupT].Ir = funcs["dup"].IrFunc
	lexer.ConstTokens[lang.SwapT].Ir = funcs["swap"].IrFunc
	lexer.ConstTokens[lang.RotT].Ir = funcs["rot"].IrFunc
	lexer.ConstTokens[lang.OverT].Ir = funcs["over"].IrFunc

	lexer.ConstTokens[lang.EqualsT].Ir = funcs["equals"].IrFunc
	lexer.ConstTokens[lang.NotEqualT].Ir = funcs["not_equals"].IrFunc
	lexer.ConstTokens[lang.GreaterT].Ir = funcs["greater"].IrFunc
	lexer.ConstTokens[lang.LessT].Ir = funcs["less"].IrFunc
	lexer.ConstTokens[lang.GreaterOrEqT].Ir = funcs["greater_eq"].IrFunc
	lexer.ConstTokens[lang.LessOrEqT].Ir = funcs["less_eq"].IrFunc

	lexer.ConstTokens[lang.NotT].Ir = funcs["not"].IrFunc
	lexer.ConstTokens[lang.OrT].Ir = funcs["or"].IrFunc
	lexer.ConstTokens[lang.AndT].Ir = funcs["and"].IrFunc
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
		argc := lexer.ConstTokens[tok.TokenType].Argc
		if argc == 0 {
			block.NewCall(function)
		} else {
			args := GetPreventValues(program, argc, block)
			if argc == 1 {
				block.NewCall(function, args[0])
			}
			if argc == 2 {
				block.NewCall(function, args[1], args[0])
			}
		}
	} else {
		argc := lexer.ConstTokens[tok.TokenType].Argc
		if argc == 0 {
			block.NewCall(function)
		} else {
			args := GetValues(program, argc, block)
			if argc == 1 {
				block.NewCall(function, args[0])
			}
			if argc == 2 {
				block.NewCall(function, args[1], args[0])
			}
		}
	}

}

func LoadBlock(program *lang.Program, block *lang.Block, fun *lang.Function, idx int) {
	irBlock := block.Ir

	for i := 0; i < len(block.Tokens); i++ {
		tok := block.Tokens[i]
		if tok.TokenType == lang.PushT {
			irBlock.NewCall(program.Funcs["push"].IrFunc, constant.NewInt(types.I32, int64(tok.Value.(int))))
		} else if lexer.ConstTokens[tok.TokenType].DefaultFunction {
			CallFunc(lexer.ConstTokens[tok.TokenType].Ir, irBlock, program, tok)
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
	funFn := program.Module.NewFunc(fun.Name, types.I32)
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
				block.Ir.NewRet(constant.NewInt(types.I32, 0))
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
