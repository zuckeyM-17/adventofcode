package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	text, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1: ", part1(string(text)))
	fmt.Println("part2: ", part2(string(text)))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	ans := 0
	for _, line := range lines {
		var onesPlace, tensPlace int
		for _, c := range line {
			if isInt(c) {
				tensPlace = int(c - '0')
				break
			}
		}

		for _, c := range reverse(line) {
			if isInt(c) {
				onesPlace = int(c - '0')
				break
			}
		}
		ans += onesPlace + tensPlace*10
	}
	return ans
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	ans := 0
	numStrs := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	rNumStrs := map[string]int{
		"eno":   1,
		"owt":   2,
		"eerht": 3,
		"ruof":  4,
		"evif":  5,
		"xis":   6,
		"neves": 7,
		"thgie": 8,
		"enin":  9,
	}
	for _, line := range lines {
		var onesPlace, tensPlace int

		for l := 0; l < len(line); l++ {
			if isInt(rune(line[l])) {
				tensPlace = int(line[l] - '0')
				break
			}
			r := l + 2
			for r-l < 6 && r < len(line) {
				if val, ok := numStrs[line[l:r]]; ok {
					tensPlace = val
					break
				}
				r++
			}
			if tensPlace != 0 {
				break
			}
		}

		rLine := reverse(line)
		for l := 0; l < len(rLine); l++ {
			if isInt(rune(rLine[l])) {
				onesPlace = int(rLine[l] - '0')
				break
			}
			r := l + 2
			for r-l < 6 && r < len(rLine) {
				if val, ok := rNumStrs[rLine[l:r]]; ok {
					onesPlace = val
					break
				}
				r++
			}
			if onesPlace != 0 {
				break
			}
		}
		ans += onesPlace + tensPlace*10
	}
	return ans
}

func isInt(c rune) bool {
	return c >= '1' && c <= '9'
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
