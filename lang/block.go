package lang

import "github.com/llir/llvm/ir"

type Block struct {
	Id     int
	Func   *Function
	Ir     *ir.Block
	Parent *Block
	Tokens []Token
	Vars   []*Var
}

func NewBlock(id int) *Block {
	return &Block{Id: id}
}
func (block *Block) GetVars() []*Var {
	vars := make([]*Var, 0)
	if block.Vars != nil {
		vars = append(vars, block.Vars...)
	}
	if block.Parent != nil {
		vars = append(vars, block.Parent.GetVars()...)
	}
	return vars
}

func (block *Block) AddVar(variable *Var) {
	variable.Parent = block
	block.Vars = append(block.Vars, variable)
}

func (block *Block) ContainsVar(name string) bool {
	vars := block.GetVars()
	for _, v := range vars {
		if v.Name == name {
			return true
		}
	}
	return false

}
