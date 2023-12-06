package main

import (
	"fmt"
	"regexp"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

type race struct {
	time int
	dist int
}

func main() {
	filename := "input.txt"

	fmt.Println("part1: ", part1(util.ReadFile(filename)))
	// fmt.Println("part2: ", part2(util.ReadFile(filename)))
}

func extractNums(s string) []int {
	var res []int
	re, err := regexp.Compile(`(\d+)`)
	if err != nil {
		panic(err)
	}
	for _, v := range re.FindAllString(s, -1) {
		res = append(res, util.Atoi(v))
	}

	return res
}

func countRace(r race) int {
	cnt := 0
	for i := 1; i <= r.time; i++ {
		if i*(r.time-i) > r.dist {
			cnt++
		}
	}
	return cnt
}

func part1(input string) int {
	lines := util.SplitLines(input)
	times, dists := extractNums(lines[0]), extractNums(lines[1])
	ans := 1
	for i, v := range times {
		ans *= countRace(race{v, dists[i]})
	}

	return ans
}
