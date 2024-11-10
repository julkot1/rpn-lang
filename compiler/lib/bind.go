package lib

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"path/filepath"
	"rpn/lang"
	"rpn/lib/decorators"
	"rpn/util/config"
)

func BindLibs(config config.TOMLConfig, program *lang.Program) {
	root := config.GetPath(config.Libs.LibRoot) + "/"
	for _, libPath := range config.Libs.LibRaw {
		base := filepath.Base(libPath)
		path := filepath.Join(root, base+".h")
		loadDecorators, err := decorators.LoadDecorators(path)
		if err != nil {
			fmt.Printf("Error loading decorators in file "+base+": %s\n", err)
		}
		var mod *ir.Module
		for _, m := range program.StaticLibsModules {
			if filepath.Base(m.SourceFilename) == (base + ".c") {
				mod = m
				break
			}
		}
		bindFunctions(program, loadDecorators, mod)

	}
}

func bindFunctions(program *lang.Program, fileDecorators *decorators.FileDecorators, mod *ir.Module) {
	for _, fun := range mod.Funcs {
		dec, f := fileDecorators.Decorators[(fun.Name())]
		if f {
			if dec.Code == 255 {
				if dec.Name == "" {
					continue
				}
				program.StaticFunctions[dec.Name] = &lang.DefaultFunc{IrFunc: fun, Type: fun.Sig}
				program.GlobalTokenTable[dec.Name] = lang.PStaticFunc
			} else {
				program.Funcs[lang.TokenType(dec.Code)] = &lang.DefaultFunc{IrFunc: fun, Type: fun.Sig}
			}
		}
	}
}
