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
	// fmt.Println("part2: ", part2(string(text)))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	ans := 0

	colorLimits := map[string]int{"blue": 14, "red": 12, "green": 13}
	for i, line := range lines {
		re, err := regexp.Compile(`(\d{1,2} (blue|red|green))`)
		if err != nil {
			panic(err)
		}

		matches := re.FindAllString(line, -1)
		ok := true
		for _, match := range matches {
			color := strings.Split(match, " ")[1]
			num, err := strconv.Atoi(strings.Split(match, " ")[0])
			if err != nil {
				panic(err)
			}
			if num > colorLimits[color] {
				ok = false
				break
			}
		}
		if ok {
			ans += i + 1
		}
	}
	return ans
}

// func part2(input string) int {

// }
