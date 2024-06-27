package lexer

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type TokenType int

const (
	PushT TokenType = iota
	PopT
	AddT
	SubT
	MulT
	PrintT
	PrintI8T
	InputT
)

type Token struct {
	TokenType TokenType
	Value     int
}

func Parse(file string) []Token {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	stringContent := string(content)
	stringTokens := strings.Split(stringContent, " ")
	result := make([]Token, len(stringTokens))
	for idx, token := range stringTokens {
		result[idx] = NewToken(token)
	}
	return result
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
	num, err := strconv.Atoi(token)
	if err != nil {
		log.Printf("Error converting string to int: %v\n", err)
	}
	return Token{TokenType: PushT, Value: num}
}
