package lang

import (
	"errors"
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"strconv"
)

type StringTable struct {
	index int
	table map[string]*ir.Global
}

func NewStringTable() *StringTable {
	return &StringTable{index: 0, table: make(map[string]*ir.Global)}
}
func (s *StringTable) nextID() string {
	s.index++
	return ".stc.string" + strconv.Itoa(s.index-1)
}
func (s *StringTable) Add(str string, program *Program) *ir.Global {
	strConst := constant.NewCharArrayFromString(str + "\x00")

	g := program.Module.NewGlobalDef(s.nextID(), strConst)
	g.Immutable = true
	g.Linkage = enum.LinkagePrivate
	s.table[str] = g
	return s.table[str]
}
func (s *StringTable) Get(str string) (*ir.Global, error) {
	x, ok := s.table[str]
	if !ok {
		return nil, errors.New(fmt.Sprintf("key %v not found in map", str))
	}
	return x, nil
}
func (s *StringTable) GetOrAdd(str string, program *Program) *ir.Global {
	g, err := s.Get(str)
	if err != nil {
		return s.Add(str, program)
	}
	return g
}
