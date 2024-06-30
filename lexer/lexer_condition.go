package lexer

import "rpn/lang"

func CreateIf(tokens []lang.Token, program *lang.Program) []lang.Token {
	newTokens := make([]lang.Token, 0)
	for idx := 0; idx < len(tokens); idx++ {
		token := tokens[idx]
		if token.TokenType == lang.IfT {
			if idx+1 < len(tokens) {
				if tokens[idx+1].TokenType != lang.BlockT {
					panic("invalid if construction")
				}
				tokens[idx+1].Value.(*lang.Block).Tokens = CreateIf(token.Value.(*lang.Block).Tokens, program)
				ifConstruct := lang.IfStatement{TrueBlock: tokens[idx+1].Value.(*lang.Block), FalseBlock: nil}
				newTokens = append(newTokens, lang.Token{TokenType: lang.IfT, Value: ifConstruct, Match: token.Match})
				idx++
			} else {
				panic("invalid if construction")
			}
		} else if token.TokenType == lang.BlockT {
			token.Value.(*lang.Block).Tokens = CreateIf(token.Value.(*lang.Block).Tokens, program)
			newTokens = append(newTokens, token)
		} else {
			newTokens = append(newTokens, token)
		}

	}
	return newTokens
}
