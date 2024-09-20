package lexer

import (
	"rpn/lang"
)

func CreateRepeat(tokens []lang.Token, program *lang.Program) []lang.Token {
	newTokens := make([]lang.Token, 0)
	for _, token := range tokens {
		if token.TokenType == lang.RepeatT {
			token.Value.(*lang.RepeatStatement).LoopBlock = lang.NewBlock(program.NewBlockIndex(), true)
		}
		newTokens = append(newTokens, token)
	}
	return newTokens
}

func AddRepeatBlock(block *lang.Block) {
	var loopBlocks []*lang.Block
	for _, token := range block.Tokens {
		if token.TokenType == lang.RepeatT {
			token.Value.(*lang.RepeatStatement).LoopBlock.Parent = block
			loopBlocks = append([]*lang.Block{token.Value.(*lang.RepeatStatement).LoopBlock}, loopBlocks...)

		}
	}
	block.Children = append(loopBlocks, block.Children...)
}
