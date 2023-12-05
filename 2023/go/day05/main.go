package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

type Rule struct {
	dest int
	src  int
	ran  int
}

func main() {
	filename := "input.txt"

	fmt.Println("part1: ", part1(util.ReadFile(filename)))
	fmt.Println("part2: ", part2(util.ReadFile(filename)))
}

func extractSeeds(s string) []int {
	var res []int
	for i, v := range strings.Split(s, " ") {
		if i == 0 {
			continue
		}
		res = append(res, util.Atoi(v))
	}
	return res
}

func applyMaps(maps []func(int) int, seed int) int {
	for _, m := range maps {
		seed = m(seed)
	}
	return seed
}

func srcToDestMap(rules []Rule) func(int) int {
	return func(i int) int {
		for _, rule := range rules {
			diff := rule.dest - rule.src
			if rule.src <= i && i < rule.src+rule.ran {
				return i + diff
			}
		}
		return i
	}
}

func extractRule(s string) Rule {
	str := strings.Split(s, " ")
	return Rule{src: util.Atoi(str[1]), dest: util.Atoi(str[0]), ran: util.Atoi(str[2])}
}

func extractRules(s string) []Rule {
	var res []Rule
	for i, v := range strings.Split(s, "\n") {
		if i == 0 {
			continue
		}
		res = append(res, extractRule(v))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part1(input string) int {
	blocks := strings.Split(input, "\n\n")
	seeds := extractSeeds(blocks[0])
	var maps []func(int) int
	for _, v := range blocks[1:] {
		rules := extractRules(v)
		maps = append(maps, srcToDestMap(rules))
	}
	m := math.MaxInt64
	for _, num := range seeds {
		m = min(m, applyMaps(maps, num))
	}

	return m
}

type Range struct {
	start int
	end   int
}

func extractSeedRanges(s string) []Range {
	var res []Range
	var nums []int
	for i, v := range strings.Split(s, " ") {
		if i == 0 {
			continue
		}
		nums = append(nums, util.Atoi(v))
	}

	for i := 0; i < len(nums); i += 2 {
		res = append(res, Range{start: nums[i], end: nums[i] + nums[i+1] - 1})
	}
	return res
}

func part2(input string) int {
	blocks := strings.Split(input, "\n\n")
	ranges := extractSeedRanges(blocks[0])
	var maps []func(int) int
	for _, v := range blocks[1:] {
		rules := extractRules(v)
		maps = append(maps, srcToDestMap(rules))
	}

	m := math.MaxInt64
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			m = min(m, applyMaps(maps, i))
		}
	}
	return m
}
