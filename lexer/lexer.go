package lexer

import (
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
	"log"
	"os"
	"rpn/lang"
	"strconv"
)

func AddOperation(name string, typ lang.TokenType, lexer *lexmachine.Lexer) {
	lexer.Add([]byte(name), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		return lang.Token{typ, 0, false, match}, nil
	})
}
func CreateLexer() *lexmachine.Lexer {
	lex := lexmachine.NewLexer()

	AddOperation(PopToken, lang.PopT, lex)
	AddOperation(AddToken, lang.AddT, lex)
	AddOperation(SubToken, lang.SubT, lex)
	AddOperation(MulToken, lang.MulT, lex)
	AddOperation(DivToken, lang.DivT, lex)
	AddOperation(DupToken, lang.DupT, lex)
	AddOperation(SwapToken, lang.SwapT, lex)
	AddOperation(RotToken, lang.RotT, lex)
	AddOperation(OverToken, lang.OverT, lex)
	AddOperation(PrintI8Token, lang.PrintI8T, lex)
	AddOperation(PrintToken, lang.PrintT, lex)
	AddOperation(InputToken, lang.InputT, lex)
	AddOperation(StackPreventToken, lang.StackPreventT, lex)
	AddOperation(FunctionDefToken, lang.FunctionDefT, lex)
	AddOperation(BlockOpenToken, lang.BlockOpenT, lex)
	AddOperation(BlockCloseToken, lang.BlockCloseT, lex)

	lex.Add([]byte(IdentifierToken), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		return lang.Token{TokenType: lang.IdentifierT, Value: string(match.Bytes), Match: match}, nil
	})
	lex.Add([]byte(`[0-9]+`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		num, err := strconv.Atoi(string(match.Bytes))
		return lang.Token{TokenType: lang.PushT, Value: num, Match: match}, err
	})

	lex.Add([]byte(`[ \t\r\n]+`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		return nil, nil
	})
	lex.Add([]byte(`//[^\n]*`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		return nil, nil
	})

	lexerErr := lex.Compile()
	if lexerErr != nil {
		log.Fatalf("Lexer compile error: %v", lexerErr)
	}

	return lex
}

func Parse(file string) *lang.Program {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	lex := CreateLexer()
	stringContent := string(content)

	scan, scanErr := lex.Scanner([]byte(stringContent))
	if scanErr != nil {
		log.Fatalf("Scanner error: %v", scanErr)
	}
	var tokens []lang.Token
	for tok, err, eos := scan.Next(); !eos; tok, err, eos = scan.Next() {
		if err != nil {
			log.Fatalf("Scan error: %v", err)
		}
		tokens = append(tokens, tok.(lang.Token))
	}

	program := lang.NewProgram()

	tokens = LexPrevent(tokens)
	tokens = CreateBlocks(tokens, program)
	tokens = CreateGlobalFunctions(tokens)
	LoadFunctions(program, tokens)

	return program
}
