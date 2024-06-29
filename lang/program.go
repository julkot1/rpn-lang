package lang

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type DefaultFunc struct {
	IrFunc *ir.Func
	Type   *types.FuncType
}

type Program struct {
	Functions    []*Function
	MainFunction *Function
	Module       *ir.Module
	Globals      map[string]*ir.Global
	Funcs        map[string]*DefaultFunc
	BlockIndex   int
}

func DefineTypes() {
	types.I8Ptr = types.NewPointer(types.I8)
}

func NewProgram() *Program {
	program := &Program{}
	program.Module = ir.NewModule()
	DefineTypes()

	return program
}

func (p *Program) NewBlockIndex() int {
	idx := p.BlockIndex
	p.BlockIndex++
	return idx
}
