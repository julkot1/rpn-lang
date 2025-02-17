package lang

import "github.com/llir/llvm/ir"

type Var struct {
	Name        string
	Type        Type
	ComplexType string
	Ir          *ir.InstAlloca
}
