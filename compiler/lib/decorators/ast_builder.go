package decorators

import (
	"errors"
	"strconv"
	"strings"
	_ "strings"
)

type PreprocessorType int32
type MacroId int32

const (
	PreprocessorMacroT PreprocessorType = iota
	PreprocessorIntT
	PreprocessorString
	PreprocessorVarT
)

var MarcosName = [...]string{
	"STC", "NAME", "CODE", "CAN_PREVENT", "TYPE", "FUNCTION"}

const (
	MacroStc MacroId = iota
	MacroName
	MacroCode
	MacroCanPrevent
)

type PreprocessorToken struct {
	Text  string
	Typ   PreprocessorType
	Value interface{}
	Args  []PreprocessorToken
}

func ParseStcHeader(input string) (*STCBindConfig, error) {

	tokens := getTokens(input)
	ast, err := getAst(tokens)
	if err != nil {
		return nil, err
	}
	analyzeAst(&ast)
	return CreateSTCBindConfig(ast)
}

func analyzeAst(ast *PreprocessorToken) *PreprocessorToken {
	if len(ast.Args) == 0 {
		i, err := strconv.Atoi(ast.Text)
		if err == nil {
			ast.Value = i
			ast.Typ = PreprocessorIntT
			return ast
		}
		if strings.HasPrefix(ast.Text, "\"") && strings.HasSuffix(ast.Text, "\"") {
			ast.Typ = PreprocessorString
			ast.Value = ast.Text[1 : len(ast.Text)-1]
			return ast
		}
		if isMacro(ast.Text) {
			ast.Typ = PreprocessorMacroT
		} else {
			ast.Typ = PreprocessorVarT
		}
		ast.Value = ast.Text
		return ast
	}
	ast.Typ = PreprocessorMacroT
	ast.Value = ast.Text
	for index, arg := range ast.Args {
		ast.Args[index] = *analyzeAst(&arg)
	}

	return ast
}

func isMacro(text string) bool {
	for _, macro := range MarcosName {
		if text == macro {
			return true
		}
	}
	return false
}

func getAst(tokens []string) (PreprocessorToken, error) {
	stack := NewStack[PreprocessorToken]()
	for _, token := range tokens {
		if token == ")" {
			args := make([]PreprocessorToken, 0)
			var tok PreprocessorToken
			for {
				tok, _ = stack.Pop()
				if tok.Text == "(" {
					break
				}
				args = append(args, tok)
			}
			tok, _ = stack.Pop()
			tok.Args = reverse(args)
			stack.Push(tok)
			continue
		}
		if token != "," {
			stack.Push(PreprocessorToken{Text: token})
		}

	}
	root, err := stack.Pop()

	return root, err
}

func reverse[T any](arr []T) []T {
	left, right := 0, len(arr)-1
	for left < right {
		// Swap elements
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}
func getTokens(input string) []string {
	str := removeWhiteSpaces(input)
	word := ""
	var tokens []string
	isString := false
	for _, char := range str {
		if char == '"' {
			if isString {
				tokens = append(tokens, word+"\"")
				word = ""
				isString = false
				continue
			}
			isString = true
		}
		if isString {
			word = word + string(char)
			continue
		}
		if isWord(char) {
			if word != "" {
				tokens = append(tokens, word)
			}
			tokens = append(tokens, string(char))
			word = ""
		} else {
			word += string(char)
		}
	}
	return tokens
}

func isWord(word int32) bool {
	if word == ',' || word == '(' || word == ')' {
		return true
	}
	return false
}

func removeWhiteSpaces(str string) string {
	strNoWhitespace := strings.ReplaceAll(str, " ", "")
	strNoWhitespace = strings.ReplaceAll(strNoWhitespace, "\t", "")
	strNoWhitespace = strings.ReplaceAll(strNoWhitespace, "\n", "")
	return strNoWhitespace
}

type Stack[T any] struct {
	elements []T
}

// NewStack initializes and returns a new stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element of the stack.
// It returns an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T // zero value of T
		return zero, errors.New("stack is empty")
	}
	// Get the top element
	top := s.elements[len(s.elements)-1]
	// Remove the top element
	s.elements = s.elements[:len(s.elements)-1]
	return top, nil
}

// Peek returns the top element without removing it.
// It returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T // zero value of T
		return zero, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.elements)
}
