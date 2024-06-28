package lexer

import (
	"fmt"
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
	"log"
	"os"
	"strconv"
)

func AddOperation(name string, typ TokenType, lexer *lexmachine.Lexer) {
	lexer.Add([]byte(name), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		return Token{typ, 0, false, match}, nil
	})
}
func CreateLexer() *lexmachine.Lexer {
	lex := lexmachine.NewLexer()

	AddOperation(`pop`, PopT, lex)
	AddOperation(`\+`, AddT, lex)
	AddOperation(`\-`, SubT, lex)
	AddOperation(`\*`, MulT, lex)
	AddOperation(`\/`, DivT, lex)
	AddOperation(`\|\>`, DupT, lex)
	AddOperation(`\|s`, SwapT, lex)
	AddOperation(`\|r`, RotT, lex)
	AddOperation(`\|o`, OverT, lex)
	AddOperation(`puts`, PrintI8T, lex)
	AddOperation(`print`, PrintT, lex)
	AddOperation(`input`, InputT, lex)
	AddOperation(`!`, StackPreventT, lex)

	lex.Add([]byte(`[0-9]+`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		num, err := strconv.Atoi(string(match.Bytes))
		return Token{PushT, num, false, match}, err
	})
	lex.Add([]byte(`\!`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		fmt.Println("foo")
		return nil, nil
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

func Parse(file string) []Token {
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
	var tokens []Token
	for tok, err, eos := scan.Next(); !eos; tok, err, eos = scan.Next() {
		if err != nil {
			log.Fatalf("Scan error: %v", err)
		}
		tokens = append(tokens, tok.(Token))
	}

	tokens = LexPrevent(tokens)

	return tokens
}
