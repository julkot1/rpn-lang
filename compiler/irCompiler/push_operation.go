package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"math"
	"rpn/lang"
	"rpn/parser"
	"strconv"
)

func push(block *ir.Block, ctx *parser.PushContext, program *lang.Program) {
	var val value.Value
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
		strPtr := block.NewGetElementPtr(types.NewArray(uint64(len(str)+1), types.I8), g, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
		val = block.NewPtrToInt(strPtr, types.I64)
		block.NewCall(program.Funcs[lang.PushT].IrFunc,
			val,
			constant.NewInt(types.I64, int64(typ)))
		return
	}

	block.NewCall(program.Funcs[lang.PushT].IrFunc,
		val,
		constant.NewInt(types.I64, int64(typ)))
}

func trimString(text string) string {
	result, _ := strconv.Unquote(text)
	return result
}

func getChar(text string) int {
	if len(text) == 3 {
		b := text[1]
		i := int(b)
		return i
	} else {
		b := text[2]
		switch b {
		case 'n':
			return int('\n')
		case 't':
			return int('\t')
		case 'r':
			return int('\r')
		case 'f':
			return int('\f')
		case '0':
			return int('0')
		}
	}
	return 0
}
