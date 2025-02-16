package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

func safeJoin(base, second string) string {
	if filepath.IsAbs(second) {
		return second
	}
	return filepath.Join(base, second)
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return ""
	}
	return dir
}
