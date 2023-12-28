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

type position struct {
	r, c int
}

func countGalaxy(g [][]rune) ([]bool, []bool, []position) {
	columns := make([]bool, len(g[0]))
	rows := make([]bool, len(g))
	positions := []position{}
	for row, line := range g {
		isEmpty := true
		for column, c := range line {
			if c == '#' {
				columns[column] = true
				positions = append(positions, position{row, column})
				isEmpty = false
			}
		}
		if !isEmpty {
			rows[row] = true
		}
	}

	return rows, columns, positions
}

func combinations(s []position) [][]position {
	var result [][]position
	for i, p := range s {
		for j := i + 1; j < len(s); j++ {
			result = append(result, []position{p, s[j]})
		}
	}
	return result
}

func part1(input string) int {
	lines := util.SplitLines(input)
	galaxy := make([][]rune, len(lines))
	for i, line := range lines {
		galaxy[i] = []rune(line)
	}
	rows, columns, positions := countGalaxy(galaxy)

	ans := 0
	c := combinations(positions)
	for _, comb := range c {
		p1, p2 := comb[0], comb[1]
		l, r := util.Min(p1.c, p2.c), util.Max(p1.c, p2.c)
		t, b := util.Min(p1.r, p2.r), util.Max(p1.r, p2.r)

		add := 0
		for i := l; i <= r; i++ {
			if !columns[i] {
				add++
			}
		}

		for i := t; i <= b; i++ {
			if !rows[i] {
				add++
			}
		}

		ans += add + r - l + b - t
	}
	return ans
}

func part2(input string) int {
	lines := util.SplitLines(input)
	galaxy := make([][]rune, len(lines))
	for i, line := range lines {
		galaxy[i] = []rune(line)
	}
	rows, columns, positions := countGalaxy(galaxy)

	ans := 0
	c := combinations(positions)
	for _, comb := range c {
		p1, p2 := comb[0], comb[1]
		l, r := util.Min(p1.c, p2.c), util.Max(p1.c, p2.c)
		t, b := util.Min(p1.r, p2.r), util.Max(p1.r, p2.r)

		add := 0
		for i := l; i <= r; i++ {
			if !columns[i] {
				add++
			}
		}

		for i := t; i <= b; i++ {
			if !rows[i] {
				add++
			}
		}

		ans += add*999_999 + r - l + b - t
	}
	return ans
}
