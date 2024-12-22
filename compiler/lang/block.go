package lang

import (
	"github.com/llir/llvm/ir"
)

type Block struct {
	Name          string
	Func          *Function
	Ir            *ir.Block
	Vars          []*Var
	FreeBlock     bool
	CreateIf      bool
	CreateLoop    bool
	LoopCondition bool
}

func NewBlockIr(id string, freeBlock bool, fun *Function, createIf, createLoop bool) *Block {
	x := &Block{Name: id, FreeBlock: freeBlock, CreateIf: createIf, CreateLoop: createLoop}
	x.Ir = fun.Ir.NewBlock(id)
	x.LoopCondition = false
	return x
}
func NewBlock(id string, freeBlock bool) *Block {
	x := &Block{Name: id, FreeBlock: freeBlock}
	x.LoopCondition = false
	return x
}

func (block *Block) AddVar(variable *Var) {
	variable.Parent = block
	block.Vars = append(block.Vars, variable)
}
