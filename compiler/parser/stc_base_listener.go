// Code generated from Stc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Stc

import "github.com/antlr4-go/antlr/v4"

// BaseStcListener is a complete listener for a parse tree produced by StcParser.
type BaseStcListener struct{}

var _ StcListener = &BaseStcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseStcListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseStcListener) ExitProg(ctx *ProgContext) {}

// EnterFunctionDef is called when production functionDef is entered.
func (s *BaseStcListener) EnterFunctionDef(ctx *FunctionDefContext) {}

// ExitFunctionDef is called when production functionDef is exited.
func (s *BaseStcListener) ExitFunctionDef(ctx *FunctionDefContext) {}

// EnterType is called when production type is entered.
func (s *BaseStcListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseStcListener) ExitType(ctx *TypeContext) {}

// EnterSubBlock is called when production subBlock is entered.
func (s *BaseStcListener) EnterSubBlock(ctx *SubBlockContext) {}

// ExitSubBlock is called when production subBlock is exited.
func (s *BaseStcListener) ExitSubBlock(ctx *SubBlockContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseStcListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseStcListener) ExitBlock(ctx *BlockContext) {}

// EnterNewStructAssign is called when production newStructAssign is entered.
func (s *BaseStcListener) EnterNewStructAssign(ctx *NewStructAssignContext) {}

// ExitNewStructAssign is called when production newStructAssign is exited.
func (s *BaseStcListener) ExitNewStructAssign(ctx *NewStructAssignContext) {}

// EnterNewOperator is called when production newOperator is entered.
func (s *BaseStcListener) EnterNewOperator(ctx *NewOperatorContext) {}

// ExitNewOperator is called when production newOperator is exited.
func (s *BaseStcListener) ExitNewOperator(ctx *NewOperatorContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseStcListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseStcListener) ExitStruct(ctx *StructContext) {}

// EnterStructElement is called when production structElement is entered.
func (s *BaseStcListener) EnterStructElement(ctx *StructElementContext) {}

// ExitStructElement is called when production structElement is exited.
func (s *BaseStcListener) ExitStructElement(ctx *StructElementContext) {}

// EnterStructBody is called when production structBody is entered.
func (s *BaseStcListener) EnterStructBody(ctx *StructBodyContext) {}

// ExitStructBody is called when production structBody is exited.
func (s *BaseStcListener) ExitStructBody(ctx *StructBodyContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseStcListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseStcListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseStcListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseStcListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterRepeatBlock is called when production repeatBlock is entered.
func (s *BaseStcListener) EnterRepeatBlock(ctx *RepeatBlockContext) {}

// ExitRepeatBlock is called when production repeatBlock is exited.
func (s *BaseStcListener) ExitRepeatBlock(ctx *RepeatBlockContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseStcListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseStcListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseStcListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseStcListener) ExitOperation(ctx *OperationContext) {}

// EnterOperaor is called when production operaor is entered.
func (s *BaseStcListener) EnterOperaor(ctx *OperaorContext) {}

// ExitOperaor is called when production operaor is exited.
func (s *BaseStcListener) ExitOperaor(ctx *OperaorContext) {}

// EnterStackOperation is called when production stackOperation is entered.
func (s *BaseStcListener) EnterStackOperation(ctx *StackOperationContext) {}

// ExitStackOperation is called when production stackOperation is exited.
func (s *BaseStcListener) ExitStackOperation(ctx *StackOperationContext) {}

// EnterPush is called when production push is entered.
func (s *BaseStcListener) EnterPush(ctx *PushContext) {}

// ExitPush is called when production push is exited.
func (s *BaseStcListener) ExitPush(ctx *PushContext) {}

// EnterArrayElement is called when production arrayElement is entered.
func (s *BaseStcListener) EnterArrayElement(ctx *ArrayElementContext) {}

// ExitArrayElement is called when production arrayElement is exited.
func (s *BaseStcListener) ExitArrayElement(ctx *ArrayElementContext) {}

// EnterArrayBase is called when production arrayBase is entered.
func (s *BaseStcListener) EnterArrayBase(ctx *ArrayBaseContext) {}

// ExitArrayBase is called when production arrayBase is exited.
func (s *BaseStcListener) ExitArrayBase(ctx *ArrayBaseContext) {}

// EnterArrayIndexShift is called when production arrayIndexShift is entered.
func (s *BaseStcListener) EnterArrayIndexShift(ctx *ArrayIndexShiftContext) {}

// ExitArrayIndexShift is called when production arrayIndexShift is exited.
func (s *BaseStcListener) ExitArrayIndexShift(ctx *ArrayIndexShiftContext) {}

// EnterArrayIndex is called when production arrayIndex is entered.
func (s *BaseStcListener) EnterArrayIndex(ctx *ArrayIndexContext) {}

// ExitArrayIndex is called when production arrayIndex is exited.
func (s *BaseStcListener) ExitArrayIndex(ctx *ArrayIndexContext) {}

// EnterCapacity is called when production capacity is entered.
func (s *BaseStcListener) EnterCapacity(ctx *CapacityContext) {}

// ExitCapacity is called when production capacity is exited.
func (s *BaseStcListener) ExitCapacity(ctx *CapacityContext) {}

// EnterArray is called when production array is entered.
func (s *BaseStcListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseStcListener) ExitArray(ctx *ArrayContext) {}

// EnterArrayDescriber is called when production arrayDescriber is entered.
func (s *BaseStcListener) EnterArrayDescriber(ctx *ArrayDescriberContext) {}

// ExitArrayDescriber is called when production arrayDescriber is exited.
func (s *BaseStcListener) ExitArrayDescriber(ctx *ArrayDescriberContext) {}

// EnterArrayNew is called when production arrayNew is entered.
func (s *BaseStcListener) EnterArrayNew(ctx *ArrayNewContext) {}

// ExitArrayNew is called when production arrayNew is exited.
func (s *BaseStcListener) ExitArrayNew(ctx *ArrayNewContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseStcListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseStcListener) ExitArgument(ctx *ArgumentContext) {}

// EnterVarAssign is called when production varAssign is entered.
func (s *BaseStcListener) EnterVarAssign(ctx *VarAssignContext) {}

// ExitVarAssign is called when production varAssign is exited.
func (s *BaseStcListener) ExitVarAssign(ctx *VarAssignContext) {}

// EnterVarAssignIdentifier is called when production varAssignIdentifier is entered.
func (s *BaseStcListener) EnterVarAssignIdentifier(ctx *VarAssignIdentifierContext) {}

// ExitVarAssignIdentifier is called when production varAssignIdentifier is exited.
func (s *BaseStcListener) ExitVarAssignIdentifier(ctx *VarAssignIdentifierContext) {}

// EnterVarType is called when production varType is entered.
func (s *BaseStcListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BaseStcListener) ExitVarType(ctx *VarTypeContext) {}

// EnterVarReference is called when production varReference is entered.
func (s *BaseStcListener) EnterVarReference(ctx *VarReferenceContext) {}

// ExitVarReference is called when production varReference is exited.
func (s *BaseStcListener) ExitVarReference(ctx *VarReferenceContext) {}

// EnterVarIdentifier is called when production varIdentifier is entered.
func (s *BaseStcListener) EnterVarIdentifier(ctx *VarIdentifierContext) {}

// ExitVarIdentifier is called when production varIdentifier is exited.
func (s *BaseStcListener) ExitVarIdentifier(ctx *VarIdentifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseStcListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseStcListener) ExitIdentifier(ctx *IdentifierContext) {}
