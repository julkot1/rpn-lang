package test

import (
	"rpn/lang"
	"testing"
)

func TestAbs(t *testing.T) {
	fn := lang.NewFunction("main")
	b1 := lang.NewBlock(fn, nil)
	fn.Blocks = []*lang.Block{b1}
	b2 := lang.NewBlock(fn, b1)

	b1.AddVar(&lang.Var{Name: "a"})
	b1.AddVar(&lang.Var{Name: "b"})
	b1.AddVar(&lang.Var{Name: "c"})

	b2.AddVar(&lang.Var{Name: "d"})
	b2.AddVar(&lang.Var{Name: "e"})
	b2.AddVar(&lang.Var{Name: "f"})

	vars1 := b2.GetVars()
	if len(vars1) != 6 {
		t.Fatalf("expected 7 variables, got %d", len(vars1))
	}
	vars2 := b1.GetVars()
	if len(vars2) != 3 {
		t.Fatalf("expected 3 variables, got %d", len(vars1))
	}

	names := [6]string{"d", "e", "f", "a", "b", "c"}

	for id, v := range vars1 {
		if v.Name != names[id] {
			t.Fatalf("expected %s, got %s", names[id], v.Name)
		}
	}
	for id, v := range vars2 {
		if v.Name != names[id+3] {
			t.Fatalf("expected %s, got %s", names[id+3], v.Name)
		}
	}
}
