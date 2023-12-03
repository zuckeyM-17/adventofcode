package main

import (
	"testing"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

type Case struct {
	want int
	got  int
}

func TestPart1(t *testing.T) {
	cases := []Case{
		{4361, part1(util.ReadFile("test.txt"))},
		{535235, part1(util.ReadFile("input.txt"))},
	}

	t.Run("part1", func(t *testing.T) {
		for _, c := range cases {
			got, want := c.got, c.want
			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})
}

func TestPart2(t *testing.T) {
	cases := []Case{
		{467835, part2(util.ReadFile("test.txt"))},
		{79844424, part2(util.ReadFile("input.txt"))},
	}

	t.Run("part2", func(t *testing.T) {
		for _, c := range cases {
			got, want := c.got, c.want
			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})
}
