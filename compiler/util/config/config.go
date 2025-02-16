package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
)

type Config struct {
	RootPath   string `toml:"root_path"`
	ClangPath  string `toml:"clang_path"`
	LLDisPath  string `toml:"llc_dis_path"`
	LinkerPath string `toml:"linker_path"`
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
	return filepath.Join(c.Config.RootPath, file)
}

func CreateTOMLConfig(path string) TOMLConfig {
	var conf TOMLConfig
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		fmt.Println("Error loading config:", err)
		return conf
	}

	return conf
}
