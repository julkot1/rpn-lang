package lang

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Function struct {
	Name   string
	Blocks []*Block
	Ir     *ir.Func
}

func NewFunction(name string, module *ir.Module) *Function {
	x := &Function{Name: name}
	x.Ir = module.NewFunc(name, types.I64)
	return x
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
