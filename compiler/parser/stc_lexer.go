// Code generated from Stc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type StcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var StcLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stclexerLexerInit() {
	staticData := &StcLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'fun'", "'{'", "'}'", "'if'", "'else'", "'repeat'", "'dup'", "'rot'",
		"'swap'", "'pop'", "':'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "STACK_PREVENTION",
		"NUMBER", "SIGNED_NUMBER", "FLOAT", "SIGNED_FLOAT", "FLOAT_E", "SIGNED_FLOAT_E",
		"BIN_NUMBER", "HEX_NUMBER", "STRING", "CHAR", "BOOL", "SIMPLE_TYPE",
		"LOGIC_OPERATOR", "ARITHMETIC_OPERATOR", "BUILD_IN_OPERATOR", "ID",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "STACK_PREVENTION", "NUMBER", "SIGNED_NUMBER", "FLOAT",
		"SIGNED_FLOAT", "FLOAT_E", "SIGNED_FLOAT_E", "DIGIT", "BIN_NUMBER",
		"HEX_NUMBER", "STRING", "CHAR", "BOOL", "ESC", "SIMPLE_TYPE", "LOGIC_OPERATOR",
		"ARITHMETIC_OPERATOR", "BUILD_IN_OPERATOR", "ID", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 297, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 4, 12, 109, 8,
		12, 11, 12, 12, 12, 110, 1, 13, 1, 13, 4, 13, 115, 8, 13, 11, 13, 12, 13,
		116, 1, 14, 4, 14, 120, 8, 14, 11, 14, 12, 14, 121, 1, 14, 1, 14, 4, 14,
		126, 8, 14, 11, 14, 12, 14, 127, 1, 15, 1, 15, 4, 15, 132, 8, 15, 11, 15,
		12, 15, 133, 1, 15, 1, 15, 4, 15, 138, 8, 15, 11, 15, 12, 15, 139, 1, 16,
		4, 16, 143, 8, 16, 11, 16, 12, 16, 144, 1, 16, 1, 16, 4, 16, 149, 8, 16,
		11, 16, 12, 16, 150, 3, 16, 153, 8, 16, 1, 16, 1, 16, 3, 16, 157, 8, 16,
		1, 16, 4, 16, 160, 8, 16, 11, 16, 12, 16, 161, 1, 17, 1, 17, 4, 17, 166,
		8, 17, 11, 17, 12, 17, 167, 1, 17, 1, 17, 4, 17, 172, 8, 17, 11, 17, 12,
		17, 173, 3, 17, 176, 8, 17, 1, 17, 1, 17, 3, 17, 180, 8, 17, 1, 17, 4,
		17, 183, 8, 17, 11, 17, 12, 17, 184, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		4, 19, 192, 8, 19, 11, 19, 12, 19, 193, 1, 20, 1, 20, 1, 20, 4, 20, 199,
		8, 20, 11, 20, 12, 20, 200, 1, 21, 1, 21, 1, 21, 5, 21, 206, 8, 21, 10,
		21, 12, 21, 209, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 3, 22, 216,
		8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 3, 23, 229, 8, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 3, 25, 251, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 3, 26, 270, 8, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 282, 8, 28, 1, 29, 1, 29, 5,
		29, 286, 8, 29, 10, 29, 12, 29, 289, 9, 29, 1, 30, 4, 30, 292, 8, 30, 11,
		30, 12, 30, 293, 1, 30, 1, 30, 0, 0, 31, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 0, 39, 19, 41, 20, 43, 21, 45, 22, 47,
		23, 49, 0, 51, 24, 53, 25, 55, 26, 57, 27, 59, 28, 61, 29, 1, 0, 15, 2,
		0, 43, 43, 45, 45, 2, 0, 69, 69, 101, 101, 1, 0, 48, 57, 2, 0, 66, 66,
		98, 98, 1, 0, 48, 49, 2, 0, 88, 88, 120, 120, 3, 0, 48, 57, 65, 70, 97,
		102, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 8, 0, 34, 34, 39, 39,
		92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 2, 0, 60, 60, 62,
		62, 4, 0, 37, 37, 42, 43, 45, 45, 47, 47, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 330,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0,
		0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 63, 1, 0, 0, 0, 3, 67, 1, 0, 0,
		0, 5, 69, 1, 0, 0, 0, 7, 71, 1, 0, 0, 0, 9, 74, 1, 0, 0, 0, 11, 79, 1,
		0, 0, 0, 13, 86, 1, 0, 0, 0, 15, 90, 1, 0, 0, 0, 17, 94, 1, 0, 0, 0, 19,
		99, 1, 0, 0, 0, 21, 103, 1, 0, 0, 0, 23, 105, 1, 0, 0, 0, 25, 108, 1, 0,
		0, 0, 27, 112, 1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 129, 1, 0, 0, 0, 33,
		142, 1, 0, 0, 0, 35, 163, 1, 0, 0, 0, 37, 186, 1, 0, 0, 0, 39, 188, 1,
		0, 0, 0, 41, 195, 1, 0, 0, 0, 43, 202, 1, 0, 0, 0, 45, 212, 1, 0, 0, 0,
		47, 228, 1, 0, 0, 0, 49, 230, 1, 0, 0, 0, 51, 250, 1, 0, 0, 0, 53, 269,
		1, 0, 0, 0, 55, 271, 1, 0, 0, 0, 57, 281, 1, 0, 0, 0, 59, 283, 1, 0, 0,
		0, 61, 291, 1, 0, 0, 0, 63, 64, 5, 102, 0, 0, 64, 65, 5, 117, 0, 0, 65,
		66, 5, 110, 0, 0, 66, 2, 1, 0, 0, 0, 67, 68, 5, 123, 0, 0, 68, 4, 1, 0,
		0, 0, 69, 70, 5, 125, 0, 0, 70, 6, 1, 0, 0, 0, 71, 72, 5, 105, 0, 0, 72,
		73, 5, 102, 0, 0, 73, 8, 1, 0, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76, 5, 108,
		0, 0, 76, 77, 5, 115, 0, 0, 77, 78, 5, 101, 0, 0, 78, 10, 1, 0, 0, 0, 79,
		80, 5, 114, 0, 0, 80, 81, 5, 101, 0, 0, 81, 82, 5, 112, 0, 0, 82, 83, 5,
		101, 0, 0, 83, 84, 5, 97, 0, 0, 84, 85, 5, 116, 0, 0, 85, 12, 1, 0, 0,
		0, 86, 87, 5, 100, 0, 0, 87, 88, 5, 117, 0, 0, 88, 89, 5, 112, 0, 0, 89,
		14, 1, 0, 0, 0, 90, 91, 5, 114, 0, 0, 91, 92, 5, 111, 0, 0, 92, 93, 5,
		116, 0, 0, 93, 16, 1, 0, 0, 0, 94, 95, 5, 115, 0, 0, 95, 96, 5, 119, 0,
		0, 96, 97, 5, 97, 0, 0, 97, 98, 5, 112, 0, 0, 98, 18, 1, 0, 0, 0, 99, 100,
		5, 112, 0, 0, 100, 101, 5, 111, 0, 0, 101, 102, 5, 112, 0, 0, 102, 20,
		1, 0, 0, 0, 103, 104, 5, 58, 0, 0, 104, 22, 1, 0, 0, 0, 105, 106, 5, 33,
		0, 0, 106, 24, 1, 0, 0, 0, 107, 109, 3, 37, 18, 0, 108, 107, 1, 0, 0, 0,
		109, 110, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111,
		26, 1, 0, 0, 0, 112, 114, 7, 0, 0, 0, 113, 115, 3, 37, 18, 0, 114, 113,
		1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0,
		0, 0, 117, 28, 1, 0, 0, 0, 118, 120, 3, 37, 18, 0, 119, 118, 1, 0, 0, 0,
		120, 121, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 125, 5, 46, 0, 0, 124, 126, 3, 37, 18, 0, 125, 124,
		1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0,
		0, 0, 128, 30, 1, 0, 0, 0, 129, 131, 7, 0, 0, 0, 130, 132, 3, 37, 18, 0,
		131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133,
		134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 5, 46, 0, 0, 136, 138,
		3, 37, 18, 0, 137, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1,
		0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 32, 1, 0, 0, 0, 141, 143, 3, 37, 18,
		0, 142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144,
		145, 1, 0, 0, 0, 145, 152, 1, 0, 0, 0, 146, 148, 5, 46, 0, 0, 147, 149,
		3, 37, 18, 0, 148, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 148, 1,
		0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 146, 1, 0, 0,
		0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 156, 7, 1, 0, 0, 155,
		157, 7, 0, 0, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 159,
		1, 0, 0, 0, 158, 160, 3, 37, 18, 0, 159, 158, 1, 0, 0, 0, 160, 161, 1,
		0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 34, 1, 0, 0,
		0, 163, 165, 7, 0, 0, 0, 164, 166, 3, 37, 18, 0, 165, 164, 1, 0, 0, 0,
		166, 167, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168,
		175, 1, 0, 0, 0, 169, 171, 5, 46, 0, 0, 170, 172, 3, 37, 18, 0, 171, 170,
		1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0,
		0, 0, 174, 176, 1, 0, 0, 0, 175, 169, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0,
		176, 177, 1, 0, 0, 0, 177, 179, 7, 1, 0, 0, 178, 180, 7, 0, 0, 0, 179,
		178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181, 183,
		3, 37, 18, 0, 182, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 182, 1,
		0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 36, 1, 0, 0, 0, 186, 187, 7, 2, 0,
		0, 187, 38, 1, 0, 0, 0, 188, 189, 5, 48, 0, 0, 189, 191, 7, 3, 0, 0, 190,
		192, 7, 4, 0, 0, 191, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 191,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 40, 1, 0, 0, 0, 195, 196, 5, 48,
		0, 0, 196, 198, 7, 5, 0, 0, 197, 199, 7, 6, 0, 0, 198, 197, 1, 0, 0, 0,
		199, 200, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		42, 1, 0, 0, 0, 202, 207, 5, 34, 0, 0, 203, 206, 3, 49, 24, 0, 204, 206,
		8, 7, 0, 0, 205, 203, 1, 0, 0, 0, 205, 204, 1, 0, 0, 0, 206, 209, 1, 0,
		0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 210, 1, 0, 0, 0,
		209, 207, 1, 0, 0, 0, 210, 211, 5, 34, 0, 0, 211, 44, 1, 0, 0, 0, 212,
		215, 5, 39, 0, 0, 213, 216, 3, 49, 24, 0, 214, 216, 8, 8, 0, 0, 215, 213,
		1, 0, 0, 0, 215, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 5, 39,
		0, 0, 218, 46, 1, 0, 0, 0, 219, 220, 5, 116, 0, 0, 220, 221, 5, 114, 0,
		0, 221, 222, 5, 117, 0, 0, 222, 229, 5, 101, 0, 0, 223, 224, 5, 102, 0,
		0, 224, 225, 5, 97, 0, 0, 225, 226, 5, 108, 0, 0, 226, 227, 5, 115, 0,
		0, 227, 229, 5, 101, 0, 0, 228, 219, 1, 0, 0, 0, 228, 223, 1, 0, 0, 0,
		229, 48, 1, 0, 0, 0, 230, 231, 5, 92, 0, 0, 231, 232, 7, 9, 0, 0, 232,
		50, 1, 0, 0, 0, 233, 234, 5, 73, 0, 0, 234, 235, 5, 54, 0, 0, 235, 251,
		5, 52, 0, 0, 236, 237, 5, 73, 0, 0, 237, 251, 5, 56, 0, 0, 238, 239, 5,
		83, 0, 0, 239, 240, 5, 116, 0, 0, 240, 251, 5, 114, 0, 0, 241, 242, 5,
		70, 0, 0, 242, 243, 5, 108, 0, 0, 243, 244, 5, 111, 0, 0, 244, 245, 5,
		97, 0, 0, 245, 251, 5, 116, 0, 0, 246, 247, 5, 66, 0, 0, 247, 248, 5, 111,
		0, 0, 248, 249, 5, 111, 0, 0, 249, 251, 5, 108, 0, 0, 250, 233, 1, 0, 0,
		0, 250, 236, 1, 0, 0, 0, 250, 238, 1, 0, 0, 0, 250, 241, 1, 0, 0, 0, 250,
		246, 1, 0, 0, 0, 251, 52, 1, 0, 0, 0, 252, 253, 5, 60, 0, 0, 253, 270,
		5, 61, 0, 0, 254, 255, 5, 62, 0, 0, 255, 270, 5, 61, 0, 0, 256, 257, 5,
		33, 0, 0, 257, 270, 5, 61, 0, 0, 258, 259, 5, 61, 0, 0, 259, 270, 5, 61,
		0, 0, 260, 270, 7, 10, 0, 0, 261, 262, 5, 97, 0, 0, 262, 263, 5, 110, 0,
		0, 263, 270, 5, 100, 0, 0, 264, 265, 5, 111, 0, 0, 265, 270, 5, 114, 0,
		0, 266, 267, 5, 110, 0, 0, 267, 268, 5, 111, 0, 0, 268, 270, 5, 116, 0,
		0, 269, 252, 1, 0, 0, 0, 269, 254, 1, 0, 0, 0, 269, 256, 1, 0, 0, 0, 269,
		258, 1, 0, 0, 0, 269, 260, 1, 0, 0, 0, 269, 261, 1, 0, 0, 0, 269, 264,
		1, 0, 0, 0, 269, 266, 1, 0, 0, 0, 270, 54, 1, 0, 0, 0, 271, 272, 7, 11,
		0, 0, 272, 56, 1, 0, 0, 0, 273, 274, 5, 116, 0, 0, 274, 275, 5, 121, 0,
		0, 275, 276, 5, 112, 0, 0, 276, 277, 5, 101, 0, 0, 277, 278, 5, 111, 0,
		0, 278, 282, 5, 102, 0, 0, 279, 280, 5, 58, 0, 0, 280, 282, 5, 61, 0, 0,
		281, 273, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 58, 1, 0, 0, 0, 283, 287,
		7, 12, 0, 0, 284, 286, 7, 13, 0, 0, 285, 284, 1, 0, 0, 0, 286, 289, 1,
		0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 60, 1, 0, 0,
		0, 289, 287, 1, 0, 0, 0, 290, 292, 7, 14, 0, 0, 291, 290, 1, 0, 0, 0, 292,
		293, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 295,
		1, 0, 0, 0, 295, 296, 6, 30, 0, 0, 296, 62, 1, 0, 0, 0, 28, 0, 110, 116,
		121, 127, 133, 139, 144, 150, 152, 156, 161, 167, 173, 175, 179, 184, 193,
		200, 205, 207, 215, 228, 250, 269, 281, 287, 293, 1, 6, 0, 0,
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

