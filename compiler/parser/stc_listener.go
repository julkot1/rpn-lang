// Code generated from Stc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Stc

import "github.com/antlr4-go/antlr/v4"

// StcListener is a complete listener for a parse tree produced by StcParser.
type StcListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterFunctionDef is called when entering the functionDef production.
	EnterFunctionDef(c *FunctionDefContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterSubBlock is called when entering the subBlock production.
	EnterSubBlock(c *SubBlockContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterNewStructAssign is called when entering the newStructAssign production.
	EnterNewStructAssign(c *NewStructAssignContext)

	// EnterNewOperator is called when entering the newOperator production.
	EnterNewOperator(c *NewOperatorContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterStructElement is called when entering the structElement production.
	EnterStructElement(c *StructElementContext)

	// EnterStructBody is called when entering the structBody production.
	EnterStructBody(c *StructBodyContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterRepeatBlock is called when entering the repeatBlock production.
	EnterRepeatBlock(c *RepeatBlockContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterOperaor is called when entering the operaor production.
	EnterOperaor(c *OperaorContext)

	// EnterStackOperation is called when entering the stackOperation production.
	EnterStackOperation(c *StackOperationContext)

	// EnterPush is called when entering the push production.
	EnterPush(c *PushContext)

	// EnterArrayElement is called when entering the arrayElement production.
	EnterArrayElement(c *ArrayElementContext)

	// EnterArrayBase is called when entering the arrayBase production.
	EnterArrayBase(c *ArrayBaseContext)

	// EnterArrayIndexShift is called when entering the arrayIndexShift production.
	EnterArrayIndexShift(c *ArrayIndexShiftContext)

	// EnterArrayIndex is called when entering the arrayIndex production.
	EnterArrayIndex(c *ArrayIndexContext)

	// EnterCapacity is called when entering the capacity production.
	EnterCapacity(c *CapacityContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterArrayDescriber is called when entering the arrayDescriber production.
	EnterArrayDescriber(c *ArrayDescriberContext)

	// EnterArrayNew is called when entering the arrayNew production.
	EnterArrayNew(c *ArrayNewContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterVarAssign is called when entering the varAssign production.
	EnterVarAssign(c *VarAssignContext)

	// EnterVarAssignIdentifier is called when entering the varAssignIdentifier production.
	EnterVarAssignIdentifier(c *VarAssignIdentifierContext)

	// EnterVarType is called when entering the varType production.
	EnterVarType(c *VarTypeContext)

	// EnterVarReference is called when entering the varReference production.
	EnterVarReference(c *VarReferenceContext)

	// EnterVarIdentifier is called when entering the varIdentifier production.
	EnterVarIdentifier(c *VarIdentifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitFunctionDef is called when exiting the functionDef production.
	ExitFunctionDef(c *FunctionDefContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitSubBlock is called when exiting the subBlock production.
	ExitSubBlock(c *SubBlockContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitNewStructAssign is called when exiting the newStructAssign production.
	ExitNewStructAssign(c *NewStructAssignContext)

	// ExitNewOperator is called when exiting the newOperator production.
	ExitNewOperator(c *NewOperatorContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitStructElement is called when exiting the structElement production.
	ExitStructElement(c *StructElementContext)

	// ExitStructBody is called when exiting the structBody production.
	ExitStructBody(c *StructBodyContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitRepeatBlock is called when exiting the repeatBlock production.
	ExitRepeatBlock(c *RepeatBlockContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitOperaor is called when exiting the operaor production.
	ExitOperaor(c *OperaorContext)

	// ExitStackOperation is called when exiting the stackOperation production.
	ExitStackOperation(c *StackOperationContext)

	// ExitPush is called when exiting the push production.
	ExitPush(c *PushContext)

	// ExitArrayElement is called when exiting the arrayElement production.
	ExitArrayElement(c *ArrayElementContext)

	// ExitArrayBase is called when exiting the arrayBase production.
	ExitArrayBase(c *ArrayBaseContext)

	// ExitArrayIndexShift is called when exiting the arrayIndexShift production.
	ExitArrayIndexShift(c *ArrayIndexShiftContext)

	// ExitArrayIndex is called when exiting the arrayIndex production.
	ExitArrayIndex(c *ArrayIndexContext)

	// ExitCapacity is called when exiting the capacity production.
	ExitCapacity(c *CapacityContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitArrayDescriber is called when exiting the arrayDescriber production.
	ExitArrayDescriber(c *ArrayDescriberContext)

	// ExitArrayNew is called when exiting the arrayNew production.
	ExitArrayNew(c *ArrayNewContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitVarAssign is called when exiting the varAssign production.
	ExitVarAssign(c *VarAssignContext)

	// ExitVarAssignIdentifier is called when exiting the varAssignIdentifier production.
	ExitVarAssignIdentifier(c *VarAssignIdentifierContext)

	// ExitVarType is called when exiting the varType production.
	ExitVarType(c *VarTypeContext)

	// ExitVarReference is called when exiting the varReference production.
	ExitVarReference(c *VarReferenceContext)

	// ExitVarIdentifier is called when exiting the varIdentifier production.
	ExitVarIdentifier(c *VarIdentifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
