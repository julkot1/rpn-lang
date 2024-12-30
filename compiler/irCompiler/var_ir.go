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
	tokens map[string]*ir.InstAlloca
}

func NewScope() *Scope {
	tokens := make(map[string]*ir.InstAlloca)
	return &Scope{tokens}
}

func (s *Scope) Push(token string, scopeEl *ir.InstAlloca) {
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

func GetVar(token string, program *lang.Program, scope util.Stack) *ir.InstAlloca {
	_, ok := program.GlobalTokenTable[token]
	if !ok {
		if TokenExists(token, scope) {
			return GetToken(token, scope)
		}
		fmt.Printf("variable '%s' does not exist:\n \tline\n", token)
		os.Exit(1)
		return nil
	}
	return nil
}

func GetToken(token string, stack util.Stack) *ir.InstAlloca {
	for _, sc := range stack.Items() {
		if sc.(*Scope).IsInScope(token) {
			return sc.(*Scope).tokens[token]
		}
	}
	return nil
}

func PushToken(token *ir.InstAlloca, block *lang.Block, program *lang.Program) {
	typ := program.Structs["variable"]

	loadPtr := block.Ir.NewGetElementPtr(typ, token, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	load := block.Ir.NewLoad(types.I64, loadPtr)

	loadTypePtr := block.Ir.NewGetElementPtr(typ, token, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	loadType := block.Ir.NewLoad(types.I64, loadTypePtr)

	block.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
		load,
		loadType)
}

func createVar(text string, block *lang.Block, scope *Scope, program *lang.Program) {
	typ := program.Structs["variable"]
	variable := block.Ir.NewAlloca(typ)
	scope.Push(text, variable)

}

func AssignVar(block *lang.Block, text string, stack util.Stack, program *lang.Program) {
	tok := GetToken(text, stack)
	if tok == nil {
		top, _ := stack.Top()
		createVar(text, block, top.(*Scope), program)
		tok = top.(*Scope).tokens[text]
	}
	args := GetValues(program, 1, block.Ir)
	typ := program.Structs["variable"]

	loadPtr := block.Ir.NewGetElementPtr(typ, tok, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	block.Ir.NewStore(args[0], loadPtr)

	loadTypePtr := block.Ir.NewGetElementPtr(typ, tok, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	block.Ir.NewStore(args[1], loadTypePtr)

}
func PushReference(block *lang.Block, text string, w *TreeWalk, ctx *parser.VarReferenceContext) {
	if TokenExists(text, w.scopeStack) {
		token := GetToken(text, w.scopeStack)
		ptrI64 := block.Ir.NewPtrToInt(token, types.I64)
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
