package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	RootPath    string `toml:"root_path"`
	Destination string `toml:"destination"`
	CompileLibs bool   `toml:"compile_libs"`
	ClangPath   string `toml:"clang_path"`
	LlcPath     string `toml:"llc_path"`
	LLDisPath   string `toml:"llc_dis_path"`
	LinkerPath  string `toml:"linker_path"`
}

type Libs struct {
	LibRoot string   `toml:"lib_root"`
	LibBin  string   `toml:"lib_path"`
	LibRaw  []string `toml:"lib_raw"`
}

type TOMLConfig struct {
	Config Config `toml:"Config"`
	Libs   Libs   `toml:"Libs"`
}

func (c *TOMLConfig) GetPath(file string) string {
	return c.Config.RootPath + "/" + file
}

func CreateTOMLConfig(path string) TOMLConfig {
	var conf TOMLConfig
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		fmt.Println("Error loading config:", err)
		return conf
	}

	return conf
}
