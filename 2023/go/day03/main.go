package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

func main() {
	filename := "input.txt"

	fmt.Println("part1: ", part1(util.ReadFile(filename)))
	fmt.Println("part2: ", part2(util.ReadFile(filename)))
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c rune) bool {
	return c == '*' || c == '#' || c == '+' || c == '-' || c == '$' || c == '@' || c == '%' || c == '&' || c == '=' || c == '/'
}

func isAst(c rune) bool {
	return c == '*'
}

type position struct {
	i, j int
}

func part1(input string) int {
	lines := util.SplitLines(input)
	ans := 0

	surroundings := []position{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	re := regexp.MustCompile(`\d+`)
	for i, line := range lines {
		indexes := re.FindAllStringIndex(line, -1)
		for _, idx := range indexes {
			l, r := idx[0], idx[1]
			isValid := false
			for j := l; j < r; j++ {
				for _, s := range surroundings {
					ni, nj := i+s.i, j+s.j
					if ni >= 0 && ni < len(lines) && nj >= 0 && nj < len(line) {
						if isSymbol(rune(lines[ni][nj])) {
							isValid = true
							break
						}
					}
				}
				if isValid {
					break
				}
			}

			if isValid {
				num, _ := strconv.Atoi(string(line[l:r]))
				ans += num
			}
		}
	}

	return ans
}

type value struct {
	num int
	pos position
}

func compact(slice []value) []value {
	uniq := make(map[value]bool)
	var res []value

	for _, v := range slice {
		if _, ok := uniq[v]; !ok {
			uniq[v] = true
			res = append(res, v)
		}
	}
	return res
}

func part2(input string) int {
	lines := util.SplitLines(input)

	surroundings := []position{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	re := regexp.MustCompile(`\d+`)
	mp := map[position][]value{}
	for i, line := range lines {
		indexes := re.FindAllStringIndex(line, -1)
		for _, idx := range indexes {
			l, r := idx[0], idx[1]
			for j := l; j < r; j++ {
				for _, s := range surroundings {
					ni, nj := i+s.i, j+s.j
					if ni >= 0 && ni < len(lines) && nj >= 0 && nj < len(line) {

						if isAst(rune(lines[ni][nj])) {
							if _, ok := mp[position{ni, nj}]; !ok {
								mp[position{ni, nj}] = []value{}
							}
							num, _ := strconv.Atoi(string(line[l:r]))
							mp[position{ni, nj}] = append(mp[position{ni, nj}], value{num, position{i, l}})
							break
						}
					}
				}
			}
		}
	}

	ans := 0
	for _, nums := range mp {
		ns := compact(nums)
		if len(ns) < 2 {
			continue
		}
		ans += ns[0].num * ns[1].num
	}

	return ans
}
