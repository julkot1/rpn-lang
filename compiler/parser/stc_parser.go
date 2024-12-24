// Code generated from Stc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Stc

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type StcParser struct {
	*antlr.BaseParser
}

var StcParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stcParserInit() {
	staticData := &StcParserStaticData
	staticData.LiteralNames = []string{
		"", "'fun'", "'{'", "'}'", "'if'", "'else'", "'repeat'", "'dup'", "'rot'",
		"'swap'", "'pop'", "'over'", "':'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "STACK_PREVENTION",
		"NUMBER", "SIGNED_NUMBER", "FLOAT", "SIGNED_FLOAT", "FLOAT_E", "SIGNED_FLOAT_E",
		"BIN_NUMBER", "HEX_NUMBER", "STRING", "CHAR", "BOOL", "SIMPLE_TYPE",
		"LOGIC_OPERATOR", "ARITHMETIC_OPERATOR", "BUILD_IN_OPERATOR", "ID",
		"WS",
	}
	staticData.RuleNames = []string{
		"prog", "functionDef", "subBlock", "block", "ifBlock", "elseBlock",
		"repeatBlock", "operation", "operaor", "stackOperation", "push", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 4, 0, 26, 8, 0, 11, 0, 12, 0, 27, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 40, 8, 2, 11, 2, 12,
		2, 41, 1, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 3, 1, 3, 4, 3, 51, 8, 3, 11,
		3, 12, 3, 52, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 60, 8, 4, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 3, 7, 69, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 4, 11, 82, 8, 11, 11, 11, 12,
		11, 83, 3, 11, 86, 8, 11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 0, 3, 1, 0, 26, 28, 1, 0, 7, 11, 1, 0, 14, 25, 88, 0, 25, 1,
		0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 48, 1, 0, 0, 0, 8, 56,
		1, 0, 0, 0, 10, 61, 1, 0, 0, 0, 12, 64, 1, 0, 0, 0, 14, 68, 1, 0, 0, 0,
		16, 72, 1, 0, 0, 0, 18, 74, 1, 0, 0, 0, 20, 76, 1, 0, 0, 0, 22, 78, 1,
		0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27,
		25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30, 5, 0, 0,
		1, 30, 1, 1, 0, 0, 0, 31, 32, 5, 1, 0, 0, 32, 33, 5, 29, 0, 0, 33, 34,
		3, 6, 3, 0, 34, 3, 1, 0, 0, 0, 35, 40, 3, 14, 7, 0, 36, 40, 3, 18, 9, 0,
		37, 40, 3, 20, 10, 0, 38, 40, 3, 22, 11, 0, 39, 35, 1, 0, 0, 0, 39, 36,
		1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0,
		41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 46, 1, 0, 0, 0, 43, 47, 3,
		2, 1, 0, 44, 47, 3, 8, 4, 0, 45, 47, 3, 12, 6, 0, 46, 43, 1, 0, 0, 0, 46,
		44, 1, 0, 0, 0, 46, 45, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 5, 1, 0, 0,
		0, 48, 50, 5, 2, 0, 0, 49, 51, 3, 4, 2, 0, 50, 49, 1, 0, 0, 0, 51, 52,
		1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 55, 5, 3, 0, 0, 55, 7, 1, 0, 0, 0, 56, 57, 5, 4, 0, 0, 57, 59, 3, 6,
		3, 0, 58, 60, 3, 10, 5, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60,
		9, 1, 0, 0, 0, 61, 62, 5, 5, 0, 0, 62, 63, 3, 6, 3, 0, 63, 11, 1, 0, 0,
		0, 64, 65, 5, 6, 0, 0, 65, 66, 3, 6, 3, 0, 66, 13, 1, 0, 0, 0, 67, 69,
		5, 13, 0, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 71, 3, 16, 8, 0, 71, 15, 1, 0, 0, 0, 72, 73, 7, 0, 0, 0, 73, 17, 1,
		0, 0, 0, 74, 75, 7, 1, 0, 0, 75, 19, 1, 0, 0, 0, 76, 77, 7, 2, 0, 0, 77,
		21, 1, 0, 0, 0, 78, 85, 5, 29, 0, 0, 79, 80, 5, 12, 0, 0, 80, 82, 5, 29,
		0, 0, 81, 79, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 81, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 23, 1, 0, 0, 0, 9, 27, 39, 41, 46, 52, 59, 68, 83, 85,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// StcParserInit initializes any static state used to implement StcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StcParserInit() {
	staticData := &StcParserStaticData
	staticData.once.Do(stcParserInit)
}

// NewStcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStcParser(input antlr.TokenStream) *StcParser {
	StcParserInit()
	this := new(StcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &StcParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Stc.g4"

	return this
}

// StcParser tokens.
const (
	StcParserEOF                 = antlr.TokenEOF
	StcParserT__0                = 1
	StcParserT__1                = 2
	StcParserT__2                = 3
	StcParserT__3                = 4
	StcParserT__4                = 5
	StcParserT__5                = 6
	StcParserT__6                = 7
	StcParserT__7                = 8
	StcParserT__8                = 9
	StcParserT__9                = 10
	StcParserT__10               = 11
	StcParserT__11               = 12
	StcParserSTACK_PREVENTION    = 13
	StcParserNUMBER              = 14
	StcParserSIGNED_NUMBER       = 15
	StcParserFLOAT               = 16
	StcParserSIGNED_FLOAT        = 17
	StcParserFLOAT_E             = 18
	StcParserSIGNED_FLOAT_E      = 19
	StcParserBIN_NUMBER          = 20
	StcParserHEX_NUMBER          = 21
	StcParserSTRING              = 22
	StcParserCHAR                = 23
	StcParserBOOL                = 24
	StcParserSIMPLE_TYPE         = 25
	StcParserLOGIC_OPERATOR      = 26
	StcParserARITHMETIC_OPERATOR = 27
	StcParserBUILD_IN_OPERATOR   = 28
	StcParserID                  = 29
	StcParserWS                  = 30
)

// StcParser rules.
const (
	StcParserRULE_prog           = 0
	StcParserRULE_functionDef    = 1
	StcParserRULE_subBlock       = 2
	StcParserRULE_block          = 3
	StcParserRULE_ifBlock        = 4
	StcParserRULE_elseBlock      = 5
	StcParserRULE_repeatBlock    = 6
	StcParserRULE_operation      = 7
	StcParserRULE_operaor        = 8
	StcParserRULE_stackOperation = 9
	StcParserRULE_push           = 10
	StcParserRULE_identifier     = 11
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFunctionDef() []IFunctionDefContext
	FunctionDef(i int) IFunctionDefContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(StcParserEOF, 0)
}

func (s *ProgContext) AllFunctionDef() []IFunctionDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDefContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDefContext); ok {
			tst[i] = t.(IFunctionDefContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) FunctionDef(i int) IFunctionDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *StcParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StcParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == StcParserT__0 {
		{
			p.SetState(24)
			p.FunctionDef()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(29)
		p.Match(StcParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefContext is an interface to support dynamic dispatch.
type IFunctionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Block() IBlockContext

	// IsFunctionDefContext differentiates from other interfaces.
	IsFunctionDefContext()
}

type FunctionDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefContext() *FunctionDefContext {
	var p = new(FunctionDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_functionDef
	return p
}

func InitEmptyFunctionDefContext(p *FunctionDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_functionDef
}

func (*FunctionDefContext) IsFunctionDefContext() {}

func NewFunctionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefContext {
	var p = new(FunctionDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_functionDef

	return p
}

func (s *FunctionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefContext) ID() antlr.TerminalNode {
	return s.GetToken(StcParserID, 0)
}

func (s *FunctionDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterFunctionDef(s)
	}
}

func (s *FunctionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitFunctionDef(s)
	}
}

func (p *StcParser) FunctionDef() (localctx IFunctionDefContext) {
	localctx = NewFunctionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StcParserRULE_functionDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(StcParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubBlockContext is an interface to support dynamic dispatch.
type ISubBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperation() []IOperationContext
	Operation(i int) IOperationContext
	AllStackOperation() []IStackOperationContext
	StackOperation(i int) IStackOperationContext
	AllPush() []IPushContext
	Push(i int) IPushContext
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	FunctionDef() IFunctionDefContext
	IfBlock() IIfBlockContext
	RepeatBlock() IRepeatBlockContext

	// IsSubBlockContext differentiates from other interfaces.
	IsSubBlockContext()
}

type SubBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubBlockContext() *SubBlockContext {
	var p = new(SubBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_subBlock
	return p
}

func InitEmptySubBlockContext(p *SubBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_subBlock
}

func (*SubBlockContext) IsSubBlockContext() {}

func NewSubBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubBlockContext {
	var p = new(SubBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_subBlock

	return p
}

func (s *SubBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SubBlockContext) AllOperation() []IOperationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationContext); ok {
			len++
		}
	}

	tst := make([]IOperationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationContext); ok {
			tst[i] = t.(IOperationContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) Operation(i int) IOperationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *SubBlockContext) AllStackOperation() []IStackOperationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStackOperationContext); ok {
			len++
		}
	}

	tst := make([]IStackOperationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStackOperationContext); ok {
			tst[i] = t.(IStackOperationContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) StackOperation(i int) IStackOperationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStackOperationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStackOperationContext)
}

func (s *SubBlockContext) AllPush() []IPushContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPushContext); ok {
			len++
		}
	}

	tst := make([]IPushContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPushContext); ok {
			tst[i] = t.(IPushContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) Push(i int) IPushContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPushContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPushContext)
}

