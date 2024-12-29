package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"math"
	"rpn/lang"
	"rpn/parser"
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
