package main

import (
	"testing"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

type Case struct {
	want int
	got  int
}

func testCases(name string, cases []Case, t *testing.T) {
	t.Run(name, func(t *testing.T) {
		for _, c := range cases {
			got, want := c.got, c.want
			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})
}

func TestPart1(t *testing.T) {
	cases := []Case{
		{374, part1(util.ReadFile("test.txt"))},
		{10033566, part1(util.ReadFile("input.txt"))},
	}

	testCases("part1", cases, t)
}

func TestPart2(t *testing.T) {
	cases := []Case{
		{560822911938, part2(util.ReadFile("input.txt"))},
	}

	testCases("part2", cases, t)
}
