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
		"", "'fun'", "'<'", "'>'", "'{'", "'}'", "':'", "'if'", "'else'", "'repeat'",
		"'('", "')'", "'dup'", "'rot'", "'swap'", "'pop'", "'over'", "'c:'",
		"'new'", "'struct'", "'!'", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "':='", "'&'", "", "'['", "']'", "'arr'", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "NEW", "STRUCT_DEFINITION", "STACK_PREVENTION", "NUMBER", "SIGNED_NUMBER",
		"FLOAT", "SIGNED_FLOAT", "FLOAT_E", "SIGNED_FLOAT_E", "BIN_NUMBER",
		"HEX_NUMBER", "STRING", "CHAR", "BOOL", "SIMPLE_TYPE", "LOGIC_OPERATOR",
		"ARITHMETIC_OPERATOR", "ASSIGN_OPERATOR", "REFERENCE_OPERATOR", "BUILD_IN_OPERATOR",
		"ARRAY_OPEN", "ARRAY_CLOSE", "ARRAY_OPERATOR", "AT", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "functionDef", "type", "subBlock", "block", "newOperator", "struct",
		"structElement", "structBody", "ifBlock", "elseBlock", "repeatBlock",
		"arguments", "operation", "operaor", "stackOperation", "push", "arrayElement",
		"arrayBase", "arrayIndexShift", "arrayIndex", "capacity", "array", "arrayDescriber",
		"arrayNew", "argument", "varAssign", "varAssignIdentifier", "varType",
		"varReference", "varIdentifier", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 259, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 1, 0, 4, 0, 67, 8, 0, 11, 0, 12, 0, 68, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 3, 1, 76, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 86, 8, 2, 1, 2, 3, 2, 89, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 101, 8, 3, 11, 3, 12, 3, 102, 1, 3, 1,
		3, 1, 3, 3, 3, 108, 8, 3, 1, 4, 1, 4, 4, 4, 112, 8, 4, 11, 4, 12, 4, 113,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 4, 8, 131, 8, 8, 11, 8, 12, 8, 132, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 3, 9, 140, 8, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11,
		147, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 4, 12, 153, 8, 12, 11, 12, 12,
		12, 154, 1, 12, 1, 12, 1, 13, 3, 13, 160, 8, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 185,
		8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 3, 19, 191, 8, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 4, 22, 202, 8, 22, 11, 22,
		12, 22, 203, 1, 22, 1, 22, 3, 22, 208, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 3, 24, 217, 8, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 3, 26, 224, 8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 4, 30, 241, 8,
		30, 11, 30, 12, 30, 242, 3, 30, 245, 8, 30, 1, 31, 3, 31, 248, 8, 31, 1,
		31, 1, 31, 1, 31, 4, 31, 253, 8, 31, 11, 31, 12, 31, 254, 3, 31, 257, 8,
		31, 1, 31, 0, 0, 32, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		0, 3, 2, 0, 33, 35, 37, 37, 1, 0, 12, 16, 1, 0, 21, 32, 276, 0, 66, 1,
		0, 0, 0, 2, 72, 1, 0, 0, 0, 4, 88, 1, 0, 0, 0, 6, 100, 1, 0, 0, 0, 8, 109,
		1, 0, 0, 0, 10, 117, 1, 0, 0, 0, 12, 121, 1, 0, 0, 0, 14, 125, 1, 0, 0,
		0, 16, 128, 1, 0, 0, 0, 18, 136, 1, 0, 0, 0, 20, 141, 1, 0, 0, 0, 22, 144,
		1, 0, 0, 0, 24, 150, 1, 0, 0, 0, 26, 159, 1, 0, 0, 0, 28, 163, 1, 0, 0,
		0, 30, 165, 1, 0, 0, 0, 32, 167, 1, 0, 0, 0, 34, 184, 1, 0, 0, 0, 36, 186,
		1, 0, 0, 0, 38, 190, 1, 0, 0, 0, 40, 192, 1, 0, 0, 0, 42, 196, 1, 0, 0,
		0, 44, 199, 1, 0, 0, 0, 46, 209, 1, 0, 0, 0, 48, 213, 1, 0, 0, 0, 50, 218,
		1, 0, 0, 0, 52, 223, 1, 0, 0, 0, 54, 227, 1, 0, 0, 0, 56, 230, 1, 0, 0,
		0, 58, 234, 1, 0, 0, 0, 60, 237, 1, 0, 0, 0, 62, 247, 1, 0, 0, 0, 64, 67,
		3, 2, 1, 0, 65, 67, 3, 12, 6, 0, 66, 64, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0,
		67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 1,
		0, 0, 0, 70, 71, 5, 0, 0, 1, 71, 1, 1, 0, 0, 0, 72, 73, 5, 1, 0, 0, 73,
		75, 5, 42, 0, 0, 74, 76, 3, 24, 12, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1, 0,
		0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 3, 8, 4, 0, 78, 3, 1, 0, 0, 0, 79, 89,
		5, 32, 0, 0, 80, 89, 5, 42, 0, 0, 81, 82, 5, 32, 0, 0, 82, 85, 5, 2, 0,
		0, 83, 86, 3, 4, 2, 0, 84, 86, 5, 42, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84,
		1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 89, 5, 3, 0, 0, 88, 79, 1, 0, 0, 0,
		88, 80, 1, 0, 0, 0, 88, 81, 1, 0, 0, 0, 89, 5, 1, 0, 0, 0, 90, 101, 3,
		10, 5, 0, 91, 101, 3, 58, 29, 0, 92, 101, 3, 52, 26, 0, 93, 101, 3, 40,
		20, 0, 94, 101, 3, 48, 24, 0, 95, 101, 3, 44, 22, 0, 96, 101, 3, 26, 13,
		0, 97, 101, 3, 30, 15, 0, 98, 101, 3, 32, 16, 0, 99, 101, 3, 62, 31, 0,
		100, 90, 1, 0, 0, 0, 100, 91, 1, 0, 0, 0, 100, 92, 1, 0, 0, 0, 100, 93,
		1, 0, 0, 0, 100, 94, 1, 0, 0, 0, 100, 95, 1, 0, 0, 0, 100, 96, 1, 0, 0,
		0, 100, 97, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101,
		102, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 107,
		1, 0, 0, 0, 104, 108, 3, 2, 1, 0, 105, 108, 3, 18, 9, 0, 106, 108, 3, 22,
		11, 0, 107, 104, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0,
		107, 108, 1, 0, 0, 0, 108, 7, 1, 0, 0, 0, 109, 111, 5, 4, 0, 0, 110, 112,
		3, 6, 3, 0, 111, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 111, 1, 0,
		0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 5, 5, 0, 0,
		116, 9, 1, 0, 0, 0, 117, 118, 5, 18, 0, 0, 118, 119, 5, 6, 0, 0, 119, 120,
		5, 42, 0, 0, 120, 11, 1, 0, 0, 0, 121, 122, 5, 19, 0, 0, 122, 123, 5, 42,
		0, 0, 123, 124, 3, 16, 8, 0, 124, 13, 1, 0, 0, 0, 125, 126, 5, 42, 0, 0,
		126, 127, 3, 56, 28, 0, 127, 15, 1, 0, 0, 0, 128, 130, 5, 4, 0, 0, 129,
		131, 3, 14, 7, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 130,
		1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 5, 5,
		0, 0, 135, 17, 1, 0, 0, 0, 136, 137, 5, 7, 0, 0, 137, 139, 3, 8, 4, 0,
		138, 140, 3, 20, 10, 0, 139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		19, 1, 0, 0, 0, 141, 142, 5, 8, 0, 0, 142, 143, 3, 8, 4, 0, 143, 21, 1,
		0, 0, 0, 144, 146, 5, 9, 0, 0, 145, 147, 3, 24, 12, 0, 146, 145, 1, 0,
		0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 3, 8, 4, 0,
		149, 23, 1, 0, 0, 0, 150, 152, 5, 10, 0, 0, 151, 153, 3, 50, 25, 0, 152,
		151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155,
		1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 5, 11, 0, 0, 157, 25, 1, 0,
		0, 0, 158, 160, 5, 20, 0, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0,
		160, 161, 1, 0, 0, 0, 161, 162, 3, 28, 14, 0, 162, 27, 1, 0, 0, 0, 163,
		164, 7, 0, 0, 0, 164, 29, 1, 0, 0, 0, 165, 166, 7, 1, 0, 0, 166, 31, 1,
		0, 0, 0, 167, 168, 7, 2, 0, 0, 168, 33, 1, 0, 0, 0, 169, 185, 5, 22, 0,
		0, 170, 185, 5, 21, 0, 0, 171, 185, 5, 24, 0, 0, 172, 185, 5, 23, 0, 0,
		173, 185, 5, 26, 0, 0, 174, 185, 5, 25, 0, 0, 175, 185, 5, 27, 0, 0, 176,
		185, 5, 28, 0, 0, 177, 185, 5, 30, 0, 0, 178, 185, 5, 29, 0, 0, 179, 185,
		5, 31, 0, 0, 180, 185, 5, 32, 0, 0, 181, 185, 3, 40, 20, 0, 182, 185, 3,
		60, 30, 0, 183, 185, 3, 44, 22, 0, 184, 169, 1, 0, 0, 0, 184, 170, 1, 0,
		0, 0, 184, 171, 1, 0, 0, 0, 184, 172, 1, 0, 0, 0, 184, 173, 1, 0, 0, 0,
		184, 174, 1, 0, 0, 0, 184, 175, 1, 0, 0, 0, 184, 176, 1, 0, 0, 0, 184,
		177, 1, 0, 0, 0, 184, 178, 1, 0, 0, 0, 184, 179, 1, 0, 0, 0, 184, 180,
		1, 0, 0, 0, 184, 181, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 183, 1, 0,
		0, 0, 185, 35, 1, 0, 0, 0, 186, 187, 3, 60, 30, 0, 187, 37, 1, 0, 0, 0,
		188, 191, 5, 21, 0, 0, 189, 191, 3, 60, 30, 0, 190, 188, 1, 0, 0, 0, 190,
		189, 1, 0, 0, 0, 191, 39, 1, 0, 0, 0, 192, 193, 3, 36, 18, 0, 193, 194,
		5, 41, 0, 0, 194, 195, 3, 38, 19, 0, 195, 41, 1, 0, 0, 0, 196, 197, 5,
		17, 0, 0, 197, 198, 5, 21, 0, 0, 198, 43, 1, 0, 0, 0, 199, 201, 5, 38,
		0, 0, 200, 202, 3, 34, 17, 0, 201, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0,
		0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205,
		207, 5, 39, 0, 0, 206, 208, 3, 42, 21, 0, 207, 206, 1, 0, 0, 0, 207, 208,
		1, 0, 0, 0, 208, 45, 1, 0, 0, 0, 209, 210, 5, 38, 0, 0, 210, 211, 5, 21,
		0, 0, 211, 212, 5, 39, 0, 0, 212, 47, 1, 0, 0, 0, 213, 214, 5, 40, 0, 0,
		214, 216, 3, 46, 23, 0, 215, 217, 3, 46, 23, 0, 216, 215, 1, 0, 0, 0, 216,
		217, 1, 0, 0, 0, 217, 49, 1, 0, 0, 0, 218, 219, 5, 42, 0, 0, 219, 51, 1,
		0, 0, 0, 220, 224, 3, 54, 27, 0, 221, 224, 3, 40, 20, 0, 222, 224, 3, 62,
		31, 0, 223, 220, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 222, 1, 0, 0, 0,
		224, 225, 1, 0, 0, 0, 225, 226, 5, 35, 0, 0, 226, 53, 1, 0, 0, 0, 227,
		228, 3, 60, 30, 0, 228, 229, 3, 56, 28, 0, 229, 55, 1, 0, 0, 0, 230, 231,
		5, 2, 0, 0, 231, 232, 3, 4, 2, 0, 232, 233, 5, 3, 0, 0, 233, 57, 1, 0,
		0, 0, 234, 235, 5, 36, 0, 0, 235, 236, 3, 60, 30, 0, 236, 59, 1, 0, 0,
		0, 237, 244, 5, 42, 0, 0, 238, 239, 5, 6, 0, 0, 239, 241, 5, 42, 0, 0,
		240, 238, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242,
		243, 1, 0, 0, 0, 243, 245, 1, 0, 0, 0, 244, 240, 1, 0, 0, 0, 244, 245,
		1, 0, 0, 0, 245, 61, 1, 0, 0, 0, 246, 248, 5, 20, 0, 0, 247, 246, 1, 0,
		0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 256, 5, 42, 0, 0,
		250, 251, 5, 6, 0, 0, 251, 253, 5, 42, 0, 0, 252, 250, 1, 0, 0, 0, 253,
		254, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257,
		1, 0, 0, 0, 256, 252, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 63, 1, 0,
		0, 0, 25, 66, 68, 75, 85, 88, 100, 102, 107, 113, 132, 139, 146, 154, 159,
		184, 190, 203, 207, 216, 223, 242, 244, 247, 254, 256,
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
	StcParserT__14               = 15
	StcParserT__15               = 16
	StcParserT__16               = 17
	StcParserNEW                 = 18
	StcParserSTRUCT_DEFINITION   = 19
	StcParserSTACK_PREVENTION    = 20
	StcParserNUMBER              = 21
	StcParserSIGNED_NUMBER       = 22
	StcParserFLOAT               = 23
	StcParserSIGNED_FLOAT        = 24
	StcParserFLOAT_E             = 25
	StcParserSIGNED_FLOAT_E      = 26
	StcParserBIN_NUMBER          = 27
	StcParserHEX_NUMBER          = 28
	StcParserSTRING              = 29
	StcParserCHAR                = 30
	StcParserBOOL                = 31
	StcParserSIMPLE_TYPE         = 32
	StcParserLOGIC_OPERATOR      = 33
	StcParserARITHMETIC_OPERATOR = 34
	StcParserASSIGN_OPERATOR     = 35
	StcParserREFERENCE_OPERATOR  = 36
	StcParserBUILD_IN_OPERATOR   = 37
	StcParserARRAY_OPEN          = 38
	StcParserARRAY_CLOSE         = 39
	StcParserARRAY_OPERATOR      = 40
	StcParserAT                  = 41
	StcParserID                  = 42
	StcParserWS                  = 43
)

