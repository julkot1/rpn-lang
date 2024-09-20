package lang

import "github.com/llir/llvm/ir"

type Function struct {
	Name   string
	Blocks []*Block
	Ir     *ir.Func
}

func NewFunction(name string) *Function {
	return &Function{Name: name}
}

func (f *Function) GetBlock(id string) *Block {
	for _, block := range f.Blocks {
		if block.Name == id {
			return block
		}
	}
	panic("Block not found")
	return nil
}
