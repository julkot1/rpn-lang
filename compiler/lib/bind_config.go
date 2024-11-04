package lib

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/llir/llvm/ir"
	"log"
	"os"
	"path/filepath"
	"rpn/lang"
	"rpn/util/config"
)

// LibConfig represents the full TOML structure
type LibConfig struct {
	Head Head `toml:"head"`
	Body Body `toml:"body"`
}

// Head represents the [head] section
type Head struct {
	Name     string    `toml:"name"`
	Includes []string  `toml:"includes"`
	Types    HeadTypes `toml:"types"`
}

// HeadTypes represents the [head.types] table within [head]
type HeadTypes struct {
	TypeName string      `toml:"type_name"`
	Name     string      `toml:"name"`
	Args     []string    `toml:"args"`
	Return   string      `toml:"return"`
	Method   Method      `toml:"method"`
	Match    []TypeMatch `toml:"match"`
}

// Method represents the [head.types.method] table within [head.types]
type Method struct {
	Name    string   `toml:"name"`
	Stc     bool     `toml:"stc"`
	Args    []string `toml:"args"`
	Return  string   `toml:"return"`
	Code    []string `toml:"code"`
	StcCode int      `toml:"stc_code"`
	StcName *string  `toml:"stc_name"`
}

// TypeMatch represents each [[head.types.match]] entry
type TypeMatch struct {
	ArgA     string `toml:"argA"`
	ArgB     string `toml:"argB"`
	Function string `toml:"function"`
}

// Body represents the [body] section
type Body struct {
	Method []Method `toml:"method"`
}

func LoadLibConfig(path string) *LibConfig {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open TOML file: %v", err)
	}
	defer file.Close()

	config := &LibConfig{}
	if _, err := toml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Failed to decode TOML: %v", err)
	}
	return config
}

type LibsConfig struct {
	RootInput  string   `toml:"root_input"`
	RootOutput string   `toml:"root_output"`
	Libs       []string `toml:"libs"`
}

func LoadConfig(path string) *LibsConfig {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open TOML file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode the TOML into the Config struct
	config := &LibsConfig{}
	if _, err := toml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Failed to decode TOML file: %v", err)
		os.Exit(1)
	}
	return config

}

func BindLibs(config config.TOMLConfig, program *lang.Program) {
	for _, libName := range config.Libs.LibRaw {

		libBindConfig := LoadLibConfig(filepath.Join(config.GetPath(config.Libs.LibConfigRoot), libName+".toml"))
		methods := getMethods(libBindConfig)
		bindMethods(methods, program, libName+".c")
	}
	return
}

func bindMethods(methods []*StcMethod, program *lang.Program, libName string) {
	for _, stcMethod := range methods {
		if stcMethod.Code == 255 {
			stcName := *stcMethod.Name
			_, exist := program.GlobalTokenTable[stcName]
			if exist {
				fmt.Println("Token " + stcName + " already exists!")
				os.Exit(1)
			}
			program.GlobalTokenTable[stcName] = lang.PStaticFunc
			fun := getIr(stcMethod, program, stcName, libName)
			typ := fun.Sig
			program.StaticFunctions[stcName] = &lang.DefaultFunc{
				IrFunc: fun,
				Type:   typ,
			}
			continue
		}
		fun := getIr(stcMethod, program, stcMethod.CName, libName)
		program.Funcs[stcMethod.Code] = &lang.DefaultFunc{IrFunc: fun}

	}
}

func getIr(method *StcMethod, program *lang.Program, name string, libName string) *ir.Func {
	module := getModule(program, libName)
	for _, fun := range module.Funcs {
		if fun.Name() == method.CName {
			return fun
		}
	}
	fmt.Println("Function " + method.CName + " not found!")
	os.Exit(1)
	return nil

}

func getModule(program *lang.Program, name string) *ir.Module {
	for _, module := range program.StaticLibsModules {

		if filepath.Base(module.SourceFilename) == name {
			return module
		}
	}
	fmt.Println("No module found!")
	os.Exit(1)
	return nil
}

type StcMethod struct {
	CName string
	Code  lang.TokenType
	Argc  int
	Name  *string
}

func toStcMethod(method Method) *StcMethod {
	stcMethod := &StcMethod{CName: "stc_" + method.Name, Code: lang.TokenType(method.StcCode), Name: method.StcName, Argc: len(method.Args)}
	return stcMethod
}
func getMethods(bindConfig *LibConfig) []*StcMethod {
	var methods []*StcMethod
	if bindConfig.Head.Types.Name != "" {
		if bindConfig.Head.Types.Method.Stc {
			methods = append(methods, toStcMethod(bindConfig.Head.Types.Method))
		}
	}
	for _, method := range bindConfig.Body.Method {
		if method.Stc {
			methods = append(methods, toStcMethod(method))
		}
	}
	return methods
}
