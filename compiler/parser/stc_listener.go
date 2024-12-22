// Code generated from Stc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Stc
import "github.com/antlr4-go/antlr/v4"

// StcListener is a complete listener for a parse tree produced by StcParser.
type StcListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterFunctionDef is called when entering the functionDef production.
	EnterFunctionDef(c *FunctionDefContext)

	// EnterSubBlock is called when entering the subBlock production.
	EnterSubBlock(c *SubBlockContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterRepeatBlock is called when entering the repeatBlock production.
	EnterRepeatBlock(c *RepeatBlockContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterOperaor is called when entering the operaor production.
	EnterOperaor(c *OperaorContext)

	// EnterStackOperation is called when entering the stackOperation production.
	EnterStackOperation(c *StackOperationContext)

	// EnterPush is called when entering the push production.
	EnterPush(c *PushContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitFunctionDef is called when exiting the functionDef production.
	ExitFunctionDef(c *FunctionDefContext)

	// ExitSubBlock is called when exiting the subBlock production.
	ExitSubBlock(c *SubBlockContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitRepeatBlock is called when exiting the repeatBlock production.
	ExitRepeatBlock(c *RepeatBlockContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitOperaor is called when exiting the operaor production.
	ExitOperaor(c *OperaorContext)

	// ExitStackOperation is called when exiting the stackOperation production.
	ExitStackOperation(c *StackOperationContext)

	// ExitPush is called when exiting the push production.
	ExitPush(c *PushContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
