package irCompiler

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"os"
	"rpn/lang"
	"rpn/parser"
	"rpn/util"
)

type Scope struct {
	tokens map[string]*lang.Var
}

func NewScope() *Scope {
	tokens := make(map[string]*lang.Var)
	return &Scope{tokens}
}

func (s *Scope) Push(token string, scopeEl *lang.Var) {
	s.tokens[token] = scopeEl
}

func (s *Scope) IsInScope(token string) bool {
	_, ok := s.tokens[token]
	return ok
}
func TokenExists(token string, stack util.Stack) bool {
	for _, sc := range stack.Items() {
		if sc.(*Scope).IsInScope(token) {
			return true
		}
	}
	return false
}

func GetVar(token string, program *lang.Program, scope util.Stack, block *lang.Block) (*ir.InstLoad, *constant.Int) {
	_, ok := program.GlobalTokenTable[token]
	if !ok {
		if TokenExists(token, scope) {
			variable := GetToken(token, scope)

			return block.Ir.NewLoad(types.I64, variable.Ir), constant.NewInt(types.I64, int64(variable.Type))
		}
		return nil, nil
	}
	return nil, nil
}

func GetToken(token string, stack util.Stack) *lang.Var {
	for _, sc := range stack.Items() {
		if sc.(*Scope).IsInScope(token) {
			return sc.(*Scope).tokens[token]
		}
	}
	return nil
}

func PushToken(variable *lang.Var, block *lang.Block, program *lang.Program) {

	load := block.Ir.NewLoad(types.I64, variable.Ir)

	block.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
		load,
		constant.NewInt(types.I64, int64(variable.Type)))
}

func createVar(text string, typ lang.Type, typS string, block *lang.Block, scope *Scope) {
	variableIr := block.Ir.NewAlloca(types.I64)
	variable := &lang.Var{Name: text, Type: typ, Ir: variableIr, ComplexType: typS[1 : len(typS)-1]}
	scope.Push(text, variable)

}

func AssignVar(block *lang.Block, ctx *parser.VarAssignContext, stack util.Stack, program *lang.Program) {

	if ctx.VarAssignIdentifier() != nil {
		name := ctx.VarAssignIdentifier().VarIdentifier().GetText()
		typ := ctx.VarAssignIdentifier().VarType().GetText()
		tok := GetToken(name, stack)
		if tok == nil {
			top, _ := stack.Top()
			createVar(name, lang.StringToType(typ, program), typ, block, top.(*Scope))
			tok = top.(*Scope).tokens[name]
		}
		args := GetValues(program, 1, block.Ir)

		block.Ir.NewStore(args[0], tok.Ir)
		return
	}
	if ctx.Identifier() != nil {
		idCtx := ctx.Identifier()
		id := GetIdentifier(idCtx.(*parser.IdentifierContext))
		token := GetToken(id.Base, stack)
		if token == nil {
			fmt.Printf("Var %s not declarated", id.Base)
			os.Exit(1)
		}
		stcStruct := program.StcStruct[token.ComplexType]
		idx := GetStructFieldIndex(id, token, program)
		load := block.Ir.NewLoad(types.I64, token.Ir)
		ptr := block.Ir.NewIntToPtr(load, types.NewPointer(stcStruct.IrType))
		gep := block.Ir.NewGetElementPtr(stcStruct.IrType, ptr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, int64(idx)))
		args := GetValues(program, 1, block.Ir)

		block.Ir.NewStore(args[0], gep)
	}

}
func PushReference(block *lang.Block, text string, w *TreeWalk, ctx *parser.VarReferenceContext) {
	if TokenExists(text, w.scopeStack) {
		token := GetToken(text, w.scopeStack)
		ptrI64 := block.Ir.NewPtrToInt(token.Ir, types.I64)
		typ := constant.NewInt(types.I64, int64(lang.REF_T))

		block.Ir.NewCall(w.program.Funcs[lang.PushT].IrFunc,
			ptrI64,
			typ)
		return
	}
	line := ctx.GetStart().GetLine()
	fmt.Printf("variable '%s' does not exist:\n \tline %v\n", text, line)
	os.Exit(1)
}

func AssignVarRef(w *TreeWalk, block *ir.Block, prevent bool) {
	args := GetValues(w.program, 2, block)
	varTyp := w.program.Structs["variable"]
	ptr := block.NewIntToPtr(args[0], types.NewPointer(varTyp))

	loadPtr := block.NewGetElementPtr(varTyp, ptr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	block.NewStore(args[1], loadPtr)

	loadTypePtr := block.NewGetElementPtr(varTyp, ptr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	block.NewStore(args[3], loadTypePtr)

	return

}
