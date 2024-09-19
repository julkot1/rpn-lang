package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
	"strconv"
)

func LoadIf(program *lang.Program, block *ir.Block, nextBlock *lang.Block, trueBlock *lang.Block) *lang.Block {
	args := GetValues(program, 1, block)
	condition := block.NewICmp(enum.IPredEQ, args[0], constant.NewInt(types.I32, 1))

	block.NewCondBr(condition, trueBlock.Ir, nextBlock.Ir)
	trueBlock.Ir.NewBr(nextBlock.Ir)
	return nextBlock
}

func LoadRepeat(program *lang.Program, fun *lang.Function, block *ir.Block, nextBlock *lang.Block, loopBlock *lang.Block) *lang.Block {
	arg := GetValues(program, 1, block)[0]
	counter := block.NewAlloca(types.I32)

	block.NewStore(constant.NewInt(types.I32, 0), counter)

	conditionName := "loop" + strconv.Itoa(program.NewBlockIndex())
	conditionBlock := fun.Ir.NewBlock(conditionName)

	block.NewBr(conditionBlock)

	counterVal := conditionBlock.NewLoad(types.I32, counter)
	compare := conditionBlock.NewICmp(enum.IPredNE, arg, counterVal)

	conditionBlock.NewCondBr(compare, loopBlock.Ir, nextBlock.Ir)

	counterValBody := loopBlock.Ir.NewLoad(types.I32, counter)
	newCounter := loopBlock.Ir.NewAdd(constant.NewInt(types.I32, 1), counterValBody)
	loopBlock.Ir.NewStore(newCounter, counter)
	loopBlock.Ir.NewBr(conditionBlock)
	return nextBlock
}
