package main

import (
	"fmt"

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

func part1(input string) int {
	lines := util.SplitLines(input)

	ans := 0
	for _, line := range lines {
		s := make([][]int, len(line)-1)
		i := 0
		s[i] = make([]int, len(util.SplitBySpace(line)))
		for j, n := range util.SplitBySpace(line) {
			s[i][j] = util.Atoi(n)
		}
		for !allZero(s[i]) {
			s[i+1] = make([]int, len(s[i])-1)
			for j := 0; j < len(s[i])-1; j++ {
				s[i+1][j] = s[i][j+1] - s[i][j]
			}
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
		s[i] = make([]int, len(util.SplitBySpace(line)))
		for j, n := range util.SplitBySpace(line) {
			s[i][j] = util.Atoi(n)
		}
		for !allZero(s[i]) {
			s[i+1] = make([]int, len(s[i])-1)
			for j := 0; j < len(s[i])-1; j++ {
				s[i+1][j] = s[i][j+1] - s[i][j]
			}
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