func (s *SubBlockContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SubBlockContext) FunctionDef() IFunctionDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *SubBlockContext) IfBlock() IIfBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *SubBlockContext) RepeatBlock() IRepeatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatBlockContext)
}

func (s *SubBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterSubBlock(s)
	}
}

func (s *SubBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitSubBlock(s)
	}
}

func (p *StcParser) SubBlock() (localctx ISubBlockContext) {
	localctx = NewSubBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, StcParserRULE_subBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case StcParserSTACK_PREVENTION, StcParserLOGIC_OPERATOR, StcParserARITHMETIC_OPERATOR, StcParserBUILD_IN_OPERATOR:
				{
					p.SetState(35)
					p.Operation()
				}

			case StcParserT__6, StcParserT__7, StcParserT__8, StcParserT__9, StcParserT__10:
				{
					p.SetState(36)
					p.StackOperation()
				}

			case StcParserNUMBER, StcParserSIGNED_NUMBER, StcParserFLOAT, StcParserSIGNED_FLOAT, StcParserFLOAT_E, StcParserSIGNED_FLOAT_E, StcParserBIN_NUMBER, StcParserHEX_NUMBER, StcParserSTRING, StcParserCHAR, StcParserBOOL, StcParserSIMPLE_TYPE:
				{
					p.SetState(37)
					p.Push()
				}

			case StcParserID:
				{
					p.SetState(38)
					p.Identifier()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case StcParserT__0:
		{
			p.SetState(43)
			p.FunctionDef()
		}

	case StcParserT__3:
		{
			p.SetState(44)
			p.IfBlock()
		}

	case StcParserT__5:
		{
			p.SetState(45)
			p.RepeatBlock()
		}

	case StcParserT__2, StcParserT__6, StcParserT__7, StcParserT__8, StcParserT__9, StcParserT__10, StcParserSTACK_PREVENTION, StcParserNUMBER, StcParserSIGNED_NUMBER, StcParserFLOAT, StcParserSIGNED_FLOAT, StcParserFLOAT_E, StcParserSIGNED_FLOAT_E, StcParserBIN_NUMBER, StcParserHEX_NUMBER, StcParserSTRING, StcParserCHAR, StcParserBOOL, StcParserSIMPLE_TYPE, StcParserLOGIC_OPERATOR, StcParserARITHMETIC_OPERATOR, StcParserBUILD_IN_OPERATOR, StcParserID:

	default:
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSubBlock() []ISubBlockContext
	SubBlock(i int) ISubBlockContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllSubBlock() []ISubBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubBlockContext); ok {
			len++
		}
	}

	tst := make([]ISubBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubBlockContext); ok {
			tst[i] = t.(ISubBlockContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) SubBlock(i int) ISubBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubBlockContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *StcParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, StcParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(StcParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073737600) != 0) {
		{
			p.SetState(49)
			p.SubBlock()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(StcParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	ElseBlock() IElseBlockContext

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_ifBlock
	return p
}

func InitEmptyIfBlockContext(p *IfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_ifBlock
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfBlockContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (p *StcParser) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, StcParserRULE_ifBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(StcParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Block()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__4 {
		{
			p.SetState(58)
			p.ElseBlock()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (p *StcParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, StcParserRULE_elseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(StcParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepeatBlockContext is an interface to support dynamic dispatch.
type IRepeatBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsRepeatBlockContext differentiates from other interfaces.
	IsRepeatBlockContext()
}

type RepeatBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatBlockContext() *RepeatBlockContext {
	var p = new(RepeatBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_repeatBlock
	return p
}

func InitEmptyRepeatBlockContext(p *RepeatBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_repeatBlock
}

func (*RepeatBlockContext) IsRepeatBlockContext() {}

func NewRepeatBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatBlockContext {
	var p = new(RepeatBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_repeatBlock

	return p
}

func (s *RepeatBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RepeatBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterRepeatBlock(s)
	}
}

func (s *RepeatBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitRepeatBlock(s)
	}
}

func (p *StcParser) RepeatBlock() (localctx IRepeatBlockContext) {
	localctx = NewRepeatBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, StcParserRULE_repeatBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(StcParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operaor() IOperaorContext
	STACK_PREVENTION() antlr.TerminalNode

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_operation
	return p
}

func InitEmptyOperationContext(p *OperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_operation
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationContext) Operaor() IOperaorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperaorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperaorContext)
}

func (s *OperationContext) STACK_PREVENTION() antlr.TerminalNode {
	return s.GetToken(StcParserSTACK_PREVENTION, 0)
}

func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (p *StcParser) Operation() (localctx IOperationContext) {
	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, StcParserRULE_operation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserSTACK_PREVENTION {
		{
			p.SetState(67)
			p.Match(StcParserSTACK_PREVENTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(70)
		p.Operaor()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperaorContext is an interface to support dynamic dispatch.
type IOperaorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOGIC_OPERATOR() antlr.TerminalNode
	ARITHMETIC_OPERATOR() antlr.TerminalNode
	BUILD_IN_OPERATOR() antlr.TerminalNode

	// IsOperaorContext differentiates from other interfaces.
	IsOperaorContext()
}

type OperaorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperaorContext() *OperaorContext {
	var p = new(OperaorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_operaor
	return p
}

func InitEmptyOperaorContext(p *OperaorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_operaor
}

func (*OperaorContext) IsOperaorContext() {}

func NewOperaorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperaorContext {
	var p = new(OperaorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_operaor

	return p
}

func (s *OperaorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperaorContext) LOGIC_OPERATOR() antlr.TerminalNode {
	return s.GetToken(StcParserLOGIC_OPERATOR, 0)
}

func (s *OperaorContext) ARITHMETIC_OPERATOR() antlr.TerminalNode {
	return s.GetToken(StcParserARITHMETIC_OPERATOR, 0)
}

func (s *OperaorContext) BUILD_IN_OPERATOR() antlr.TerminalNode {
	return s.GetToken(StcParserBUILD_IN_OPERATOR, 0)
}

func (s *OperaorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperaorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperaorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterOperaor(s)
	}
}

func (s *OperaorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitOperaor(s)
	}
}

func (p *StcParser) Operaor() (localctx IOperaorContext) {
	localctx = NewOperaorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, StcParserRULE_operaor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&469762048) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStackOperationContext is an interface to support dynamic dispatch.
type IStackOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStackOperationContext differentiates from other interfaces.
	IsStackOperationContext()
}

type StackOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStackOperationContext() *StackOperationContext {
	var p = new(StackOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_stackOperation
	return p
}

func InitEmptyStackOperationContext(p *StackOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_stackOperation
}

func (*StackOperationContext) IsStackOperationContext() {}

func NewStackOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StackOperationContext {
	var p = new(StackOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_stackOperation

	return p
}

func (s *StackOperationContext) GetParser() antlr.Parser { return s.parser }
func (s *StackOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StackOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StackOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterStackOperation(s)
	}
}

func (s *StackOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitStackOperation(s)
	}
}

func (p *StcParser) StackOperation() (localctx IStackOperationContext) {
	localctx = NewStackOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, StcParserRULE_stackOperation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3968) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPushContext is an interface to support dynamic dispatch.
type IPushContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIGNED_NUMBER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	SIGNED_FLOAT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	SIGNED_FLOAT_E() antlr.TerminalNode
	FLOAT_E() antlr.TerminalNode
	BIN_NUMBER() antlr.TerminalNode
	HEX_NUMBER() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	SIMPLE_TYPE() antlr.TerminalNode

	// IsPushContext differentiates from other interfaces.
	IsPushContext()
}

type PushContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPushContext() *PushContext {
	var p = new(PushContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_push
	return p
}

func InitEmptyPushContext(p *PushContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_push
}

func (*PushContext) IsPushContext() {}

func NewPushContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PushContext {
	var p = new(PushContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_push

	return p
}

func (s *PushContext) GetParser() antlr.Parser { return s.parser }

func (s *PushContext) SIGNED_NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserSIGNED_NUMBER, 0)
}

func (s *PushContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserNUMBER, 0)
}

func (s *PushContext) SIGNED_FLOAT() antlr.TerminalNode {
	return s.GetToken(StcParserSIGNED_FLOAT, 0)
}

func (s *PushContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(StcParserFLOAT, 0)
}

func (s *PushContext) SIGNED_FLOAT_E() antlr.TerminalNode {
	return s.GetToken(StcParserSIGNED_FLOAT_E, 0)
}

func (s *PushContext) FLOAT_E() antlr.TerminalNode {
	return s.GetToken(StcParserFLOAT_E, 0)
}

func (s *PushContext) BIN_NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserBIN_NUMBER, 0)
}

func (s *PushContext) HEX_NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserHEX_NUMBER, 0)
}

func (s *PushContext) CHAR() antlr.TerminalNode {
	return s.GetToken(StcParserCHAR, 0)
}

func (s *PushContext) STRING() antlr.TerminalNode {
	return s.GetToken(StcParserSTRING, 0)
}

func (s *PushContext) BOOL() antlr.TerminalNode {
	return s.GetToken(StcParserBOOL, 0)
}

func (s *PushContext) SIMPLE_TYPE() antlr.TerminalNode {
	return s.GetToken(StcParserSIMPLE_TYPE, 0)
}

func (s *PushContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PushContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PushContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterPush(s)
	}
}

func (s *PushContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitPush(s)
	}
}

func (p *StcParser) Push() (localctx IPushContext) {
	localctx = NewPushContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, StcParserRULE_push)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67092480) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(StcParserID)
}

func (s *IdentifierContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(StcParserID, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *StcParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, StcParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__11 {
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == StcParserT__11 {
			{
				p.SetState(79)
				p.Match(StcParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(80)
				p.Match(StcParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
