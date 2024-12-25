package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
	"rpn/util"
)

type Scope struct {
	tokens map[string]*ScopeElement
}

type ScopeElement struct {
	element *ir.InstAlloca
	typ     *ir.InstAlloca
}

func NewScope() *Scope {
	tokens := make(map[string]*ScopeElement)
	return &Scope{tokens}
}

func (s *Scope) Push(token string, scopeEl *ScopeElement) {
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
func GetToken(token string, stack util.Stack) *ScopeElement {
	for _, sc := range stack.Items() {
		if sc.(*Scope).IsInScope(token) {
			return sc.(*Scope).tokens[token]
		}
	}
	return nil
}

func ParseToken(text string, block *lang.Block, fun *lang.Function, program *lang.Program, scope util.Stack) {
	tok, ok := program.GlobalTokenTable[text]
	if !ok {
		if TokenExists(text, scope) {
			token := GetToken(text, scope)
			PushToken(token, block, program)
			return
		}
		top, _ := scope.Top()
		createVar(text, block, top.(*Scope))
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

func PushToken(token *ScopeElement, block *lang.Block, program *lang.Program) {
	load := block.Ir.NewLoad(types.I64, token.element)
	loadType := block.Ir.NewLoad(types.I64, token.typ)
	block.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
		load,
		loadType)
}

func createVar(text string, block *lang.Block, scope *Scope) {
	variable := block.Ir.NewAlloca(types.I64)
	variableType := block.Ir.NewAlloca(types.I64)
	scope.Push(text, &ScopeElement{variable, variableType})

}

func AssignVar(block *lang.Block, text string, stack util.Stack, program *lang.Program) {
	tok := GetToken(text, stack)
	if tok == nil {
		top, _ := stack.Top()
		createVar(text, block, top.(*Scope))
		tok = top.(*Scope).tokens[text]
	}
	args := GetValues(program, 1, block.Ir)
	block.Ir.NewStore(args[0], tok.element)
	block.Ir.NewStore(args[1], tok.typ)

}
