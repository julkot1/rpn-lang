package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
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
