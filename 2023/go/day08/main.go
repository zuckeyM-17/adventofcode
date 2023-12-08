package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

func main() {
	filename := "input.txt"

	fmt.Println("part1: ", part1(util.ReadFile(filename)))
	fmt.Println("part2: ", part2(util.ReadFile(filename)))
}

func extractNetwork(lines []string) map[string][]string {
	network := map[string][]string{}

	for _, line := range lines {
		words := strings.Split(line, " = ")
		re, err := regexp.Compile(`([A-Z\d]{3})`)
		if err != nil {
			panic(err)
		}
		lr := re.FindAllString(words[1], -1)
		network[words[0]] = []string{lr[0], lr[1]}
	}

	return network
}

func instructions(line string) []int {
	var res []int
	for _, c := range line {
		if c == 'R' {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	}

	return res
}

func part1(input string) int {
	blocks := util.SplitByEmptyLines(input)

	inst := instructions(blocks[0])
	network := extractNetwork(util.SplitLines((blocks[1])))

	cnt := 1
	i := 1
	cur := network["AAA"][inst[0]]
	for cur != "ZZZ" {
		if i >= len(inst) {
			i = 0
		}
		cur = network[cur][inst[i]]
		i++
		cnt++
	}

	return cnt
}

func part2(input string) int {
	blocks := util.SplitByEmptyLines(input)

	inst := instructions(blocks[0])
	network := extractNetwork(util.SplitLines((blocks[1])))

	nodes := []string{}
	for k, _ := range network {
		if k[2] == 'A' {
			nodes = append(nodes, k)
		}
	}

	cnts := []int{}

	for _, node := range nodes {
		cnt := 1
		i := 1
		cur := network[node][inst[0]]
		for cur[2] != 'Z' {
			if i >= len(inst) {
				i = 0
			}
			cur = network[cur][inst[i]]
			i++
			cnt++
		}
		cnts = append(cnts, cnt)
	}

	return lcm(cnts)
}

func lcm(nums []int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		res = res * num / gcd(res, num)
	}

	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
