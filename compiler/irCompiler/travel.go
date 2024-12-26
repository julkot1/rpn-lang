package irCompiler

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"rpn/lang"
	"rpn/parser"
	"rpn/util"
)

type TreeWalk struct {
	*parser.BaseStcListener
	program       *lang.Program
	blockStack    util.Stack
	functionStack util.Stack
	scopeStack    util.Stack
	nonFreeBlock  bool
	ifBlock       bool
	repeatBlock   bool
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
	w.functionStack.Push(lang.NewFunction(ctx.ID().GetText(), w.program.Module))
}

func (w *TreeWalk) ExitFunctionDef(ctx *parser.FunctionDefContext) {
	pop, err := w.functionStack.Pop()

	if err != nil {
		os.Exit(-1)
	}
	fun := pop.(*lang.Function)
	w.program.Functions = append(w.program.Functions, fun)
}

func (w *TreeWalk) EnterVarAssing(ctx *parser.VarAssingContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}
	AssignVar(top.(*lang.Block), ctx.VarIdentifier().GetText(), w.scopeStack, w.program)
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
		topF, _ := w.functionStack.Top()
		ParseToken(ctx.GetText(), top.(*lang.Block), topF.(*lang.Function), w.program, w.scopeStack, ctx)
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

}
