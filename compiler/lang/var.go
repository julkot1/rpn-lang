package lang

import "github.com/llir/llvm/ir"

type Var struct {
	Name        string
	Type        Type
	ComplexType string
	Size        int64
	Ir          *ir.InstAlloca
}
