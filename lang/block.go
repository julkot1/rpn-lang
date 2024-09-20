package lang

import (
	"github.com/llir/llvm/ir"
)

type Block struct {
	Name      string
	Func      *Function
	Ir        *ir.Block
	Parent    *Block
	Children  []*Block
	Tokens    []Token
	Vars      []*Var
	FreeBlock bool
}

func NewBlock(id string, freeBlock bool) *Block {
	return &Block{Name: id, FreeBlock: freeBlock}
}

func (block *Block) GetVars() []*Var {
	vars := make([]*Var, 0)
	if block.Vars != nil {
		vars = append(vars, block.Vars...)
	}
	if block.Parent != nil {
		vars = append(vars, block.Parent.GetVars()...)
	}
	return vars
}

func (block *Block) AddVar(variable *Var) {
	variable.Parent = block
	block.Vars = append(block.Vars, variable)
}

func (block *Block) ContainsVar(name string) bool {
	vars := block.GetVars()
	for _, v := range vars {
		if v.Name == name {
			return true
		}
	}
	return false

}

func containsParent(forbidden []*Block, block *Block) bool {
	for _, v := range forbidden {
		if v.Parent == block {
			return true
		}
	}
	return false
}
func containsSelf(forbidden []*Block, block *Block) bool {
	for _, v := range forbidden {
		if v == block {
			return true
		}
	}
	return false
}
func (block *Block) NextFreeBlock(forbidden ...*Block) *Block {
	children := block.Children
	afterBlock := true
	if containsParent(forbidden, block) {
		afterBlock = false
	}
	for _, child := range children {
		if containsSelf(forbidden, child) {
			afterBlock = true
			continue
		}
		if child.FreeBlock && afterBlock {
			return child
		}
	}
	// if it has no free block as children find it in parent
	return block.Parent.NextFreeBlock(forbidden...)

}
