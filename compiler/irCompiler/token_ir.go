package irCompiler

import (
	"fmt"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"os"
	"rpn/lang"
	"rpn/parser"
	"rpn/util"
)

func ParseToken(id *parser.IdentifierContext, block *lang.Block, fun *lang.Function, program *lang.Program, scope util.Stack, ctx *parser.IdentifierContext) {
	identifier := GetIdentifier(id)
	text := identifier.Base
	if fun.IsParameter(text) {
		if len(identifier.References) == 0 {
			param, paramType := fun.GetParam(id.ID(0).GetText())
			if param.Type != lang.ANY_T {
				block.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
					param.Ir,
					constant.NewInt(types.I64, int64(param.Type)))
			} else {
				block.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
					param.Ir,
					paramType)
			}

		} else {
			param, _ := fun.GetParam(id.ID(0).GetText())
			stcStruct := program.StcStruct[param.ComplexType]
			idx := GetStructFieldIndex(identifier, param.ComplexType, program)
			ptr := block.Ir.NewIntToPtr(param.Ir, types.NewPointer(stcStruct.IrType))
			gep := block.Ir.NewGetElementPtr(stcStruct.IrType, ptr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, int64(idx)))
			load2 := block.Ir.NewLoad(types.I64, gep)

			block.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
				load2,
				constant.NewInt(types.I64, int64(stcStruct.Args[idx].Type)))
		}
		return
	}

	tok, ok := program.GlobalTokenTable[text]
	if !ok {
		if TokenExists(text, scope) {
			PushIdentifier(text, scope, identifier, program, block)

			return
		}
		line := ctx.GetStart().GetLine()
		fmt.Printf("variable '%s' does not exist:\n \tline %v\n", text, line)
		os.Exit(1)
		return
	}
	switch tok {
	case lang.PStaticFunc:
		CallFunc(program.StaticFunctions[text].IrFunc, block.Ir, program, false)
		break
	case lang.PStcFunction:
		CallFunc(program.GetFunction(text, fun).Ir, block.Ir, program, false)
		break
	default:
		break
	}
}

func PushIdentifier(text string, scope util.Stack, identifier *Identifier, program *lang.Program, block *lang.Block) {
	token := GetToken(text, scope)
	if len(identifier.References) == 0 {
		PushToken(token, block, program)
	} else {
		stcStruct := program.StcStruct[token.ComplexType]
		idx := GetStructFieldIndex(identifier, token.ComplexType, program)
		load := block.Ir.NewLoad(types.I64, token.Ir)
		ptr := block.Ir.NewIntToPtr(load, types.NewPointer(stcStruct.IrType))
		gep := block.Ir.NewGetElementPtr(stcStruct.IrType, ptr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, int64(idx)))
		load2 := block.Ir.NewLoad(types.I64, gep)

		block.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
			load2,
			constant.NewInt(types.I64, int64(stcStruct.Args[idx].Type)))
	}
}
func GetStructFieldIndex(identifier *Identifier, complexType string, program *lang.Program) int {
	ref := identifier.References[0]
	stcStruct := program.StcStruct[complexType]
	for i, s := range stcStruct.Args {
		if s.Name == ref {
			return i
		}
	}
	return 0

}
