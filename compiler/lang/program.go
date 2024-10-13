package lang

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"strconv"
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
	LoopIndex    int
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

func (p *Program) NewBlockIndex() string {
	idx := p.BlockIndex
	p.BlockIndex++
	return "block" + strconv.Itoa(idx)
}
func (p *Program) NewLoopIndex() string {
	idx := p.LoopIndex
	p.LoopIndex++
	return "loop" + strconv.Itoa(idx)
}
