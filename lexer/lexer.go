package lexer

import (
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
	"log"
	"os"
	"strconv"
)

type TokenType int

const (
	PushT TokenType = iota
	PopT
	AddT
	SubT
	MulT
	PrintT
	DivT
	PrintI8T
	InputT
	DupT
	SwapT
	OverT
	RotT
)

type Token struct {
	TokenType TokenType
	Value     interface{}
	Match     *machines.Match
}

func AddOperation(name string, typ TokenType, lexer *lexmachine.Lexer) {
	lexer.Add([]byte(name), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		return Token{typ, 0, match}, nil
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

	lex.Add([]byte(`[0-9]+`), func(scan *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		num, err := strconv.Atoi(string(match.Bytes))
		return Token{PushT, num, match}, err
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

	return tokens
}

func NewToken(token string) Token {
	if token == "+" {
		return Token{TokenType: AddT}
	}
	if token == "-" {
		return Token{TokenType: SubT}
	}
	if token == "*" {
		return Token{TokenType: MulT}
	}
	if token == "print" {
		return Token{TokenType: PrintT}
	}
	if token == "puts" {
		return Token{TokenType: PrintI8T}
	}
	if token == "input" {
		return Token{TokenType: InputT}
	}
	if token == "|>" {
		return Token{TokenType: DupT}
	}
	if token == "|s" {
		return Token{TokenType: SwapT}
	}
	if token == "|r" {
		return Token{TokenType: RotT}
	}
	if token == "|o" {
		return Token{TokenType: OverT}
	}
	num, err := strconv.Atoi(token)
	if err != nil {
		log.Printf("Error converting string to int: %v\n", err)
	}
	return Token{TokenType: PushT, Value: num}
}
