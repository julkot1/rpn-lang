package main

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
	num, err := strconv.Atoi(token)
	if err != nil {
		log.Printf("Error converting string to int: %v\n", err)
	}
	return Token{TokenType: PushT, Value: num}
}
