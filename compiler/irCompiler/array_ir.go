package irCompiler

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"math"
	"os"
	"rpn/lang"
	"rpn/parser"
	"rpn/util"
	"strconv"
)

func CreateArray(ctx *parser.ArrayContext, block *lang.Block, program *lang.Program) {
	length, capacity := getSize(ctx)
	arrType := program.Structs["array"]
	varType := program.Structs["variable"]

	lengthConst := constant.NewInt(types.I64, length)
	capacityConst := constant.NewInt(types.I64, capacity)
	arrElemType := types.NewArray(uint64(capacity), varType)

	elems := getElems(varType, length, capacity, ctx, program)
	idx := program.NewArrayIndex()
	ar := program.Module.NewGlobalDef("struct.arr.elements."+idx, constant.NewArray(
		arrElemType, elems...,
	))

	zeroIndex := constant.NewInt(types.I32, 0)
	x := program.Module.NewGlobalDef("struct.arr."+idx, constant.NewStruct(
		arrType,
		lengthConst,
		capacityConst,
		constant.NewGetElementPtr(arrElemType, ar, zeroIndex, zeroIndex),
	))

	block.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
		block.Ir.NewPtrToInt(x, types.I64),
		constant.NewInt(types.I64, int64(lang.ARRAY_T)))
}

func getToken(ctx parser.IArrayElementContext, varType *types.StructType, program *lang.Program) *constant.Struct {
	var val constant.Constant
	var typ lang.Type

	if ctx.NUMBER() != nil {
		typ = lang.INT_T
		i, _ := strconv.Atoi(ctx.NUMBER().GetText())
		val = constant.NewInt(types.I64, int64(i))
	}
	if ctx.SIGNED_NUMBER() != nil {
		typ = lang.INT_T
		i, _ := strconv.Atoi(ctx.SIGNED_NUMBER().GetText())
		val = constant.NewInt(types.I64, int64(i))
	}
	if ctx.HEX_NUMBER() != nil {
		typ = lang.INT_T
		i, _ := strconv.ParseInt(ctx.HEX_NUMBER().GetText()[2:], 16, 64)
		val = constant.NewInt(types.I64, int64(i))
	}
	if ctx.BIN_NUMBER() != nil {
		typ = lang.INT_T
		i, _ := strconv.ParseInt(ctx.BIN_NUMBER().GetText()[2:], 2, 64)
		val = constant.NewInt(types.I64, int64(i))
	}
	if ctx.FLOAT() != nil {
		typ = lang.FLOAT_T
		i, _ := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
		f := math.Float64bits(i)
		val = constant.NewInt(types.I64, int64(f))
	}
	if ctx.FLOAT_E() != nil {
		typ = lang.FLOAT_T
		i, _ := strconv.ParseFloat(ctx.FLOAT_E().GetText(), 64)
		f := math.Float64bits(i)
		val = constant.NewInt(types.I64, int64(f))
	}
	if ctx.SIGNED_FLOAT() != nil {
		typ = lang.FLOAT_T
		i, _ := strconv.ParseFloat(ctx.SIGNED_FLOAT().GetText(), 64)
		f := math.Float64bits(i)
		val = constant.NewInt(types.I64, int64(f))
	}
	if ctx.SIGNED_FLOAT_E() != nil {
		typ = lang.FLOAT_T
		i, _ := strconv.ParseFloat(ctx.SIGNED_FLOAT_E().GetText(), 64)
		f := math.Float64bits(i)
		val = constant.NewInt(types.I64, int64(f))
	}

	if ctx.CHAR() != nil {
		typ = lang.CHAR_T
		text := ctx.CHAR().GetText()
		ch := getChar(text)
		val = constant.NewInt(types.I64, int64(ch))
	}
	if ctx.STRING() != nil {
		typ = lang.STRING_T
		text := ctx.STRING().GetText()
		str := trimString(text)
		g := program.StringTable.GetOrAdd(str, program)
		strPtr := constant.NewGetElementPtr(types.NewArray(uint64(len(str)+1), types.I8), g, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
		val = constant.NewPtrToInt(strPtr, types.I64)
	}

	if ctx.Array() != nil {
		typ = lang.ARRAY_T
		ctxArr := ctx.Array()
		arrConst, _ := createSubArray(ctxArr, program)
		val = constant.NewPtrToInt(arrConst, types.I64)
	}
	return constant.NewStruct(varType, val, constant.NewInt(types.I64, int64(typ)))

}

func createSubArray(ctx parser.IArrayContext, program *lang.Program) (*ir.Global, int64) {
	length, capacity := getSize(ctx.(*parser.ArrayContext))
	arrType := program.Structs["array"]
	varType := program.Structs["variable"]

	lengthConst := constant.NewInt(types.I64, length)
	capacityConst := constant.NewInt(types.I64, capacity)
	arrElemType := types.NewArray(uint64(capacity), varType)

	elems := getElems(varType, length, capacity, ctx.(*parser.ArrayContext), program)
	idx := program.NewArrayIndex()
	ar := program.Module.NewGlobalDef("struct.arr.elements."+idx, constant.NewArray(
		arrElemType, elems...,
	))

	zeroIndex := constant.NewInt(types.I32, 0)
	x := program.Module.NewGlobalDef("struct.arr."+idx, constant.NewStruct(
		arrType,
		lengthConst,
		capacityConst,
		constant.NewGetElementPtr(arrElemType, ar, zeroIndex, zeroIndex),
	))
	return x, capacity
}

func getElems(varType *types.StructType, length int64, capacity int64, ctx *parser.ArrayContext, program *lang.Program) []constant.Constant {
	elems := make([]constant.Constant, 0)

	for _, e := range ctx.AllArrayElement() {
		elem := getToken(e, varType, program)
		elems = append(elems, elem)
	}
	for i := 0; i < int(capacity-length); i++ {
		elem := constant.NewStruct(varType, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0))
		elems = append(elems, elem)
	}
	return elems
}

