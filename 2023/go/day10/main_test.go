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
		{8, part1(util.ReadFile("test.txt"))},
		{6815, part1(util.ReadFile("input.txt"))},
	}

	testCases("part1", cases, t)
}

func TestPart2(t *testing.T) {
	cases := []Case{
		{10, part2(util.ReadFile("test2.txt"))},
		// {1026, part2(util.ReadFile("input.txt"))},
	}

	testCases("part2", cases, t)
}