// StcLexerInit initializes any static state used to implement StcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewStcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func StcLexerInit() {
	staticData := &StcLexerLexerStaticData
	staticData.once.Do(stclexerLexerInit)
}

// NewStcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewStcLexer(input antlr.CharStream) *StcLexer {
	StcLexerInit()
	l := new(StcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &StcLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Stc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// StcLexer tokens.
const (
	StcLexerT__0                = 1
	StcLexerT__1                = 2
	StcLexerT__2                = 3
	StcLexerT__3                = 4
	StcLexerT__4                = 5
	StcLexerT__5                = 6
	StcLexerT__6                = 7
	StcLexerT__7                = 8
	StcLexerT__8                = 9
	StcLexerT__9                = 10
	StcLexerT__10               = 11
	StcLexerSTACK_PREVENTION    = 12
	StcLexerNUMBER              = 13
	StcLexerSIGNED_NUMBER       = 14
	StcLexerFLOAT               = 15
	StcLexerSIGNED_FLOAT        = 16
	StcLexerFLOAT_E             = 17
	StcLexerSIGNED_FLOAT_E      = 18
	StcLexerBIN_NUMBER          = 19
	StcLexerHEX_NUMBER          = 20
	StcLexerSTRING              = 21
	StcLexerCHAR                = 22
	StcLexerBOOL                = 23
	StcLexerSIMPLE_TYPE         = 24
	StcLexerLOGIC_OPERATOR      = 25
	StcLexerARITHMETIC_OPERATOR = 26
	StcLexerBUILD_IN_OPERATOR   = 27
	StcLexerID                  = 28
	StcLexerWS                  = 29
)