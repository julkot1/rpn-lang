package irCompiler

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"os"
	"rpn/lang"
	"rpn/parser"
	"rpn/util"
)

type TreeWalk struct {
	*parser.BaseStcListener
	program        *lang.Program
	blockStack     util.Stack
	functionStack  util.Stack
	scopeStack     util.Stack
	nonFreeBlock   bool
	ifBlock        bool
	repeatBlock    bool
	notCreateScope bool
}

func getLineText(lineNumber int, tokenStream *antlr.CommonTokenStream) string {
	allTokens := tokenStream.GetAllTokens()
	lineText := ""

	for _, token := range allTokens {
		if token.GetLine() == lineNumber {
			lineText += token.GetText()
		}
	}
	return lineText
}

func NewTreeWalk() *TreeWalk {
	x := new(TreeWalk)
	x.blockStack = util.NewStack()
	x.functionStack = util.NewStack()
	x.nonFreeBlock = false
	x.ifBlock = false
	x.repeatBlock = false
	x.notCreateScope = false
	x.scopeStack = util.NewStack()

	return x
}

func (w *TreeWalk) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	err := w.program.AddGlobalToken(ctx.ID().GetText(), lang.PStcFunction)
	if err != nil {
		line := ctx.GetStart().GetLine()
		lineText := getLineText(line, w.program.Stream)
		fmt.Printf("Function %s is defined in other place\n \tline: %v: %s", ctx.ID().GetText(), line, lineText)
		os.Exit(-1)
	}

	params := make([]*lang.FuncParam, 0)
	if ctx.Arguments() != nil {
		args := ctx.Arguments().AllArgument()
		for _, arg := range args {
			param := &lang.FuncParam{}
			param.Name = arg.ID().GetText()
			if arg.VarType() == nil {
				param.Type = lang.ANY_T
			} else {
				param.Type = lang.StringToType(arg.VarType().GetText(), w.program)
			}
			param.Ir = ir.NewParam(param.Name, types.I64)
			if param.Type == lang.Struct_T {
				param.ComplexType = arg.VarType().Type_().GetText()
			}
			params = append(params, param)
		}

		ok, str := isUniqueParam(params)
		if !ok {
			fmt.Printf("parametr is not unique: %s in function %s\n\t line %v:", str, ctx.ID().GetText(), ctx.GetStart().GetLine())
			os.Exit(-1)
		}

		for _, arg := range args {
			param := &lang.FuncParam{}
			param.Name = arg.ID().GetText() + ".typ"
			if arg.VarType() == nil {
				param.Type = lang.ANY_T
			} else {
				param.Type = lang.StringToType(arg.VarType().GetText(), w.program)
			}
			param.Ir = ir.NewParam(param.Name, types.I64)
			params = append(params, param)
		}
	}
	w.functionStack.Push(lang.NewFunction(ctx.ID().GetText(), w.program.Module, params))

}
func isUnique(arr []string) (bool, string) {
	seen := make(map[string]bool)
	for _, v := range arr {
		if seen[v] {
			return false, v // Duplicate found
		}
		seen[v] = true
	}
	return true, "" // All elements are unique
}
func isUniqueParam(arr []*lang.FuncParam) (bool, string) {
	seen := make(map[string]bool)
	for _, v := range arr {
		if seen[v.Name] {
			return false, v.Name // Duplicate found
		}
		seen[v.Name] = true
	}
	return true, "" // All elements are unique
}
func (w *TreeWalk) ExitFunctionDef(ctx *parser.FunctionDefContext) {
	pop, err := w.functionStack.Pop()

	if err != nil {
		os.Exit(-1)
	}
	fun := pop.(*lang.Function)
	w.program.Functions = append(w.program.Functions, fun)
}

func (w *TreeWalk) EnterVarAssign(ctx *parser.VarAssignContext) {
	top, err := w.blockStack.Top()
	topF, e := w.functionStack.Top()
	if err != nil || e != nil {
		os.Exit(-1)
	}
	if ctx.ArrayIndex() != nil {
		AssignArrayElement(top.(*lang.Block), ctx.ArrayIndex(), w.scopeStack, w.program, topF.(*lang.Function))
	} else {
		AssignVar(top.(*lang.Block), ctx, w.scopeStack, w.program)
	}
}