// StcParser rules.
const (
	StcParserRULE_prog                = 0
	StcParserRULE_functionDef         = 1
	StcParserRULE_type                = 2
	StcParserRULE_subBlock            = 3
	StcParserRULE_block               = 4
	StcParserRULE_newOperator         = 5
	StcParserRULE_struct              = 6
	StcParserRULE_structElement       = 7
	StcParserRULE_structBody          = 8
	StcParserRULE_ifBlock             = 9
	StcParserRULE_elseBlock           = 10
	StcParserRULE_repeatBlock         = 11
	StcParserRULE_arguments           = 12
	StcParserRULE_operation           = 13
	StcParserRULE_operaor             = 14
	StcParserRULE_stackOperation      = 15
	StcParserRULE_push                = 16
	StcParserRULE_arrayElement        = 17
	StcParserRULE_arrayBase           = 18
	StcParserRULE_arrayIndexShift     = 19
	StcParserRULE_arrayIndex          = 20
	StcParserRULE_capacity            = 21
	StcParserRULE_array               = 22
	StcParserRULE_arrayDescriber      = 23
	StcParserRULE_arrayNew            = 24
	StcParserRULE_argument            = 25
	StcParserRULE_varAssign           = 26
	StcParserRULE_varAssignIdentifier = 27
	StcParserRULE_varType             = 28
	StcParserRULE_varReference        = 29
	StcParserRULE_varIdentifier       = 30
	StcParserRULE_identifier          = 31
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
	AllStruct_() []IStructContext
	Struct_(i int) IStructContext

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

func (s *ProgContext) AllStruct_() []IStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructContext); ok {
			len++
		}
	}

	tst := make([]IStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructContext); ok {
			tst[i] = t.(IStructContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Struct_(i int) IStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructContext); ok {
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

	return t.(IStructContext)
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == StcParserT__0 || _la == StcParserSTRUCT_DEFINITION {
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case StcParserT__0:
			{
				p.SetState(64)
				p.FunctionDef()
			}

		case StcParserSTRUCT_DEFINITION:
			{
				p.SetState(65)
				p.Struct_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
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
		p.SetState(72)
		p.Match(StcParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__9 {
		{
			p.SetState(74)
			p.Arguments()
		}

	}
	{
		p.SetState(77)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIMPLE_TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode
	Type_() ITypeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) SIMPLE_TYPE() antlr.TerminalNode {
	return s.GetToken(StcParserSIMPLE_TYPE, 0)
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(StcParserID, 0)
}

func (s *TypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *StcParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, StcParserRULE_type)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(StcParserSIMPLE_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(StcParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.Match(StcParserSIMPLE_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(StcParserT__1)
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

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(83)
				p.Type_()
			}

		case 2:
			{
				p.SetState(84)
				p.Match(StcParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(87)
			p.Match(StcParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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
	AllNewOperator() []INewOperatorContext
	NewOperator(i int) INewOperatorContext
	AllVarReference() []IVarReferenceContext
	VarReference(i int) IVarReferenceContext
	AllVarAssign() []IVarAssignContext
	VarAssign(i int) IVarAssignContext
	AllArrayIndex() []IArrayIndexContext
	ArrayIndex(i int) IArrayIndexContext
	AllArrayNew() []IArrayNewContext
	ArrayNew(i int) IArrayNewContext
	AllArray() []IArrayContext
	Array(i int) IArrayContext
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

func (s *SubBlockContext) AllNewOperator() []INewOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewOperatorContext); ok {
			len++
		}
	}

	tst := make([]INewOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewOperatorContext); ok {
			tst[i] = t.(INewOperatorContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) NewOperator(i int) INewOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewOperatorContext); ok {
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

	return t.(INewOperatorContext)
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

func (s *SubBlockContext) AllVarAssign() []IVarAssignContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarAssignContext); ok {
			len++
		}
	}

	tst := make([]IVarAssignContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarAssignContext); ok {
			tst[i] = t.(IVarAssignContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) VarAssign(i int) IVarAssignContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAssignContext); ok {
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

	return t.(IVarAssignContext)
}

func (s *SubBlockContext) AllArrayIndex() []IArrayIndexContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayIndexContext); ok {
			len++
		}
	}

	tst := make([]IArrayIndexContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayIndexContext); ok {
			tst[i] = t.(IArrayIndexContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) ArrayIndex(i int) IArrayIndexContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexContext); ok {
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

	return t.(IArrayIndexContext)
}

func (s *SubBlockContext) AllArrayNew() []IArrayNewContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayNewContext); ok {
			len++
		}
	}

	tst := make([]IArrayNewContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayNewContext); ok {
			tst[i] = t.(IArrayNewContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) ArrayNew(i int) IArrayNewContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayNewContext); ok {
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

	return t.(IArrayNewContext)
}

func (s *SubBlockContext) AllArray() []IArrayContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayContext); ok {
			len++
		}
	}

	tst := make([]IArrayContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayContext); ok {
			tst[i] = t.(IArrayContext)
			i++
		}
	}

	return tst
}

