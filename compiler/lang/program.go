package lang

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
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
	PVar
)

type Program struct {
	Tree   antlr.Tree
	Stream *antlr.CommonTokenStream

	StringTable       *StringTable
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
	Structs           map[string]*types.StructType
}

func (p *Program) GetFunction(name string, function *Function) *Function {
	if function.Name == name {
		return function
	}
	for _, fun := range p.Functions {
		if fun.Name == name {
			return fun
		}
	}
	return nil
}

func DefineTypes() {
	types.I8Ptr = types.NewPointer(types.I8)
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

func NewProgram() *Program {
	program := &Program{}
	program.Module = ir.NewModule()
	program.StaticFunctions = make(map[string]*DefaultFunc)
	program.GlobalTokenTable = make(map[string]ProgramTokenType)
	program.Funcs = make(map[TokenType]*DefaultFunc)
	program.StringTable = NewStringTable()
	program.StaticLibsModules = make([]*ir.Module, 0)
	return program

}

func (p *Program) AddGlobalToken(name string, t ProgramTokenType) error {
	if _, ok := p.GlobalTokenTable[name]; ok {
		return fmt.Errorf("global token already defined: %s", name)
	}
	p.GlobalTokenTable[name] = t
	return nil
}
