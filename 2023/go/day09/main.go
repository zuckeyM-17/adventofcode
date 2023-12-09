package main

import (
	"fmt"
	"strings"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

func main() {
	filename := "input.txt"

	fmt.Println("part1: ", part1(util.ReadFile(filename)))
	fmt.Println("part2: ", part2(util.ReadFile(filename)))
}

func allZero(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}
	return true
}

func calcDiff(slice []int) []int {
	diff := make([]int, len(slice)-1)
	for i := 0; i < len(slice)-1; i++ {
		diff[i] = slice[i+1] - slice[i]
	}
	return diff
}

func extractNums(text string) []int {
	res := []int{}
	for _, n := range strings.Split(text, " ") {
		res = append(res, util.Atoi(n))
	}
	return res
}

func part1(input string) int {
	lines := util.SplitLines(input)

	ans := 0
	for _, line := range lines {
		s := make([][]int, len(line)-1)
		i := 0
		s[i] = extractNums(line)
		for !allZero(s[i]) {
			s[i+1] = calcDiff(s[i])
			i++
		}
		t := 0
		for ; i >= 0; i-- {
			t += s[i][len(s[i])-1]
		}
		ans += t
	}

	return ans
}

func part2(input string) int {
	lines := util.SplitLines(input)

	ans := 0
	for _, line := range lines {
		s := make([][]int, len(line)-1)
		i := 0
		s[i] = extractNums(line)
		for !allZero(s[i]) {
			s[i+1] = calcDiff(s[i])
			i++
		}
		t := 0
		for ; i >= 0; i-- {
			t = s[i][0] - t
		}

		ans += t
	}

	return ans
}
