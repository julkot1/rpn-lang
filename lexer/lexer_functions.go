package lexer

import (
	"rpn/lang"
	"rpn/util"
)

func CreateBlocks(tokens []lang.Token) []lang.Token {
	newTokens := make([]lang.Token, 0)
	blocksStack := util.NewStack()
	for idx := 0; idx < len(tokens); idx++ {
		tok := tokens[idx]
		if tok.TokenType == lang.BlockOpenT {
			newBlock := lang.NewBlock(nil, nil)
			newBlock.Tokens = make([]lang.Token, 0)
			blocksStack.Push(newBlock)
		} else if tok.TokenType == lang.BlockCloseT {
			currentBlock, _ := blocksStack.Pop()
			parentBlock, _ := blocksStack.Top()
			if parentBlock == nil {
				newTokens = append(newTokens, lang.Token{TokenType: lang.BlockT, Value: currentBlock})
			} else {
				x := parentBlock.(*lang.Block)
				x.Tokens = append(x.Tokens, lang.Token{TokenType: lang.BlockT, Value: currentBlock})
			}

		} else {
			block, _ := blocksStack.Top()
			if block == nil {
				newTokens = append(newTokens, tok)
			} else {
				x := block.(*lang.Block)
				x.Tokens = append(x.Tokens, tok)
			}
		}

	}
	return newTokens
}

func CreateGlobalFunctions(tokens []lang.Token) []lang.Token {
	newTokens := make([]lang.Token, 0)
	newFuncDef := false
	var name lang.Token
	var block lang.Token
	for idx := 0; idx < len(tokens); idx++ {
		tok := tokens[idx]
		if tok.TokenType == lang.FunctionDefT {
			if newFuncDef {
				panic("invalid function definition")
			}
			newFuncDef = true
			name = tokens[idx+1]
			block = tokens[idx+2]
			idx += 2
			newTokens = append(newTokens, lang.Token{TokenType: lang.FunctionT, Value: CreateFunction(name, block)})
			continue
		}

	}
	return newTokens
}

func CreateFunction(name lang.Token, block lang.Token) *lang.Function {
	if name.TokenType != lang.IdentifierT || block.TokenType != lang.BlockT {
		panic("invalid function definition")
	}
	fun := lang.NewFunction(name.Value.(string))
	fun.Blocks = make([]*lang.Block, 0)
	fun.Blocks = append(fun.Blocks, block.Value.(*lang.Block))
	return fun
}
