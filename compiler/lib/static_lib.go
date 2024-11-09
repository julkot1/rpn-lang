package lib

import (
	"rpn/lang"
)

func isStcFunction(name string) bool {
	if len(name) < 4 {
		return false
	}
	return name[0] == 's' && name[1] == 't' && name[2] == 'c' && name[3] == '_'
}

func GenerateDefinitions(program *lang.Program) {
	for _, mod := range program.StaticLibsModules {
		for _, fun := range mod.Funcs {
			if isStcFunction(fun.Name()) {
				program.Module.NewFunc(fun.Name(), fun.Sig.RetType, fun.Params...)
			}
		}
	}
}

func MatchDefaultOperators(program *lang.Program) {

}
