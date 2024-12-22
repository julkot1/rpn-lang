package irCompiler

import (
	"fmt"
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
	nonFreeBlock  bool
	ifBlock       bool
}

func NewTreeWalk() *TreeWalk {
	x := new(TreeWalk)
	x.blockStack = util.NewStack()
	x.functionStack = util.NewStack()
	x.nonFreeBlock = false
	x.ifBlock = false
	return x
}

func (w *TreeWalk) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	err := w.program.AddGlobalToken(ctx.ID().GetText(), lang.PStcFunction)
	if err != nil {
		fmt.Println(err)
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

func (w *TreeWalk) EnterSubBlock(ctx *parser.SubBlockContext) {
	top, err := w.functionStack.Top()
	if err != nil {
		os.Exit(-1)
	}

	w.blockStack.Push(lang.NewBlockIr(w.program.NewBlockIndex(), !w.nonFreeBlock, top.(*lang.Function), false))
	w.nonFreeBlock = false
}

func (w *TreeWalk) ExitSubBlock(ctx *parser.SubBlockContext) {

	pop, err := w.blockStack.Pop()
	top, _ := w.functionStack.Top()
	pop.(*lang.Block).CreateIf = w.ifBlock
	w.ifBlock = false

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
	typ := lang.StrToTokenType(ctx.GetText())
	prevent := ctx.STACK_PREVENTION() != nil

	CallFunc(w.program.Funcs[typ].IrFunc, top.(*lang.Block).Ir, w.program, prevent)
}

func (w *TreeWalk) EnterIdentifier(ctx *parser.IdentifierContext) {
	top, err := w.blockStack.Top()
	if err != nil {
		os.Exit(-1)
	}

	if ctx.GetText() == lang.PrintToken {
		CallFunc(w.program.Funcs[lang.PrintT].IrFunc, top.(*lang.Block).Ir, w.program, false)
	}
	if ctx.GetText() == lang.InputToken {
		CallFunc(w.program.Funcs[lang.InputT].IrFunc, top.(*lang.Block).Ir, w.program, false)
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
