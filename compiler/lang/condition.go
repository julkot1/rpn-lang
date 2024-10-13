package lang

type IfStatement struct {
	TrueBlock  *Block
	FalseBlock *Block
}
type RepeatStatement struct {
	LoopBlock *Block
}
