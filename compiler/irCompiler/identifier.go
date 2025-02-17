package irCompiler

import "rpn/parser"

type Identifier struct {
	Base       string
	References []string
}

func GetIdentifier(ctx *parser.IdentifierContext) *Identifier {
	ids := ctx.AllID()
	refs := make([]string, len(ids)-1)
	for idx, id := range ids[1:] {
		refs[idx] = id.GetText()
	}

	id := &Identifier{Base: ids[0].GetText(), References: refs}
	return id
}
