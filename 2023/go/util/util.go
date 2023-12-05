package util

import (
	"os"
	"strconv"
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

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return num
}
