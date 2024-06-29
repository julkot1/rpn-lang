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
