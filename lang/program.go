package lang

type Program struct {
	Functions    []*Function
	MainFunction *Function

	BlockIndex int
}

func NewProgram() *Program {
	return &Program{nil, nil, 0}
}

func (p *Program) NewIndex() int {
	idx := p.BlockIndex
	p.BlockIndex++
	return idx
}
