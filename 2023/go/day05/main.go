package main

import (
	"fmt"
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

func min(s []int) int {
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func part1(input string) int {
	blocks := strings.Split(input, "\n\n")
	seeds := extractSeeds(blocks[0])
	var maps []func(int) int
	for _, v := range blocks[1:] {
		rules := extractRules(v)
		maps = append(maps, srcToDestMap(rules))
	}
	var dests []int
	for _, m := range seeds {
		dests = append(dests, applyMaps(maps, m))
	}

	return min(dests)
}
