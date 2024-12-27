package lang

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Function struct {
	Name   string
	Blocks []*Block
	Ir     *ir.Func
	params []*ir.Param
}

func (f *Function) Params(params []string) {
	f.params = make([]*ir.Param, 0)
	for _, param := range params {
		paramIr := ir.NewParam(param, types.I64)
		f.params = append(f.params, paramIr)
	}
}

func NewFunction(name string, module *ir.Module, params []string) *Function {
	x := &Function{Name: name}
	x.Params(params)
	x.Ir = module.NewFunc(name, types.I64, x.params...)
	return x
}

func (f *Function) GetBlock(id string) *Block {
	for _, block := range f.Blocks {
		if block.Name == id {
			return block
		}
	}
	panic("Block not found")
	return nil
}

func (f *Function) IsParameter(text string) bool {
	paramsNames := getParamsNames(f.Ir.Params)
	if len(paramsNames) == 0 {
		return false
	}
	for _, name := range paramsNames {
		if name == text {
			return true
		}
	}
	return false
}

func (f *Function) GetParam(text string) (*ir.Param, *ir.Param) {
	var value *ir.Param
	var valueTyp *ir.Param
	for _, p := range f.params {
		if p.Name() == text {
			value = p
		} else if p.Name() == (text + ".typ") {
			valueTyp = p
		}
		if value != nil && valueTyp != nil {
			return value, valueTyp
		}
	}
	return nil, nil
}

func getParamsNames(params []*ir.Param) []string {
	result := make([]string, 0)

	for _, param := range params {
		result = append(result, param.Name())
	}
	return result
}
