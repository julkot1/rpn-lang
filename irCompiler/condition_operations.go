package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
	"strconv"
)

func LoadIf(program *lang.Program, statement *lang.IfStatement, fun *lang.Function, block *ir.Block, nextBlock *lang.Block) *lang.Block {
	args := GetValues(program, 1, block)
	condition := block.NewICmp(enum.IPredEQ, args[0], constant.NewInt(types.I32, 1))
	statement.TrueBlock.Ir = ir.NewBlock(statement.TrueBlock.Name)
	nextBlock.Ir = ir.NewBlock(nextBlock.Name)
	block.NewCondBr(condition, statement.TrueBlock.Ir, nextBlock.Ir)
	LoadBlock(program, statement.TrueBlock, fun)
	statement.TrueBlock.Ir.NewBr(nextBlock.Ir)
	return nextBlock
}

func LoadRepeat(program *lang.Program, statement *lang.RepeatStatement, fun *lang.Function, block *ir.Block, nextBlock *lang.Block) *lang.Block {
	arg := GetValues(program, 1, block)[0]
	counter := block.NewAlloca(types.I32)

	block.NewStore(constant.NewInt(types.I32, 0), counter)

	conditionName := "loop" + strconv.Itoa(program.NewBlockIndex())
	conditionBlock := fun.Ir.NewBlock(conditionName)

	block.NewBr(conditionBlock)

	counterVal := conditionBlock.NewLoad(types.I32, counter)
	compare := conditionBlock.NewICmp(enum.IPredNE, arg, counterVal)

	statement.LoopBlock.Ir = fun.Ir.NewBlock(statement.LoopBlock.Name)
	nextBlock.Ir = fun.Ir.NewBlock(nextBlock.Name)

	conditionBlock.NewCondBr(compare, statement.LoopBlock.Ir, nextBlock.Ir)
	LoadBlock(program, statement.LoopBlock, fun)

	counterValBody := conditionBlock.NewLoad(types.I32, counter)
	newCounter := statement.LoopBlock.Ir.NewAdd(constant.NewInt(types.I32, 1), counterValBody)
	statement.LoopBlock.Ir.NewStore(newCounter, counter)
	statement.LoopBlock.Ir.NewBr(conditionBlock)

	return nextBlock
}
