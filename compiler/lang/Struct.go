package lang

import "github.com/llir/llvm/ir/types"

type StcStruct struct {
	Name   string
	IrType types.Type
	Args   []StructElement
}

type StructElement struct {
	Name string
	Type Type
}
