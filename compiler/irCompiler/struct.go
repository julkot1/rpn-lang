package irCompiler

import (
	"fmt"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"os"
	"rpn/lang"
	"rpn/parser"
	"rpn/util"
)

func getArgs(id []parser.IStructElementContext, name string, program *lang.Program) []lang.StructElement {
	args := make([]lang.StructElement, len(id))
	for idx, i := range id {
		args[idx] = lang.StructElement{Name: i.ID().GetText(), Type: lang.StringToType(i.VarType().GetText(), program)}
	}
	if !util.IsUnique(args) {
		fmt.Println("Arguments not unique in struct: " + name)
		os.Exit(1)
	}

	return args
}

func createStructDefinition(name string, args []lang.StructElement, program *lang.Program) {
	irType := createStructIr(args)
	x := &lang.StcStruct{Name: name, Args: args, IrType: irType}
	program.StcStruct[name] = x

}

func createStructIr(args []lang.StructElement) *types.StructType {
	irArgs := make([]types.Type, len(args))

	for idx, _ := range args {
		irArgs[idx] = types.I64
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
	alloc := top.Ir.NewAlloca(stcStruct.IrType)

	ptrI64 := top.Ir.NewPtrToInt(alloc, types.I64)
	typ := constant.NewInt(types.I64, int64(lang.REF_T))

	top.Ir.NewCall(program.Funcs[lang.PushT].IrFunc,
		ptrI64,
		typ)

}
