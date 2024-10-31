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
	"rpn/lib"
	"rpn/util/config"
	"strings"
)

func Save(program *lang.Program, config config.TOMLConfig) {

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
	program := lexer.Parse(config.GetPath(path))
	program.Module.DataLayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
	if config.Config.CompileLibs {
		CompileLibs(config, program)
	}
	lib.GenerateDefinitions(program)

	irCompiler.LoadProgram(program)
	Save(program, config)
	LinkFiles(program, config)
	ExecScript(config)

}

func LinkFiles(program *lang.Program, tomlConfig config.TOMLConfig) {
	args := getFilesToLink(tomlConfig)
	args = append(args, "-o")
	args = append(args, tomlConfig.GetPath(tomlConfig.Config.Destination)+"/output.ll")
	cmd := exec.Command(tomlConfig.Config.LinkerPath, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: Invalid linker process:", err)
		return
	}

}

func getFilesToLink(config config.TOMLConfig) []string {
	var str []string
	for _, libPath := range config.Libs.LibRaw {
		base := filepath.Base(libPath)
		ext := filepath.Ext(base)
		nameWithoutExt := strings.TrimSuffix(base, ext)
		outPath := config.GetPath(config.Config.Destination) + "/" + nameWithoutExt + ".ll"
		str = append(str, outPath)
	}
	str = append(str, config.GetPath(config.Config.Destination)+"/output.ll")
	return str

}

func CompileLibs(config config.TOMLConfig, program *lang.Program) {
	root := config.GetPath(config.Libs.LibRoot) + "/"
	for _, libPath := range config.Libs.LibRaw {
		base := filepath.Base(libPath)
		ext := filepath.Ext(base)
		nameWithoutExt := strings.TrimSuffix(base, ext)
		outPath := config.GetPath(config.Config.Destination) + "/" + nameWithoutExt + ".ll"
		cmd := exec.Command(config.Config.ClangPath, "-O2", "-S", "-emit-llvm", root+libPath, "-o", outPath)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error compiling lib:", libPath, err)
			return
		}
		mod, err2 := asm.ParseFile(outPath)
		if err2 != nil {
			fmt.Println("Error loading lib:", libPath, err2)
			return
		}
		program.StaticLibsModules = append(program.StaticLibsModules, mod)
		fmt.Println("Compiled lib:", libPath)

	}
}
