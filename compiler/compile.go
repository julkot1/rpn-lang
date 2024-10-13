package main

import (
	"bufio"
	"fmt"
	"github.com/llir/llvm/asm"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"rpn/irCompiler"
	"rpn/lang"
	"rpn/lexer"
	"rpn/util/config"
	"strings"
)

func Save(program *lang.Program, config config.TOMLConfig) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Print the current directory
	fmt.Println("Current working directory:", dir)

	file, err := os.Create(config.GetPath(config.Config.Destination + "/output.ll"))
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
func ExecScript(config config.TOMLConfig) {
	cmd := exec.Command(config.GetPath(config.Config.CompilationScript))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running script:", err)
		return
	}

	fmt.Println("Compilation script successful, binary created")
}

func Compile(config config.TOMLConfig, path string) {
	if config.Config.CompileLibs {
		CompileLibs(config)
	}

	program := lexer.Parse(config.GetPath(path))
	irCompiler.LoadProgram(program)
	Save(program, config)
	ExecScript(config)

}

func CompileLibs(config config.TOMLConfig) {
	root := config.GetPath(config.Libs.LibRoot) + "/"
	for _, libPath := range config.Libs.LibRaw {
		base := filepath.Base(libPath)
		ext := filepath.Ext(base)
		nameWithoutExt := strings.TrimSuffix(base, ext)
		outPath := config.GetPath(config.Config.Destination) + "/" + nameWithoutExt + ".ll"
		cmd := exec.Command(config.Config.ClangPath, "-S", "-emit-llvm", root+libPath, "-o", outPath)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error compiling lib:", libPath, err)
			return
		}
		_, err2 := asm.ParseFile(outPath)
		if err2 != nil {
			fmt.Println("Error loading lib:", libPath, err2)
			return
		}
		fmt.Println("Compiled lib:", libPath)

	}
}