func (s *SubBlockContext) Array(i int) IArrayContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
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

	return t.(IArrayContext)
}

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
	p.EnterRule(localctx, 6, StcParserRULE_subBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(90)
					p.NewOperator()
				}

			case 2:
				{
					p.SetState(91)
					p.VarReference()
				}

			case 3:
				{
					p.SetState(92)
					p.VarAssign()
				}

			case 4:
				{
					p.SetState(93)
					p.ArrayIndex()
				}

			case 5:
				{
					p.SetState(94)
					p.ArrayNew()
				}

			case 6:
				{
					p.SetState(95)
					p.Array()
				}

			case 7:
				{
					p.SetState(96)
					p.Operation()
				}

			case 8:
				{
					p.SetState(97)
					p.StackOperation()
				}

			case 9:
				{
					p.SetState(98)
					p.Push()
				}

			case 10:
				{
					p.SetState(99)
					p.Identifier()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case StcParserT__0:
		{
			p.SetState(104)
			p.FunctionDef()
		}

	case StcParserT__6:
		{
			p.SetState(105)
			p.IfBlock()
		}

	case StcParserT__8:
		{
			p.SetState(106)
			p.RepeatBlock()
		}

	case StcParserT__4, StcParserT__11, StcParserT__12, StcParserT__13, StcParserT__14, StcParserT__15, StcParserNEW, StcParserSTACK_PREVENTION, StcParserNUMBER, StcParserSIGNED_NUMBER, StcParserFLOAT, StcParserSIGNED_FLOAT, StcParserFLOAT_E, StcParserSIGNED_FLOAT_E, StcParserBIN_NUMBER, StcParserHEX_NUMBER, StcParserSTRING, StcParserCHAR, StcParserBOOL, StcParserSIMPLE_TYPE, StcParserLOGIC_OPERATOR, StcParserARITHMETIC_OPERATOR, StcParserASSIGN_OPERATOR, StcParserREFERENCE_OPERATOR, StcParserBUILD_IN_OPERATOR, StcParserARRAY_OPEN, StcParserARRAY_OPERATOR, StcParserID:

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
	p.EnterRule(localctx, 8, StcParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(StcParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6047313293312) != 0) {
		{
			p.SetState(110)
			p.SubBlock()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(StcParserT__4)
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

// INewOperatorContext is an interface to support dynamic dispatch.
type INewOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEW() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsNewOperatorContext differentiates from other interfaces.
	IsNewOperatorContext()
}

type NewOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewOperatorContext() *NewOperatorContext {
	var p = new(NewOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_newOperator
	return p
}

func InitEmptyNewOperatorContext(p *NewOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_newOperator
}

func (*NewOperatorContext) IsNewOperatorContext() {}

func NewNewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewOperatorContext {
	var p = new(NewOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_newOperator

	return p
}

func (s *NewOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *NewOperatorContext) NEW() antlr.TerminalNode {
	return s.GetToken(StcParserNEW, 0)
}

func (s *NewOperatorContext) ID() antlr.TerminalNode {
	return s.GetToken(StcParserID, 0)
}

func (s *NewOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterNewOperator(s)
	}
}

func (s *NewOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitNewOperator(s)
	}
}

func (p *StcParser) NewOperator() (localctx INewOperatorContext) {
	localctx = NewNewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, StcParserRULE_newOperator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(StcParserNEW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(StcParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(StcParserID)
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

// IStructContext is an interface to support dynamic dispatch.
type IStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT_DEFINITION() antlr.TerminalNode
	ID() antlr.TerminalNode
	StructBody() IStructBodyContext

	// IsStructContext differentiates from other interfaces.
	IsStructContext()
}

type StructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructContext() *StructContext {
	var p = new(StructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_struct
	return p
}

func InitEmptyStructContext(p *StructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_struct
}

func (*StructContext) IsStructContext() {}

func NewStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructContext {
	var p = new(StructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_struct

	return p
}

func (s *StructContext) GetParser() antlr.Parser { return s.parser }

func (s *StructContext) STRUCT_DEFINITION() antlr.TerminalNode {
	return s.GetToken(StcParserSTRUCT_DEFINITION, 0)
}

func (s *StructContext) ID() antlr.TerminalNode {
	return s.GetToken(StcParserID, 0)
}

func (s *StructContext) StructBody() IStructBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (p *StcParser) Struct_() (localctx IStructContext) {
	localctx = NewStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, StcParserRULE_struct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(StcParserSTRUCT_DEFINITION)
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
	{
		p.SetState(123)
		p.StructBody()
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

// IStructElementContext is an interface to support dynamic dispatch.
type IStructElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	VarType() IVarTypeContext

	// IsStructElementContext differentiates from other interfaces.
	IsStructElementContext()
}

type StructElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructElementContext() *StructElementContext {
	var p = new(StructElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_structElement
	return p
}

func InitEmptyStructElementContext(p *StructElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_structElement
}

func (*StructElementContext) IsStructElementContext() {}

func NewStructElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructElementContext {
	var p = new(StructElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_structElement

	return p
}

func (s *StructElementContext) GetParser() antlr.Parser { return s.parser }

func (s *StructElementContext) ID() antlr.TerminalNode {
	return s.GetToken(StcParserID, 0)
}

func (s *StructElementContext) VarType() IVarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *StructElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterStructElement(s)
	}
}

func (s *StructElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitStructElement(s)
	}
}

func (p *StcParser) StructElement() (localctx IStructElementContext) {
	localctx = NewStructElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, StcParserRULE_structElement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.VarType()
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

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructElement() []IStructElementContext
	StructElement(i int) IStructElementContext

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_structBody
	return p
}

func InitEmptyStructBodyContext(p *StructBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_structBody
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) AllStructElement() []IStructElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructElementContext); ok {
			len++
		}
	}

	tst := make([]IStructElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructElementContext); ok {
			tst[i] = t.(IStructElementContext)
			i++
		}
	}

	return tst
}

func (s *StructBodyContext) StructElement(i int) IStructElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructElementContext); ok {
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

	return t.(IStructElementContext)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterStructBody(s)
	}
}

