package lexer

import (
	"rpn/lang"
	"rpn/util"
)

func CreateBlocks(tokens []lang.Token, program *lang.Program) []lang.Token {
	newTokens := make([]lang.Token, 0)
	blocksStack := util.NewStack()
	closeCount := 0
	for idx := 0; idx < len(tokens); idx++ {
		tok := tokens[idx]
		if tok.TokenType == lang.BlockOpenT {

			newBlock := lang.NewBlock(program.NewBlockIndex())
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
			if idx+1 < len(tokens) {
				nextToken := tokens[idx+1]
				if nextToken.TokenType != lang.BlockOpenT && nextToken.TokenType != lang.FunctionDefT {
					closeCount++
					if closeCount%2 == 1 {
						newBlock := lang.NewBlock(program.NewBlockIndex())
						newBlock.Tokens = make([]lang.Token, 0)
						blocksStack.Push(newBlock)
					} else {
						currentBlock, _ = blocksStack.Pop()
						parentBlock, _ = blocksStack.Top()
						if parentBlock == nil {
							newTokens = append(newTokens, lang.Token{TokenType: lang.BlockT, Value: currentBlock})
						} else {
							x := parentBlock.(*lang.Block)
							x.Tokens = append(x.Tokens, lang.Token{TokenType: lang.BlockT, Value: currentBlock})
						}
					}
				}
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
	if !blocksStack.IsEmpty() {
		currentBlock, _ := blocksStack.Pop()
		newTokens = append(newTokens, lang.Token{TokenType: lang.BlockT, Value: currentBlock})
	}
	return newTokens
}

func CreateGlobalFunctions(tokens []lang.Token) []lang.Token {
	newTokens := make([]lang.Token, 0)

	var name lang.Token
	var block lang.Token
	for idx := 0; idx < len(tokens); idx++ {
		tok := tokens[idx]
		if tok.TokenType == lang.FunctionDefT {
			name = tokens[idx+1]
			block = tokens[idx+2]
			idx += 2
			newTokens = append(newTokens, lang.Token{TokenType: lang.FunctionT, Value: CreateFunction(name, block)})
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
	AssignFuncToBlock(fun, block.Value.(*lang.Block))
	fun.Blocks = append(fun.Blocks, block.Value.(*lang.Block))
	return fun
}

func AssignFuncToBlock(fun *lang.Function, block *lang.Block) {
	block.Func = fun
	for _, tok := range block.Tokens {
		if tok.TokenType == lang.BlockT {
			AssignFuncToBlock(fun, tok.Value.(*lang.Block))
		}
	}

}

func LoadFunctions(program *lang.Program, tokens []lang.Token) {
	program.Functions = make([]*lang.Function, 0)
	for _, token := range tokens {
		switch token.TokenType {
		case lang.FunctionT:
			program.Functions = append(program.Functions, token.Value.(*lang.Function))
			if token.Value.(*lang.Function).Name == "main" {
				program.MainFunction = token.Value.(*lang.Function)
			}
			break
		default:
			panic("invalid Token")
		}
	}
}
