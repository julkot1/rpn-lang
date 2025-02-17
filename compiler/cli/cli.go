package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"rpn/util/config"
)

type CompilationConfig struct {
	OutFile      string
	Optimization int
	EmitLLVM     bool
}
type CLI struct {
	Root           *cobra.Command
	CompilationCfg CompilationConfig
}

func CreateCli() *CLI {
	c := &CLI{}
	c.Root = &cobra.Command{Use: "stc"}
	c.Root.AddCommand(createCompileLibs(c))
	c.Root.AddCommand(createCompile(c))

	return c
}

func createCompile(cli *CLI) *cobra.Command {
	var cmdCompile = &cobra.Command{
		Use:  "compile [files...]",
		Args: cobra.MinimumNArgs(1),

		Short: "Compile program",
		Long: `Compile given files to executable file
Usage:
	compile [file names...] [flags]

Flags:
	--output, (-o) [output file]: Specifies the name of the output file.
	--emit-llvm -l:              Generates output in llvm
	-O2 -O3 -O:                   Specifies the optimization level

`,

		Run: func(cmd *cobra.Command, args []string) {
			// Validate optimization level
			if cli.CompilationCfg.Optimization < 0 || cli.CompilationCfg.Optimization > 3 {
				log.Fatalf("Error: -O level should be between 0 and 3")
			}

			fmt.Printf("Compiling files with optimization level -O%d:\n", cli.CompilationCfg.Optimization)
			//conf := config.CreateTOMLConfig("/etc/stc/stconfig.toml")
			conf := config.CreateTOMLConfig("config.debug.toml")
			for _, file := range args {
				Compile(conf, file, cli.CompilationCfg)
			}
		},
	}
	cmdCompile.Flags().StringVarP(&cli.CompilationCfg.OutFile, "output", "o", "out", "Specifies the name of the output file.")
	cmdCompile.Flags().IntVarP(&cli.CompilationCfg.Optimization, "opt", "O", 0, "Specifies the optimization level")
	cmdCompile.Flags().BoolVarP(&cli.CompilationCfg.EmitLLVM, "emit-llvm", "l", false, "Generates output in llvm")

	return cmdCompile
}

func createCompileLibs(cli *CLI) *cobra.Command {
	var cmdLibs = &cobra.Command{
		Use:   "libs",
		Short: "Compile libs",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Compile stc default libs")
			//conf := config.CreateTOMLConfig("/etc/stc/stconfig.toml")
			conf := config.CreateTOMLConfig("config.debug.toml")

			CompileStcLibs(conf)
		},
	}

	return cmdLibs
}
