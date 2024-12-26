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
		"", "'fun'", "'{'", "'}'", "'if'", "'else'", "'repeat'", "'('", "')'",
		"'dup'", "'rot'", "'swap'", "'pop'", "'over'", "':'", "'!'", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "':='", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STACK_PREVENTION",
		"NUMBER", "SIGNED_NUMBER", "FLOAT", "SIGNED_FLOAT", "FLOAT_E", "SIGNED_FLOAT_E",
		"BIN_NUMBER", "HEX_NUMBER", "STRING", "CHAR", "BOOL", "SIMPLE_TYPE",
		"LOGIC_OPERATOR", "ARITHMETIC_OPERATOR", "ASSIGN_OPERATOR", "REFERENCE_OPERATOR",
		"BUILD_IN_OPERATOR", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "functionDef", "subBlock", "block", "ifBlock", "elseBlock",
		"repeatBlock", "arguments", "operation", "operaor", "stackOperation",
		"push", "varAssing", "varReference", "varIdentifier", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 130, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 4, 0, 34, 8, 0, 11, 0, 12, 0, 35, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3,
		1, 43, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 53,
		8, 2, 11, 2, 12, 2, 54, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 3, 1, 3, 4,
		3, 64, 8, 3, 11, 3, 12, 3, 65, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 73,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 80, 8, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 4, 7, 86, 8, 7, 11, 7, 12, 7, 87, 1, 7, 1, 7, 1, 8, 3, 8, 93, 8, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 4, 14, 112, 8, 14, 11, 14, 12,
		14, 113, 3, 14, 116, 8, 14, 1, 15, 3, 15, 119, 8, 15, 1, 15, 1, 15, 1,
		15, 4, 15, 124, 8, 15, 11, 15, 12, 15, 125, 3, 15, 128, 8, 15, 1, 15, 0,
		0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 3,
		2, 0, 28, 30, 32, 32, 1, 0, 9, 13, 1, 0, 16, 27, 134, 0, 33, 1, 0, 0, 0,
		2, 39, 1, 0, 0, 0, 4, 52, 1, 0, 0, 0, 6, 61, 1, 0, 0, 0, 8, 69, 1, 0, 0,
		0, 10, 74, 1, 0, 0, 0, 12, 77, 1, 0, 0, 0, 14, 83, 1, 0, 0, 0, 16, 92,
		1, 0, 0, 0, 18, 96, 1, 0, 0, 0, 20, 98, 1, 0, 0, 0, 22, 100, 1, 0, 0, 0,
		24, 102, 1, 0, 0, 0, 26, 105, 1, 0, 0, 0, 28, 108, 1, 0, 0, 0, 30, 118,
		1, 0, 0, 0, 32, 34, 3, 2, 1, 0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0,
		35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 5,
		0, 0, 1, 38, 1, 1, 0, 0, 0, 39, 40, 5, 1, 0, 0, 40, 42, 5, 33, 0, 0, 41,
		43, 3, 14, 7, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0,
		0, 0, 44, 45, 3, 6, 3, 0, 45, 3, 1, 0, 0, 0, 46, 53, 3, 16, 8, 0, 47, 53,
		3, 20, 10, 0, 48, 53, 3, 22, 11, 0, 49, 53, 3, 26, 13, 0, 50, 53, 3, 24,
		12, 0, 51, 53, 3, 30, 15, 0, 52, 46, 1, 0, 0, 0, 52, 47, 1, 0, 0, 0, 52,
		48, 1, 0, 0, 0, 52, 49, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 51, 1, 0, 0,
		0, 53, 54, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 59,
		1, 0, 0, 0, 56, 60, 3, 2, 1, 0, 57, 60, 3, 8, 4, 0, 58, 60, 3, 12, 6, 0,
		59, 56, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1,
		0, 0, 0, 60, 5, 1, 0, 0, 0, 61, 63, 5, 2, 0, 0, 62, 64, 3, 4, 2, 0, 63,
		62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0,
		0, 66, 67, 1, 0, 0, 0, 67, 68, 5, 3, 0, 0, 68, 7, 1, 0, 0, 0, 69, 70, 5,
		4, 0, 0, 70, 72, 3, 6, 3, 0, 71, 73, 3, 10, 5, 0, 72, 71, 1, 0, 0, 0, 72,
		73, 1, 0, 0, 0, 73, 9, 1, 0, 0, 0, 74, 75, 5, 5, 0, 0, 75, 76, 3, 6, 3,
		0, 76, 11, 1, 0, 0, 0, 77, 79, 5, 6, 0, 0, 78, 80, 3, 14, 7, 0, 79, 78,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 3, 6, 3, 0,
		82, 13, 1, 0, 0, 0, 83, 85, 5, 7, 0, 0, 84, 86, 3, 30, 15, 0, 85, 84, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 90, 5, 8, 0, 0, 90, 15, 1, 0, 0, 0, 91, 93, 5, 15,
		0, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95,
		3, 18, 9, 0, 95, 17, 1, 0, 0, 0, 96, 97, 7, 0, 0, 0, 97, 19, 1, 0, 0, 0,
		98, 99, 7, 1, 0, 0, 99, 21, 1, 0, 0, 0, 100, 101, 7, 2, 0, 0, 101, 23,
		1, 0, 0, 0, 102, 103, 3, 28, 14, 0, 103, 104, 5, 30, 0, 0, 104, 25, 1,
		0, 0, 0, 105, 106, 5, 31, 0, 0, 106, 107, 3, 28, 14, 0, 107, 27, 1, 0,
		0, 0, 108, 115, 5, 33, 0, 0, 109, 110, 5, 14, 0, 0, 110, 112, 5, 33, 0,
		0, 111, 109, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113,
		114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 111, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 29, 1, 0, 0, 0, 117, 119, 5, 15, 0, 0, 118, 117, 1, 0,
		0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 127, 5, 33, 0, 0,
		121, 122, 5, 14, 0, 0, 122, 124, 5, 33, 0, 0, 123, 121, 1, 0, 0, 0, 124,
		125, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 128,
		1, 0, 0, 0, 127, 123, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 31, 1, 0,
		0, 0, 15, 35, 42, 52, 54, 59, 65, 72, 79, 87, 92, 113, 115, 118, 125, 127,
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
	StcParserT__12               = 13
	StcParserT__13               = 14
	StcParserSTACK_PREVENTION    = 15
	StcParserNUMBER              = 16
	StcParserSIGNED_NUMBER       = 17
	StcParserFLOAT               = 18
	StcParserSIGNED_FLOAT        = 19
	StcParserFLOAT_E             = 20
	StcParserSIGNED_FLOAT_E      = 21
	StcParserBIN_NUMBER          = 22
	StcParserHEX_NUMBER          = 23
	StcParserSTRING              = 24
	StcParserCHAR                = 25
	StcParserBOOL                = 26
	StcParserSIMPLE_TYPE         = 27
	StcParserLOGIC_OPERATOR      = 28
	StcParserARITHMETIC_OPERATOR = 29
	StcParserASSIGN_OPERATOR     = 30
	StcParserREFERENCE_OPERATOR  = 31
	StcParserBUILD_IN_OPERATOR   = 32
	StcParserID                  = 33
	StcParserWS                  = 34
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
	StcParserRULE_arguments      = 7
	StcParserRULE_operation      = 8
	StcParserRULE_operaor        = 9
	StcParserRULE_stackOperation = 10
	StcParserRULE_push           = 11
	StcParserRULE_varAssing      = 12
	StcParserRULE_varReference   = 13
	StcParserRULE_varIdentifier  = 14
	StcParserRULE_identifier     = 15
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == StcParserT__0 {
		{
			p.SetState(32)
			p.FunctionDef()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
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
	Arguments() IArgumentsContext

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

func (s *FunctionDefContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(StcParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__6 {
		{
			p.SetState(41)
			p.Arguments()
		}

	}
	{
		p.SetState(44)
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
	AllVarReference() []IVarReferenceContext
	VarReference(i int) IVarReferenceContext
	AllVarAssing() []IVarAssingContext
	VarAssing(i int) IVarAssingContext
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

func (s *SubBlockContext) AllVarReference() []IVarReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarReferenceContext); ok {
			len++
		}
	}

	tst := make([]IVarReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarReferenceContext); ok {
			tst[i] = t.(IVarReferenceContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) VarReference(i int) IVarReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarReferenceContext); ok {
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

	return t.(IVarReferenceContext)
}

func (s *SubBlockContext) AllVarAssing() []IVarAssingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarAssingContext); ok {
			len++
		}
	}

	tst := make([]IVarAssingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarAssingContext); ok {
			tst[i] = t.(IVarAssingContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) VarAssing(i int) IVarAssingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAssingContext); ok {
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

	return t.(IVarAssingContext)
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
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(46)
					p.Operation()
				}

			case 2:
				{
					p.SetState(47)
					p.StackOperation()
				}

			case 3:
				{
					p.SetState(48)
					p.Push()
				}

			case 4:
				{
					p.SetState(49)
					p.VarReference()
				}

			case 5:
				{
					p.SetState(50)
					p.VarAssing()
				}

			case 6:
				{
					p.SetState(51)
					p.Identifier()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case StcParserT__0:
		{
			p.SetState(56)
			p.FunctionDef()
		}

	case StcParserT__3:
		{
			p.SetState(57)
			p.IfBlock()
		}

	case StcParserT__5:
		{
			p.SetState(58)
			p.RepeatBlock()
		}

	case StcParserT__2, StcParserT__8, StcParserT__9, StcParserT__10, StcParserT__11, StcParserT__12, StcParserSTACK_PREVENTION, StcParserNUMBER, StcParserSIGNED_NUMBER, StcParserFLOAT, StcParserSIGNED_FLOAT, StcParserFLOAT_E, StcParserSIGNED_FLOAT_E, StcParserBIN_NUMBER, StcParserHEX_NUMBER, StcParserSTRING, StcParserCHAR, StcParserBOOL, StcParserSIMPLE_TYPE, StcParserLOGIC_OPERATOR, StcParserARITHMETIC_OPERATOR, StcParserASSIGN_OPERATOR, StcParserREFERENCE_OPERATOR, StcParserBUILD_IN_OPERATOR, StcParserID:

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
		p.SetState(61)
		p.Match(StcParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17179852288) != 0) {
		{
			p.SetState(62)
			p.SubBlock()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
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
		p.SetState(69)
		p.Match(StcParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Block()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__4 {
		{
			p.SetState(71)
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
		p.SetState(74)
		p.Match(StcParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
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
	Arguments() IArgumentsContext

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

func (s *RepeatBlockContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(StcParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__6 {
		{
			p.SetState(78)
			p.Arguments()
		}

	}
	{
		p.SetState(81)
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllIdentifier() []IIdentifierContext {
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

func (s *ArgumentsContext) Identifier(i int) IIdentifierContext {
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

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *StcParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, StcParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(StcParserT__6)
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

	for ok := true; ok; ok = _la == StcParserSTACK_PREVENTION || _la == StcParserID {
		{
			p.SetState(84)
			p.Identifier()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.Match(StcParserT__7)
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
	p.EnterRule(localctx, 16, StcParserRULE_operation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserSTACK_PREVENTION {
		{
			p.SetState(91)
			p.Match(StcParserSTACK_PREVENTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(94)
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
	ASSIGN_OPERATOR() antlr.TerminalNode

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

func (s *OperaorContext) ASSIGN_OPERATOR() antlr.TerminalNode {
	return s.GetToken(StcParserASSIGN_OPERATOR, 0)
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
	p.EnterRule(localctx, 18, StcParserRULE_operaor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6174015488) != 0) {
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
	p.EnterRule(localctx, 20, StcParserRULE_stackOperation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15872) != 0) {
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
	p.EnterRule(localctx, 22, StcParserRULE_push)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268369920) != 0) {
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

// IVarAssingContext is an interface to support dynamic dispatch.
type IVarAssingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarIdentifier() IVarIdentifierContext
	ASSIGN_OPERATOR() antlr.TerminalNode

	// IsVarAssingContext differentiates from other interfaces.
	IsVarAssingContext()
}

type VarAssingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssingContext() *VarAssingContext {
	var p = new(VarAssingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varAssing
	return p
}

func InitEmptyVarAssingContext(p *VarAssingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varAssing
}

func (*VarAssingContext) IsVarAssingContext() {}

func NewVarAssingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssingContext {
	var p = new(VarAssingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_varAssing

	return p
}

func (s *VarAssingContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssingContext) VarIdentifier() IVarIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarIdentifierContext)
}

func (s *VarAssingContext) ASSIGN_OPERATOR() antlr.TerminalNode {
	return s.GetToken(StcParserASSIGN_OPERATOR, 0)
}

func (s *VarAssingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAssingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterVarAssing(s)
	}
}

func (s *VarAssingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitVarAssing(s)
	}
}

func (p *StcParser) VarAssing() (localctx IVarAssingContext) {
	localctx = NewVarAssingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, StcParserRULE_varAssing)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.VarIdentifier()
	}
	{
		p.SetState(103)
		p.Match(StcParserASSIGN_OPERATOR)
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

// IVarReferenceContext is an interface to support dynamic dispatch.
type IVarReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REFERENCE_OPERATOR() antlr.TerminalNode
	VarIdentifier() IVarIdentifierContext

	// IsVarReferenceContext differentiates from other interfaces.
	IsVarReferenceContext()
}

type VarReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarReferenceContext() *VarReferenceContext {
	var p = new(VarReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varReference
	return p
}

func InitEmptyVarReferenceContext(p *VarReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varReference
}

func (*VarReferenceContext) IsVarReferenceContext() {}

func NewVarReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarReferenceContext {
	var p = new(VarReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_varReference

	return p
}

func (s *VarReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *VarReferenceContext) REFERENCE_OPERATOR() antlr.TerminalNode {
	return s.GetToken(StcParserREFERENCE_OPERATOR, 0)
}

func (s *VarReferenceContext) VarIdentifier() IVarIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarIdentifierContext)
}

func (s *VarReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterVarReference(s)
	}
}

func (s *VarReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitVarReference(s)
	}
}

func (p *StcParser) VarReference() (localctx IVarReferenceContext) {
	localctx = NewVarReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, StcParserRULE_varReference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(StcParserREFERENCE_OPERATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.VarIdentifier()
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

// IVarIdentifierContext is an interface to support dynamic dispatch.
type IVarIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsVarIdentifierContext differentiates from other interfaces.
	IsVarIdentifierContext()
}

type VarIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarIdentifierContext() *VarIdentifierContext {
	var p = new(VarIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varIdentifier
	return p
}

func InitEmptyVarIdentifierContext(p *VarIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varIdentifier
}

func (*VarIdentifierContext) IsVarIdentifierContext() {}

func NewVarIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarIdentifierContext {
	var p = new(VarIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_varIdentifier

	return p
}

func (s *VarIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VarIdentifierContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(StcParserID)
}

func (s *VarIdentifierContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(StcParserID, i)
}

func (s *VarIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterVarIdentifier(s)
	}
}

func (s *VarIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitVarIdentifier(s)
	}
}

func (p *StcParser) VarIdentifier() (localctx IVarIdentifierContext) {
	localctx = NewVarIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, StcParserRULE_varIdentifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__13 {
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == StcParserT__13 {
			{
				p.SetState(109)
				p.Match(StcParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(110)
				p.Match(StcParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(113)
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	STACK_PREVENTION() antlr.TerminalNode

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

func (s *IdentifierContext) STACK_PREVENTION() antlr.TerminalNode {
	return s.GetToken(StcParserSTACK_PREVENTION, 0)
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
	p.EnterRule(localctx, 30, StcParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserSTACK_PREVENTION {
		{
			p.SetState(117)
			p.Match(StcParserSTACK_PREVENTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(120)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__13 {
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == StcParserT__13 {
			{
				p.SetState(121)
				p.Match(StcParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(122)
				p.Match(StcParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(125)
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
