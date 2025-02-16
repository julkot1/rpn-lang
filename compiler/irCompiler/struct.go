package irCompiler

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir/types"
	"os"
	"rpn/lang"
	"rpn/util"
)

func getArgs(id []antlr.TerminalNode, name string) []lang.StructElement {
	args := make([]lang.StructElement, len(id))
	for idx, i := range id {
		args[idx] = lang.StructElement{Name: i.GetText(), Type: lang.ANY_T}
	}
	if !util.IsUnique(args) {
		fmt.Println("Arguments not unique in struct: " + name)
		os.Exit(1)
	}

	return args
}

func createStructDefinition(name string, args []lang.StructElement, program *lang.Program) {
	irType := createStructIr(args, program)
	x := &lang.StcStruct{Name: name, Args: args, IrType: irType}
	program.StcStruct[name] = x

}

func createStructIr(args []lang.StructElement, program *lang.Program) *types.StructType {
	irArgs := make([]types.Type, len(args))
	varType := program.Structs["variable"]
	for idx, _ := range args {
		irArgs[idx] = varType
	}
	return types.NewStruct(irArgs...)
}

func createStruct(name string, top *lang.Block, program *lang.Program) {
	x, ok := program.GlobalTokenTable[name]
	if !ok {
		fmt.Println("No such struct: " + name)
		os.Exit(1)
	}
	if x != lang.PStruct {
		fmt.Println("Token " + name + " is not a struct")
	}
	stcStruct := program.StcStruct[name]
	top.Ir.NewAlloca(stcStruct.IrType)

}
