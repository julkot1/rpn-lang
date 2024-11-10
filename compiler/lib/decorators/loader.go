package decorators

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type FileDecorators struct {
	Decorators map[string]STCBindConfig
}

func LoadDecorators(path string) (*FileDecorators, error) {
	fileDecorators := &FileDecorators{Decorators: map[string]STCBindConfig{}}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	err = findDecorators(fileDecorators, file)
	if err != nil {
		return nil, err
	}
	return fileDecorators, nil
}

func findDecorators(decorators *FileDecorators, file *os.File) error {
	scanner := bufio.NewScanner(file)
	prevLine := ""
	for scanner.Scan() {
		line := scanner.Text()
		if isDecorator(prevLine) && isFunctionDef(line) {
			name := getFuncName(line)
			decorator, err := ParseStcHeader(prevLine)
			if err != nil {
				return err
			}
			parseDecorator(name, decorator, decorators)

		}
		prevLine = line
	}
	return nil
}

func parseDecorator(name string, decorator *STCBindConfig, decorators *FileDecorators) {
	decorators.Decorators[name] = *decorator
}
func isDecorator(line string) bool {
	return strings.HasPrefix(line, "STC(")
}
func getFuncName(line string) string {
	strArray := strings.Split(line, " ")
	s := strArray[1]
	end := 0
	for idx, val := range s {
		if val == '(' {
			end = idx
			break
		}
	}
	return s[:end]
}
func isFunctionDef(line string) bool {
	funcDeclRegex := regexp.MustCompile(`^\s*\w+\s+\w+\s*\([^)]*\)\s*;?\s*$`)
	return funcDeclRegex.MatchString(line)
}