func (w *TreeWalk) EnterVarReference(ctx *parser.VarReferenceContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	PushReference(top.(*lang.Block), ctx.VarIdentifier().GetText(), w, ctx)
}

func (w *TreeWalk) EnterSubBlock(ctx *parser.SubBlockContext) {
	top, err := w.functionStack.Top()
	if err != nil {
		os.Exit(-1)
	}

	w.blockStack.Push(lang.NewBlockIr(w.program.NewBlockIndex(), !w.nonFreeBlock, top.(*lang.Function), false, false))
	w.nonFreeBlock = false
}

func (w *TreeWalk) ExitSubBlock(ctx *parser.SubBlockContext) {

	pop, err := w.blockStack.Pop()
	top, _ := w.functionStack.Top()
	pop.(*lang.Block).CreateIf = w.ifBlock
	pop.(*lang.Block).CreateLoop = w.repeatBlock

	w.ifBlock = false
	w.repeatBlock = false

	if err != nil {
		os.Exit(-1)
	}
	addBlock := true
	for _, b := range top.(*lang.Function).Blocks {
		if b.Name == pop.(*lang.Block).Name {
			addBlock = false
			break
		}
	}
	if addBlock {
		top.(*lang.Function).Blocks = append(top.(*lang.Function).Blocks, pop.(*lang.Block))
	}

}

func (w *TreeWalk) EnterBlock(ctx *parser.BlockContext) {
	if !w.notCreateScope {

		w.nonFreeBlock = false
		w.notCreateScope = !w.notCreateScope
	}
	scope := NewScope()
	w.scopeStack.Push(scope)
}

func (w *TreeWalk) ExitBlock(ctx *parser.BlockContext) {
	w.scopeStack.Pop()
}

func (w *TreeWalk) EnterPush(ctx *parser.PushContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	push(top.(*lang.Block).Ir, ctx, w.program)
}
func (w *TreeWalk) EnterOperation(ctx *parser.OperationContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	typ := lang.StrToTokenType(ctx.Operaor().GetText())

	prevent := ctx.STACK_PREVENTION() != nil
	switch typ {
	case lang.AssignT:
		AssignVarRef(w, top.(*lang.Block).Ir, prevent)
		break
	default:
		CallFunc(w.program.Funcs[typ].IrFunc, top.(*lang.Block).Ir, w.program, prevent)
		break

	}

}

func (w *TreeWalk) EnterStackOperation(ctx *parser.StackOperationContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	typ := lang.StrToTokenType(ctx.GetText())

	CallFunc(w.program.Funcs[typ].IrFunc, top.(*lang.Block).Ir, w.program, false)
	if typ == lang.PopT {
		CallFunc(w.program.Funcs[lang.PopTypeT].IrFunc, top.(*lang.Block).Ir, w.program, false)

	}
}
func (w *TreeWalk) EnterIdentifier(ctx *parser.IdentifierContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}

	if ctx.ID(0).GetText() == lang.PrintToken {

		CallFunc(w.program.Funcs[lang.PrintT].IrFunc, top.(*lang.Block).Ir, w.program, ctx.STACK_PREVENTION() != nil)
	} else if ctx.ID(0).GetText() == lang.InputToken {
		CallFunc(w.program.Funcs[lang.InputT].IrFunc, top.(*lang.Block).Ir, w.program, false)
	} else {
		_, ok := ctx.GetParent().(*parser.VarAssignContext)
		if ok {
			return
		}
		topF, _ := w.functionStack.Top()
		ParseToken(ctx, top.(*lang.Block), topF.(*lang.Function), w.program, w.scopeStack, ctx)
	}

}

func (w *TreeWalk) EnterIfBlock(ctx *parser.IfBlockContext) {
	top, err := w.functionStack.Top()
	if err != nil {
		os.Exit(-1)
	}

	pop, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}

	top.(*lang.Function).Blocks = append(top.(*lang.Function).Blocks, pop.(*lang.Block))
	w.nonFreeBlock = true
}
func (w *TreeWalk) ExitIfBlock(ctx *parser.IfBlockContext) {
	w.ifBlock = true
}

