package lang

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type FuncParam struct {
	Name        string
	Type        Type
	ComplexType string
	Size        int64
	Ir          *ir.Param
}
type Function struct {
	Name   string
	Blocks []*Block
	Ir     *ir.Func
	Params []*FuncParam
}

func NewFunction(name string, module *ir.Module, params []*FuncParam) *Function {
	x := &Function{Name: name}
	x.Params = params
	irParams := make([]*ir.Param, len(params))
	for i, p := range params {
		irParams[i] = p.Ir
	}
	x.Ir = module.NewFunc(name, types.I64, irParams...)
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

func (f *Function) GetParam(text string) (*FuncParam, *ir.Param) {
	var value *FuncParam
	var valueTyp *ir.Param
	for _, p := range f.Params {
		if p.Name == text {
			value = p
		} else if p.Name == (text + ".typ") {
			valueTyp = p.Ir
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
