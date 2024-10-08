package main

import (
	"bufio"
	"log"
	"os"
	"rpn/irCompiler"
	"rpn/lang"
	"rpn/lexer"
	"strings"
)

func Save(program *lang.Program) {
	file, err := os.Create("output.ll")
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

func main() {

	program := lexer.Parse("app.rpn")
	irCompiler.LoadProgram(program)
	Save(program)

}
