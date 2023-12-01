package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func word2int(word string) int {
	num, e := strconv.Atoi(word)
	if e == nil {
		return num
	}
	nums := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9}
	return nums[word]
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	pattern := `\d`
	return sumTwoDigits(lines, pattern)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	pattern := `\d|one|two|three|four|five|six|seven|eight|nine`
	return sumTwoDigits(lines, pattern)
}

func sumTwoDigits(lines []string, pattern string) int {
	ans := 0
	for _, line := range lines {
		re, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}

		matches := re.FindAllString(line, -1)
		tensPlace := word2int(matches[0])
		onesPlace := word2int(matches[len(matches)-1])
		ans += onesPlace + tensPlace*10
	}
	return ans
}
