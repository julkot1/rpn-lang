package lexer

import (
	"rpn/lang"
)

func IndexBlocks(function *lang.Function) []*lang.Block {
	newBlocks := make([]*lang.Block, 0)
	index(function.Blocks[0], &newBlocks)
	return newBlocks

}

func index(block *lang.Block, blocks *[]*lang.Block) {
	newTokens := make([]lang.Token, 0)
	*blocks = append(*blocks, block)
	for _, token := range block.Tokens {
		if token.TokenType == lang.BlockT {
			index(token.Value.(*lang.Block), blocks)
		} else {
			newTokens = append(newTokens, token)
		}
	}
	block.Tokens = newTokens

}
