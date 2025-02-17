package lang

import "github.com/llir/llvm/ir"

type Var struct {
	Name string
	Type Type
	Ir   *ir.InstAlloca
}
