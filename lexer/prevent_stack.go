package lexer

import "rpn/lang"

func LexPrevent(tokens []lang.Token) []lang.Token {
	newTokens := make([]lang.Token, 0)
	idx := 0
	for _, tok := range tokens {
		if tok.TokenType == lang.StackPreventT {
			if ConstTokens[newTokens[idx-1].TokenType].CanPrevent == false {
				panic("Syntax Error: \n Can not stack prevent operation: " + newTokens[idx-1].String())
			}
			newTokens[idx-1].StackPrevent = true
		} else {
			newTokens = append(newTokens, tok)
			idx += 1
		}
	}
	return newTokens
}
