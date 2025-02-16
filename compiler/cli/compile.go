package cli

import (
	"bufio"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/asm"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"rpn/irCompiler"
	"rpn/lang"
	"rpn/lib"
	"rpn/parser"
	"rpn/util/config"
	"strings"
)

func Save(program *lang.Program, config config.TOMLConfig) {

	file, err := os.Create(safeJoin(getCurrentDir(), "output.ll"))
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	out := program.Module.String()
	out = strings.Replace(out, "declare i32 (i8*, ...) @printf()", "declare i32 @printf(i8*, ...)", -1)
	out = strings.Replace(out, "declare i32 (i8*, ...) @scanf()", "declare i32 @scanf(i8*, ...)", -1)

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(out)
	writer.Flush()
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
}

func CompileStcLibs(config config.TOMLConfig) {
	root := config.GetPath(config.Libs.LibRoot) + "/"
	for _, libPath := range config.Libs.LibRaw {
		base := filepath.Base(libPath)

		outPath := filepath.Join(config.GetPath(config.Libs.LibBin), base+".ll")
		cmd := exec.Command(config.Config.ClangPath, "-O2", "-S", "-emit-llvm", filepath.Join(root, libPath+".c"), "-o", outPath)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error compiling lib:", libPath, err)
			os.Exit(1)
		}
		fmt.Println("Compiled lib:", libPath)

	}
}

func Compile(config config.TOMLConfig, path string, compilationCfg CompilationConfig) {
	input, _ := antlr.NewFileStream(path)
	lex := parser.NewStcLexer(input)
	stream := antlr.NewCommonTokenStream(lex, 0)
	p := parser.NewStcParser(stream)
	p.AddErrorListener(antlr.NewDefaultErrorListener())
	tree := p.Prog()

	program := lang.NewProgram()
	program.Tree = tree
	program.Stream = stream
	program.Module.DataLayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"

	LoadLibs(config, program)
	lib.BindLibs(config, program)
	lib.GenerateDefinitions(program)

	irCompiler.LoadProgram(program)

	Save(program, config)
	LinkFiles(program, config)

	BuildBinary(config, compilationCfg)
}

func BuildBinary(config config.TOMLConfig, cfg CompilationConfig) {
	llPath := safeJoin(getCurrentDir(), "output.ll")

	cmd := exec.Command(
		config.Config.ClangPath, llPath, "-lm", "-fPIE", "-pie", getOpt(cfg.Optimization), "-o", cfg.OutFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error compiling binary:", err)
	}
	if cfg.EmitLLVM {
		cmd = exec.Command(config.Config.LLDisPath, llPath, "-o", cfg.OutFile+".ll")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error compiling binary:", err)
		}

	}

}

func getOpt(opt int) string {
	return fmt.Sprintf("-O%d", opt)
}

func LinkFiles(program *lang.Program, tomlConfig config.TOMLConfig) {
	args := getFilesToLink(tomlConfig)
	args = append(args, "-o")
	args = append(args, safeJoin(getCurrentDir(), "output.ll"))
	cmd := exec.Command(tomlConfig.Config.LinkerPath, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: Invalid linker process:", err)
		os.Exit(-1)
	}

}

func getFilesToLink(config config.TOMLConfig) []string {
	var str []string
	for _, libPath := range config.Libs.LibRaw {
		base := filepath.Base(libPath)
		ext := filepath.Ext(base)
		nameWithoutExt := strings.TrimSuffix(base, ext)
		outPath := safeJoin(config.GetPath(config.Libs.LibBin), nameWithoutExt+".ll")
		str = append(str, outPath)
	}
	str = append(str, safeJoin(getCurrentDir(), "output.ll"))
	return str

}

func LoadLibs(config config.TOMLConfig, program *lang.Program) {
	for _, libPath := range config.Libs.LibRaw {
		base := filepath.Base(libPath)

		libBinPath := filepath.Join(config.GetPath(config.Libs.LibBin), base+".ll")
		mod, err2 := asm.ParseFile(libBinPath)
		if err2 != nil {
			fmt.Println("Error loading lib:", libPath, err2)
			os.Exit(1)
		}
		program.StaticLibsModules = append(program.StaticLibsModules, mod)
	}
}