func (s *StructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitStructBody(s)
	}
}

func (p *StcParser) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, StcParserRULE_structBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(StcParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == StcParserID {
		{
			p.SetState(129)
			p.StructElement()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(StcParserT__4)
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
	p.EnterRule(localctx, 18, StcParserRULE_ifBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(StcParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Block()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__7 {
		{
			p.SetState(138)
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
	p.EnterRule(localctx, 20, StcParserRULE_elseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(StcParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
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
	p.EnterRule(localctx, 22, StcParserRULE_repeatBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(StcParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__9 {
		{
			p.SetState(145)
			p.Arguments()
		}

	}
	{
		p.SetState(148)
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
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext

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

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
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

	return t.(IArgumentContext)
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
	p.EnterRule(localctx, 24, StcParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(StcParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == StcParserID {
		{
			p.SetState(151)
			p.Argument()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(StcParserT__10)
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
	p.EnterRule(localctx, 26, StcParserRULE_operation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserSTACK_PREVENTION {
		{
			p.SetState(158)
			p.Match(StcParserSTACK_PREVENTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(161)
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
	p.EnterRule(localctx, 28, StcParserRULE_operaor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&197568495616) != 0) {
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
	p.EnterRule(localctx, 30, StcParserRULE_stackOperation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126976) != 0) {
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
	p.EnterRule(localctx, 32, StcParserRULE_push)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8587837440) != 0) {
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

// IArrayElementContext is an interface to support dynamic dispatch.
type IArrayElementContext interface {
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
	ArrayIndex() IArrayIndexContext
	VarIdentifier() IVarIdentifierContext
	Array() IArrayContext

	// IsArrayElementContext differentiates from other interfaces.
	IsArrayElementContext()
}

type ArrayElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayElementContext() *ArrayElementContext {
	var p = new(ArrayElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayElement
	return p
}

func InitEmptyArrayElementContext(p *ArrayElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayElement
}

func (*ArrayElementContext) IsArrayElementContext() {}

func NewArrayElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayElementContext {
	var p = new(ArrayElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_arrayElement

	return p
}

func (s *ArrayElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayElementContext) SIGNED_NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserSIGNED_NUMBER, 0)
}

func (s *ArrayElementContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserNUMBER, 0)
}

func (s *ArrayElementContext) SIGNED_FLOAT() antlr.TerminalNode {
	return s.GetToken(StcParserSIGNED_FLOAT, 0)
}

func (s *ArrayElementContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(StcParserFLOAT, 0)
}

func (s *ArrayElementContext) SIGNED_FLOAT_E() antlr.TerminalNode {
	return s.GetToken(StcParserSIGNED_FLOAT_E, 0)
}

func (s *ArrayElementContext) FLOAT_E() antlr.TerminalNode {
	return s.GetToken(StcParserFLOAT_E, 0)
}

func (s *ArrayElementContext) BIN_NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserBIN_NUMBER, 0)
}

func (s *ArrayElementContext) HEX_NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserHEX_NUMBER, 0)
}

func (s *ArrayElementContext) CHAR() antlr.TerminalNode {
	return s.GetToken(StcParserCHAR, 0)
}

func (s *ArrayElementContext) STRING() antlr.TerminalNode {
	return s.GetToken(StcParserSTRING, 0)
}

func (s *ArrayElementContext) BOOL() antlr.TerminalNode {
	return s.GetToken(StcParserBOOL, 0)
}

func (s *ArrayElementContext) SIMPLE_TYPE() antlr.TerminalNode {
	return s.GetToken(StcParserSIMPLE_TYPE, 0)
}

func (s *ArrayElementContext) ArrayIndex() IArrayIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayIndexContext)
}

func (s *ArrayElementContext) VarIdentifier() IVarIdentifierContext {
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

func (s *ArrayElementContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArrayElement(s)
	}
}

func (s *ArrayElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArrayElement(s)
	}
}

func (p *StcParser) ArrayElement() (localctx IArrayElementContext) {
	localctx = NewArrayElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, StcParserRULE_arrayElement)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(StcParserSIGNED_NUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Match(StcParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
			p.Match(StcParserSIGNED_FLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(172)
			p.Match(StcParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(173)
			p.Match(StcParserSIGNED_FLOAT_E)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(174)
			p.Match(StcParserFLOAT_E)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(175)
			p.Match(StcParserBIN_NUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(176)
			p.Match(StcParserHEX_NUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(177)
			p.Match(StcParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(178)
			p.Match(StcParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(179)
			p.Match(StcParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(180)
			p.Match(StcParserSIMPLE_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(181)
			p.ArrayIndex()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(182)
			p.VarIdentifier()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(183)
			p.Array()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IArrayBaseContext is an interface to support dynamic dispatch.
type IArrayBaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarIdentifier() IVarIdentifierContext

	// IsArrayBaseContext differentiates from other interfaces.
	IsArrayBaseContext()
}

type ArrayBaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayBaseContext() *ArrayBaseContext {
	var p = new(ArrayBaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayBase
	return p
}

func InitEmptyArrayBaseContext(p *ArrayBaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayBase
}

func (*ArrayBaseContext) IsArrayBaseContext() {}

func NewArrayBaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayBaseContext {
	var p = new(ArrayBaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_arrayBase

	return p
}

func (s *ArrayBaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayBaseContext) VarIdentifier() IVarIdentifierContext {
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

func (s *ArrayBaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayBaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayBaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArrayBase(s)
	}
}

func (s *ArrayBaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArrayBase(s)
	}
}

func (p *StcParser) ArrayBase() (localctx IArrayBaseContext) {
	localctx = NewArrayBaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, StcParserRULE_arrayBase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
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

// IArrayIndexShiftContext is an interface to support dynamic dispatch.
type IArrayIndexShiftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	VarIdentifier() IVarIdentifierContext

	// IsArrayIndexShiftContext differentiates from other interfaces.
	IsArrayIndexShiftContext()
}

type ArrayIndexShiftContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayIndexShiftContext() *ArrayIndexShiftContext {
	var p = new(ArrayIndexShiftContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayIndexShift
	return p
}

func InitEmptyArrayIndexShiftContext(p *ArrayIndexShiftContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayIndexShift
}

func (*ArrayIndexShiftContext) IsArrayIndexShiftContext() {}

func NewArrayIndexShiftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayIndexShiftContext {
	var p = new(ArrayIndexShiftContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_arrayIndexShift

	return p
}

func (s *ArrayIndexShiftContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayIndexShiftContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserNUMBER, 0)
}

func (s *ArrayIndexShiftContext) VarIdentifier() IVarIdentifierContext {
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

func (s *ArrayIndexShiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexShiftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayIndexShiftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArrayIndexShift(s)
	}
}

func (s *ArrayIndexShiftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArrayIndexShift(s)
	}
}

func (p *StcParser) ArrayIndexShift() (localctx IArrayIndexShiftContext) {
	localctx = NewArrayIndexShiftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, StcParserRULE_arrayIndexShift)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case StcParserNUMBER:
		{
			p.SetState(188)
			p.Match(StcParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StcParserID:
		{
			p.SetState(189)
			p.VarIdentifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IArrayIndexContext is an interface to support dynamic dispatch.
type IArrayIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArrayBase() IArrayBaseContext
	AT() antlr.TerminalNode
	ArrayIndexShift() IArrayIndexShiftContext

	// IsArrayIndexContext differentiates from other interfaces.
	IsArrayIndexContext()
}

type ArrayIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayIndexContext() *ArrayIndexContext {
	var p = new(ArrayIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayIndex
	return p
}

func InitEmptyArrayIndexContext(p *ArrayIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayIndex
}

func (*ArrayIndexContext) IsArrayIndexContext() {}

func NewArrayIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayIndexContext {
	var p = new(ArrayIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_arrayIndex

	return p
}

func (s *ArrayIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayIndexContext) ArrayBase() IArrayBaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayBaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayBaseContext)
}

func (s *ArrayIndexContext) AT() antlr.TerminalNode {
	return s.GetToken(StcParserAT, 0)
}

func (s *ArrayIndexContext) ArrayIndexShift() IArrayIndexShiftContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexShiftContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayIndexShiftContext)
}

func (s *ArrayIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArrayIndex(s)
	}
}

func (s *ArrayIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArrayIndex(s)
	}
}

func (p *StcParser) ArrayIndex() (localctx IArrayIndexContext) {
	localctx = NewArrayIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, StcParserRULE_arrayIndex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.ArrayBase()
	}
	{
		p.SetState(193)
		p.Match(StcParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.ArrayIndexShift()
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

// ICapacityContext is an interface to support dynamic dispatch.
type ICapacityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsCapacityContext differentiates from other interfaces.
	IsCapacityContext()
}

type CapacityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapacityContext() *CapacityContext {
	var p = new(CapacityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_capacity
	return p
}

func InitEmptyCapacityContext(p *CapacityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_capacity
}

func (*CapacityContext) IsCapacityContext() {}

func NewCapacityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapacityContext {
	var p = new(CapacityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_capacity

	return p
}

func (s *CapacityContext) GetParser() antlr.Parser { return s.parser }

func (s *CapacityContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserNUMBER, 0)
}

func (s *CapacityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapacityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CapacityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterCapacity(s)
	}
}

func (s *CapacityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitCapacity(s)
	}
}

func (p *StcParser) Capacity() (localctx ICapacityContext) {
	localctx = NewCapacityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, StcParserRULE_capacity)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(StcParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(StcParserNUMBER)
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

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARRAY_OPEN() antlr.TerminalNode
	ARRAY_CLOSE() antlr.TerminalNode
	AllArrayElement() []IArrayElementContext
	ArrayElement(i int) IArrayElementContext
	Capacity() ICapacityContext

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) ARRAY_OPEN() antlr.TerminalNode {
	return s.GetToken(StcParserARRAY_OPEN, 0)
}

func (s *ArrayContext) ARRAY_CLOSE() antlr.TerminalNode {
	return s.GetToken(StcParserARRAY_CLOSE, 0)
}

func (s *ArrayContext) AllArrayElement() []IArrayElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayElementContext); ok {
			len++
		}
	}

	tst := make([]IArrayElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayElementContext); ok {
			tst[i] = t.(IArrayElementContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) ArrayElement(i int) IArrayElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayElementContext); ok {
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

	return t.(IArrayElementContext)
}

func (s *ArrayContext) Capacity() ICapacityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapacityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapacityContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *StcParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, StcParserRULE_array)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(StcParserARRAY_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4681512255488) != 0) {
		{
			p.SetState(200)
			p.ArrayElement()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(205)
		p.Match(StcParserARRAY_CLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__16 {
		{
			p.SetState(206)
			p.Capacity()
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

// IArrayDescriberContext is an interface to support dynamic dispatch.
type IArrayDescriberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARRAY_OPEN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	ARRAY_CLOSE() antlr.TerminalNode

	// IsArrayDescriberContext differentiates from other interfaces.
	IsArrayDescriberContext()
}

type ArrayDescriberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDescriberContext() *ArrayDescriberContext {
	var p = new(ArrayDescriberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayDescriber
	return p
}

func InitEmptyArrayDescriberContext(p *ArrayDescriberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayDescriber
}

func (*ArrayDescriberContext) IsArrayDescriberContext() {}

func NewArrayDescriberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDescriberContext {
	var p = new(ArrayDescriberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_arrayDescriber

	return p
}

func (s *ArrayDescriberContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDescriberContext) ARRAY_OPEN() antlr.TerminalNode {
	return s.GetToken(StcParserARRAY_OPEN, 0)
}

func (s *ArrayDescriberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(StcParserNUMBER, 0)
}

func (s *ArrayDescriberContext) ARRAY_CLOSE() antlr.TerminalNode {
	return s.GetToken(StcParserARRAY_CLOSE, 0)
}

func (s *ArrayDescriberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDescriberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDescriberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArrayDescriber(s)
	}
}

func (s *ArrayDescriberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArrayDescriber(s)
	}
}

func (p *StcParser) ArrayDescriber() (localctx IArrayDescriberContext) {
	localctx = NewArrayDescriberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, StcParserRULE_arrayDescriber)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(StcParserARRAY_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(StcParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(StcParserARRAY_CLOSE)
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

// IArrayNewContext is an interface to support dynamic dispatch.
type IArrayNewContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARRAY_OPERATOR() antlr.TerminalNode
	AllArrayDescriber() []IArrayDescriberContext
	ArrayDescriber(i int) IArrayDescriberContext

	// IsArrayNewContext differentiates from other interfaces.
	IsArrayNewContext()
}

type ArrayNewContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayNewContext() *ArrayNewContext {
	var p = new(ArrayNewContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayNew
	return p
}

func InitEmptyArrayNewContext(p *ArrayNewContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_arrayNew
}

func (*ArrayNewContext) IsArrayNewContext() {}

func NewArrayNewContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayNewContext {
	var p = new(ArrayNewContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_arrayNew

	return p
}

func (s *ArrayNewContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayNewContext) ARRAY_OPERATOR() antlr.TerminalNode {
	return s.GetToken(StcParserARRAY_OPERATOR, 0)
}

func (s *ArrayNewContext) AllArrayDescriber() []IArrayDescriberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayDescriberContext); ok {
			len++
		}
	}

	tst := make([]IArrayDescriberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayDescriberContext); ok {
			tst[i] = t.(IArrayDescriberContext)
			i++
		}
	}

	return tst
}

func (s *ArrayNewContext) ArrayDescriber(i int) IArrayDescriberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDescriberContext); ok {
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

	return t.(IArrayDescriberContext)
}

func (s *ArrayNewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayNewContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayNewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArrayNew(s)
	}
}

func (s *ArrayNewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArrayNew(s)
	}
}

func (p *StcParser) ArrayNew() (localctx IArrayNewContext) {
	localctx = NewArrayNewContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, StcParserRULE_arrayNew)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(StcParserARRAY_OPERATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.ArrayDescriber()
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(215)
			p.ArrayDescriber()
		}

	} else if p.HasError() { // JIM
		goto errorExit
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

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) ID() antlr.TerminalNode {
	return s.GetToken(StcParserID, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *StcParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, StcParserRULE_argument)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(StcParserID)
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

// IVarAssignContext is an interface to support dynamic dispatch.
type IVarAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN_OPERATOR() antlr.TerminalNode
	VarAssignIdentifier() IVarAssignIdentifierContext
	ArrayIndex() IArrayIndexContext
	Identifier() IIdentifierContext

	// IsVarAssignContext differentiates from other interfaces.
	IsVarAssignContext()
}

type VarAssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssignContext() *VarAssignContext {
	var p = new(VarAssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varAssign
	return p
}

func InitEmptyVarAssignContext(p *VarAssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varAssign
}

func (*VarAssignContext) IsVarAssignContext() {}

func NewVarAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssignContext {
	var p = new(VarAssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_varAssign

	return p
}

func (s *VarAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssignContext) ASSIGN_OPERATOR() antlr.TerminalNode {
	return s.GetToken(StcParserASSIGN_OPERATOR, 0)
}

func (s *VarAssignContext) VarAssignIdentifier() IVarAssignIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAssignIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarAssignIdentifierContext)
}

func (s *VarAssignContext) ArrayIndex() IArrayIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayIndexContext)
}

func (s *VarAssignContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *VarAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterVarAssign(s)
	}
}

func (s *VarAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitVarAssign(s)
	}
}

func (p *StcParser) VarAssign() (localctx IVarAssignContext) {
	localctx = NewVarAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, StcParserRULE_varAssign)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(220)
			p.VarAssignIdentifier()
		}

	case 2:
		{
			p.SetState(221)
			p.ArrayIndex()
		}

	case 3:
		{
			p.SetState(222)
			p.Identifier()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(225)
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

// IVarAssignIdentifierContext is an interface to support dynamic dispatch.
type IVarAssignIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarIdentifier() IVarIdentifierContext
	VarType() IVarTypeContext

	// IsVarAssignIdentifierContext differentiates from other interfaces.
	IsVarAssignIdentifierContext()
}

type VarAssignIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssignIdentifierContext() *VarAssignIdentifierContext {
	var p = new(VarAssignIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varAssignIdentifier
	return p
}

func InitEmptyVarAssignIdentifierContext(p *VarAssignIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varAssignIdentifier
}

func (*VarAssignIdentifierContext) IsVarAssignIdentifierContext() {}

func NewVarAssignIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssignIdentifierContext {
	var p = new(VarAssignIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_varAssignIdentifier

	return p
}

func (s *VarAssignIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssignIdentifierContext) VarIdentifier() IVarIdentifierContext {
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

func (s *VarAssignIdentifierContext) VarType() IVarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *VarAssignIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssignIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAssignIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterVarAssignIdentifier(s)
	}
}

func (s *VarAssignIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitVarAssignIdentifier(s)
	}
}

func (p *StcParser) VarAssignIdentifier() (localctx IVarAssignIdentifierContext) {
	localctx = NewVarAssignIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, StcParserRULE_varAssignIdentifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.VarIdentifier()
	}
	{
		p.SetState(228)
		p.VarType()
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

// IVarTypeContext is an interface to support dynamic dispatch.
type IVarTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsVarTypeContext differentiates from other interfaces.
	IsVarTypeContext()
}

type VarTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarTypeContext() *VarTypeContext {
	var p = new(VarTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varType
	return p
}

func InitEmptyVarTypeContext(p *VarTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StcParserRULE_varType
}

func (*VarTypeContext) IsVarTypeContext() {}

func NewVarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarTypeContext {
	var p = new(VarTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StcParserRULE_varType

	return p
}

func (s *VarTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VarTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.EnterVarType(s)
	}
}

func (s *VarTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StcListener); ok {
		listenerT.ExitVarType(s)
	}
}

func (p *StcParser) VarType() (localctx IVarTypeContext) {
	localctx = NewVarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, StcParserRULE_varType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(StcParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Type_()
	}
	{
		p.SetState(232)
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
	p.EnterRule(localctx, 58, StcParserRULE_varReference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(StcParserREFERENCE_OPERATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
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
	p.EnterRule(localctx, 60, StcParserRULE_varIdentifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__5 {
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == StcParserT__5 {
			{
				p.SetState(238)
				p.Match(StcParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(239)
				p.Match(StcParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(242)
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
	p.EnterRule(localctx, 62, StcParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserSTACK_PREVENTION {
		{
			p.SetState(246)
			p.Match(StcParserSTACK_PREVENTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(249)
		p.Match(StcParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == StcParserT__5 {
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == StcParserT__5 {
			{
				p.SetState(250)
				p.Match(StcParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(251)
				p.Match(StcParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(254)
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
