package irCompiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
)

func DefineFuncs(program *lang.Program) {

	program.Funcs[lang.PushT] = &lang.DefaultFunc{IrFunc: DefinePushFunction(program)}
	program.Funcs[lang.PopT] = &lang.DefaultFunc{IrFunc: DefinePopFunction(program)}
	program.Funcs[lang.PopTypeT] = &lang.DefaultFunc{IrFunc: DefineTypePopFunction(program)}

	program.Funcs[lang.DupT] = &lang.DefaultFunc{IrFunc: DefineDupFunc(program)}
	program.Funcs[lang.SwapT] = &lang.DefaultFunc{IrFunc: DefineSwapFunction(program)}
	program.Funcs[lang.RotT] = &lang.DefaultFunc{IrFunc: DefineRotFunction(program)}
	program.Funcs[lang.OverT] = &lang.DefaultFunc{IrFunc: DefineOverFunction(program)}

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

func DefineStructs() map[string]*types.StructType {
	structs := make(map[string]*types.StructType)

	variable := types.NewStruct(types.I64, types.I64)
	structs["variable"] = variable

	array := types.NewStruct(types.I64, types.I64, types.NewPointer(variable))
	structs["array"] = array

	return structs
}

func CallFunc(function *ir.Func, block *ir.Block, program *lang.Program, prevent bool) {
	if prevent == true {
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

func LoadProgram(program *lang.Program) {
	program.Globals = DefineGlobals(program.Module)
	program.Structs = DefineStructs()

	DefineFuncs(program)

	treeWalk := NewTreeWalk()
	treeWalk.program = program
	antlr.ParseTreeWalkerDefault.Walk(treeWalk, program.Tree)

	createJumps(program)
}
func createJumps(program *lang.Program) {
	for _, fun := range program.Functions {
		for idx, block := range fun.Blocks {
			if block.CreateIf {
				LoadIf(program, block.Ir, NextFreeBlock(fun, idx, program), fun.Blocks[idx+1])
			} else if block.CreateLoop {
				i, b := MatchLoopCondition(fun, idx, program)
				LoadRepeat(program, block, NextLoopFreeBlock(fun, i, program), fun.Blocks[idx+1], b)
			} else if block.Ir.Term == nil {
				block.Ir.NewBr(NextFreeBlock(fun, idx, program).Ir)
			}

		}
	}
}
func NextLoopFreeBlock(fun *lang.Function, idx int, program *lang.Program) *lang.Block {
	if idx+1 >= len(fun.Blocks) {
		b := lang.NewBlockIr(program.NewBlockIndex(), true, fun, false, false)
		b.Ir.NewRet(constant.NewInt(types.I64, 0))
		fun.Blocks = append(fun.Blocks, b)
		return b
	}
	return fun.Blocks[idx+1]
}
func MatchLoopCondition(fun *lang.Function, idx int, program *lang.Program) (int, *lang.Block) {
	counter := 0
	for i := idx; i < len(fun.Blocks); i++ {
		if fun.Blocks[i].LoopCondition {
			counter++
		} else if fun.Blocks[i].CreateLoop {
			counter--
		}
		if counter == 0 {
			return i, fun.Blocks[i]
		}
	}
	b := lang.NewBlockIr(program.NewBlockIndex(), true, fun, false, false)
	b.Ir.NewRet(constant.NewInt(types.I64, 0))
	fun.Blocks = append(fun.Blocks, b)
	return len(fun.Blocks) - 1, b
}

func NextFreeBlock(fun *lang.Function, idx int, program *lang.Program) *lang.Block {
	counter := 0
	endValue := 0
	if fun.Blocks[idx].FreeBlock && !fun.Blocks[idx].CreateIf {
		endValue = 1
	}
	if len(fun.Blocks) >= idx-1 && idx-1 >= 0 {
		if fun.Blocks[idx-1].CreateLoop && fun.Blocks[idx].CreateIf {
			endValue = 0
		}
	}

	for i := idx + 1; i < len(fun.Blocks); i++ {
		if fun.Blocks[i].FreeBlock {
			counter++
		} else {
			counter--
		}
		if counter == endValue {
			return fun.Blocks[i]
		}
	}
	b := lang.NewBlockIr(program.NewBlockIndex(), true, fun, false, false)
	b.Ir.NewRet(constant.NewInt(types.I64, 0))
	fun.Blocks = append(fun.Blocks, b)
	return b
}
