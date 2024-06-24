package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Save(program *Program) {
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
	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(out)
	writer.Flush()
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
}

func main() {

	tokens := Parse("app.rpn")
	program := NewProgram()
	LoadProgram(program, tokens)
	Save(program)

}
