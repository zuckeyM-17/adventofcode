package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

func main() {
	filename := "input.txt"

	fmt.Println("part1: ", part1(util.ReadFile(filename)))
	fmt.Println("part2: ", part2(util.ReadFile(filename)))
}

func compact(s []string) []int {
	var res []int

	for _, v := range s {
		if trimmedStr := strings.TrimSpace(v); trimmedStr != "" {
			num, _ := strconv.Atoi(trimmedStr)
			res = append(res, num)
		}
	}
	return res
}

func part1(input string) int {
	lines := util.SplitLines(input)

	prefix := `Card\s+\d+:\s+`
	re, err := regexp.Compile(prefix)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}
	ans := 0
	for _, line := range lines {
		l := strings.Split(re.ReplaceAllString(line, ""), "|")
		wins, nums := compact(strings.Split(l[0], " ")), compact(strings.Split(l[1], " "))
		cnt := 0
		for _, v := range nums {
			for _, w := range wins {
				if v == w {
					cnt++
					break
				}
			}
		}
		var w int
		if cnt == 1 {
			w = 1
		} else {
			w = int(math.Pow(2, float64(cnt-1)))
		}

		ans += w
	}

	return ans
}

func part2(input string) int {
	lines := util.SplitLines(input)

	prefix := `Card\s+\d+:\s+`
	re, err := regexp.Compile(prefix)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}
	cards := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cards[i] = 1
	}
	for i, line := range lines {
		l := strings.Split(re.ReplaceAllString(line, ""), "|")
		wins, nums := compact(strings.Split(l[0], " ")), compact(strings.Split(l[1], " "))
		cnt := 0
		for _, v := range nums {
			for _, w := range wins {
				if v == w {
					cnt++
					break
				}
			}
		}

		for j := 1; j <= cnt; j++ {
			cards[i+j] += cards[i]
		}
	}

	ans := 0
	for _, v := range cards {
		ans += v
	}

	return ans
}