func getSize(ctx *parser.ArrayContext) (int64, int64) {
	length, capacity := len(ctx.AllArrayElement()), 0
	if ctx.Capacity() != nil {
		capacity, _ = strconv.Atoi(ctx.Capacity().NUMBER().GetText())
	} else {
		capacity = length
	}
	return int64(length), int64(capacity)
}

func getElementAtIndex(block *lang.Block, scope util.Stack, program *lang.Program, base string, index string, function *lang.Function) {
	arr, arrT := GetVar(base, program, scope, block)

	fun := program.Funcs[lang.AtT].IrFunc

	indexValue, indexType := getIndex(block, index, program, scope, function)
	if arr != nil {

		block.Ir.NewCall(fun, indexValue, arr, indexType, arrT)
		return
	}
	arrValue, arrTypeValue := function.GetParam(base)
	if arrValue == nil {
		fmt.Printf("variable '%s' does not exist:\n \tline\n", index)
		os.Exit(1)
	}
	block.Ir.NewCall(fun, indexValue, arrValue.Ir, indexType, arrTypeValue)

}

func getIndex(block *lang.Block, index string, program *lang.Program, scope util.Stack, function *lang.Function) (value.Value, value.Value) {
	val, err := strconv.Atoi(index)
	if err != nil {
		load, typeT := GetVar(index, program, scope, block)

		if load != nil {
			return load, typeT
		} else {
			loadValue, typeValue := function.GetParam(index)
			if loadValue == nil {
				fmt.Printf("variable '%s' does not exist:\n \tline\n", index)
				os.Exit(1)
			}
			return loadValue.Ir, typeValue
		}

	}
	return constant.NewInt(types.I64, int64(val)), constant.NewInt(types.I64, int64(lang.INT_T))
}

func AssignArrayElement(block *lang.Block, index parser.IArrayIndexContext, scope util.Stack, program *lang.Program, function *lang.Function) {

	indexValue, indexType := getIndex(block, index.ArrayIndexShift().GetText(), program, scope, function)
	args := GetValues(program, 1, block.Ir)
	setFun := program.Funcs[lang.SetT].IrFunc

	arr, arrT := GetVar(index.ArrayBase().GetText(), program, scope, block)
	if arr == nil {
		param, typ := function.GetParam(index.ArrayBase().GetText())
		if param == nil {
			fmt.Printf("variable '%s' does not exist:\n \tline\n", index.ArrayBase().GetText())
			os.Exit(1)
		}
		block.Ir.NewCall(setFun, indexValue, param.Ir, args[0], indexType, typ, args[1])
	} else {
		block.Ir.NewCall(setFun, indexValue, arr, args[0], indexType, arrT, args[1])

	}
}