func (w *TreeWalk) EnterRepeatBlock(ctx *parser.RepeatBlockContext) {
	top, err := w.functionStack.Top()
	if err != nil {
		os.Exit(-1)
	}

	pop, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	top.(*lang.Function).Blocks = append(top.(*lang.Function).Blocks, pop.(*lang.Block))

	w.nonFreeBlock = true

	if ctx.Arguments() == nil {
		return
	}
	if len(ctx.Arguments().AllArgument()) != 1 {
		fmt.Printf("Too many arguments repeat statment \n\tline %v:", ctx.Arguments().GetStart().GetLine())
		os.Exit(-1)
	}
	arg := ctx.Arguments().AllArgument()[0].GetText()
	scope := NewScope()
	w.scopeStack.Push(scope)
	createVar(arg, lang.INT_T, "", pop.(*lang.Block), scope, w.program)

	block := pop.(*lang.Block)

	block.Ir.NewStore(constant.NewInt(types.I64, -1), scope.tokens[arg].Ir)
	scope.tokens[arg].Type = lang.INT_T
	block.Vars["loop_index"] = scope.tokens[arg]
	w.notCreateScope = true

}
func (w *TreeWalk) ExitRepeatBlock(ctx *parser.RepeatBlockContext) {
	w.repeatBlock = true
	top, err := w.functionStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	b := lang.NewBlockIr(w.program.NewBlockIndex(), true, top.(*lang.Function), false, false)
	b.LoopCondition = true
	top.(*lang.Function).Blocks = append(top.(*lang.Function).Blocks, b)

	if ctx.Arguments() != nil {
		w.scopeStack.Pop()
	}

}

func (w *TreeWalk) EnterArray(ctx *parser.ArrayContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	parent := ctx.GetParent()
	_, ok := parent.(*parser.SubBlockContext)
	if ok {
		CreateArray(ctx, top.(*lang.Block), w.program)
	}
}
func (w *TreeWalk) EnterArrayIndex(ctx *parser.ArrayIndexContext) {
	parent := ctx.GetParent()
	_, ok := parent.(*parser.VarAssignContext)
	if ok {
		return
	}
	base := ctx.ArrayBase().GetText()
	index := ctx.ArrayIndexShift().GetText()

	topBlock, err := w.blockStack.Top()
	topF, err := w.functionStack.Top()
	if err != nil {
		os.Exit(-1)
	}

	getElementAtIndex(topBlock.(*lang.Block), w.scopeStack, w.program, base, index, topF.(*lang.Function))
}

func (w *TreeWalk) EnterStruct(ctx *parser.StructContext) {
	name := ctx.ID().GetText()
	err := w.program.AddGlobalToken(name, lang.PStruct)
	if err != nil {
		os.Exit(-1)
	}

	args := getArgs(ctx.StructBody().AllStructElement(), name, w.program)
	createStructDefinition(name, args, w.program)

}

func (w *TreeWalk) EnterNewOperator(ctx *parser.NewOperatorContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	name := ctx.ID().GetText()
	createStruct(name, top.(*lang.Block), w.program)

}

func (w *TreeWalk) EnterNewStructAssign(ctx *parser.NewStructAssignContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	name := ctx.ID(0).GetText()
	varName := ctx.ID(1).GetText()

	createStructAndAssign(name, varName, top.(*lang.Block), w.program, w.scopeStack)

}

func createStructAndAssign(structName string, varName string, block *lang.Block, program *lang.Program, stack util.Stack) {

	tok := GetToken(varName, stack)
	if tok == nil {
		top, _ := stack.Top()
		createVar(varName, lang.Struct_T, "<"+structName+">", block, top.(*Scope), program)
		tok = top.(*Scope).tokens[varName]
	}

	x, ok := program.GlobalTokenTable[structName]
	if !ok {
		fmt.Println("No such struct: " + structName)
		os.Exit(1)
	}
	if x != lang.PStruct {
		fmt.Println("Token " + structName + " is not a struct")
	}
	stcStruct := program.StcStruct[structName]
	alloc := block.Ir.NewAlloca(stcStruct.IrType)

	intPtr := block.Ir.NewPtrToInt(alloc, types.I64)
	block.Ir.NewStore(intPtr, tok.Ir)
}
