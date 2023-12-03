package util

import (
	"os"
	"strings"
)

func ReadFile(filename string) string {
	text, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(text)
}

func SplitLines(text string) []string {
	return strings.Split(text, "\n")
}
