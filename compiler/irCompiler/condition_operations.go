package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
)

func LoadIf(program *lang.Program, block *ir.Block, nextBlock *lang.Block, trueBlock *lang.Block) *lang.Block {
	args := GetValues(program, 1, block)
	condition := block.NewICmp(enum.IPredEQ, args[0], constant.NewInt(types.I64, 1))

	block.NewCondBr(condition, trueBlock.Ir, nextBlock.Ir)
	trueBlock.Ir.NewBr(nextBlock.Ir)
	return nextBlock
}

func LoadRepeat(program *lang.Program, block *lang.Block, nextBlock, loopBlock, conditionBlock *lang.Block) *lang.Block {

	arg := GetValues(program, 1, block.Ir)[0]
	counter, ok := block.Vars["loop_index"]
	var loadPtr *ir.InstGetElementPtr
	if !ok {
		counter = block.Ir.NewAlloca(types.I64)
		block.Ir.NewStore(constant.NewInt(types.I64, -1), counter)
	} else {
		typ := program.Structs["variable"]
		loadPtr = block.Ir.NewGetElementPtr(typ, counter, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	}

	conditionBlockIr := conditionBlock.Ir

	block.Ir.NewBr(conditionBlockIr)

	var counterValBody *ir.InstLoad
	if ok {
		counterValBody = conditionBlock.Ir.NewLoad(types.I64, loadPtr)
	} else {
		counterValBody = conditionBlock.Ir.NewLoad(types.I64, counter)
	}
	newCounter := conditionBlock.Ir.NewAdd(constant.NewInt(types.I64, 1), counterValBody)
	if !ok {
		conditionBlock.Ir.NewStore(newCounter, counter)
	} else {
		conditionBlock.Ir.NewStore(newCounter, loadPtr)
	}

	var counterVal *ir.InstLoad
	if ok {
		counterVal = conditionBlockIr.NewLoad(types.I64, loadPtr)
	} else {
		counterVal = conditionBlockIr.NewLoad(types.I64, counter)
	}
	compare := conditionBlockIr.NewICmp(enum.IPredNE, arg, counterVal)

	conditionBlockIr.NewCondBr(compare, loopBlock.Ir, nextBlock.Ir)

	loopBlock.Ir.NewBr(conditionBlockIr)
	return nextBlock
}
