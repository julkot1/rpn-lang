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
type ProgramTokenType int

const (
	PConst ProgramTokenType = iota
	PStcFunction
	PDefaultFunc
	PStaticFunc
)

type Program struct {
	GlobalTokenTable  map[string]ProgramTokenType
	Functions         []*Function
	MainFunction      *Function
	Module            *ir.Module
	StaticLibsModules []*ir.Module
	StaticFunctions   map[string]*DefaultFunc
	Globals           map[string]*ir.Global
	Funcs             map[TokenType]*DefaultFunc
	BlockIndex        int
	LoopIndex         int
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
