package lang

type Function struct {
	Name   string
	Blocks []*Block
}

func NewFunction(name string) *Function {
	return &Function{Name: name}
}
