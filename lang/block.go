package lang

type Block struct {
	Func   *Function
	Parent *Block
	Tokens []Token
	Vars   []*Var
}

func NewBlock(fn *Function, parent *Block) *Block {
	return &Block{Func: fn, Parent: parent}
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
