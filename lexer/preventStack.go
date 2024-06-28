package lexer

func LexPrevent(tokens []Token) []Token {
	newTokens := make([]Token, 0)
	idx := 0
	for _, tok := range tokens {
		if tok.TokenType == StackPreventT {
			if newTokens[idx-1].CanPrevent() == false {
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
