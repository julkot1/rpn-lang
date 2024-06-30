package irCompiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"rpn/lang"
)

func DefineBinLogicFunction(program *lang.Program, name string, cmp enum.IPred) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, name)
	result := binFnBody.NewICmp(cmp, a, b)
	zexted := binFnBody.NewZExt(result, types.I32)
	binFnBody.NewCall(program.Funcs["push"].IrFunc, zexted)
	binFnBody.NewRet(nil)
	return binFn
}

func DefineEqualsFunction(program *lang.Program) *ir.Func {
	return DefineBinLogicFunction(program, "equals", enum.IPredEQ)

}
func DefineGreaterFunction(program *lang.Program) *ir.Func {
	return DefineBinLogicFunction(program, "grater", enum.IPredSGT)

}
func DefineLessFunction(program *lang.Program) *ir.Func {
	return DefineBinLogicFunction(program, "less", enum.IPredSLT)

}
func DefineNotEqualsFunction(program *lang.Program) *ir.Func {
	return DefineBinLogicFunction(program, "not_equals", enum.IPredNE)

}
func DefineGreaterOrEqFunction(program *lang.Program) *ir.Func {
	return DefineBinLogicFunction(program, "greater_eq", enum.IPredSGE)
}

func DefineLessOrEqFunction(program *lang.Program) *ir.Func {
	return DefineBinLogicFunction(program, "less_eq", enum.IPredSLE)

}
func DefineAndOperation(program *lang.Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "and")
	result := binFnBody.NewAnd(a, b)
	//zexted := binFnBody.NewZExt(result, types.I32)
	binFnBody.NewCall(program.Funcs["push"].IrFunc, result)
	binFnBody.NewRet(nil)
	return binFn

}
func DefineOrOperation(program *lang.Program) *ir.Func {
	binFnBody, binFn, a, b := DefineBinaryFunction(program, "or")
	result := binFnBody.NewOr(a, b)
	//zexted := binFnBody.NewZExt(result, types.I32)
	binFnBody.NewCall(program.Funcs["push"].IrFunc, result)
	binFnBody.NewRet(nil)
	return binFn

}
func DefineNotOperation(program *lang.Program) *ir.Func {
	notFn := program.Module.NewFunc("not_op", types.Void, ir.NewParam("a", types.I32))
	notFnBody := notFn.NewBlock("entry")

	op1 := notFn.Params[0]
	op1T := notFnBody.NewTrunc(op1, types.I1)
	result := notFnBody.NewXor(constant.True, op1T)
	zexted := notFnBody.NewZExt(result, types.I32)
	notFnBody.NewCall(program.Funcs["push"].IrFunc, zexted)
	notFnBody.NewRet(nil)
	return notFn

}
