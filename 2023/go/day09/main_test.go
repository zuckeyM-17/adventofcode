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
		{114, part1(util.ReadFile("test.txt"))},
		{1584748274, part1(util.ReadFile("input.txt"))},
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
		{2, part2(util.ReadFile("test.txt"))},
		{1026, part2(util.ReadFile("input.txt"))},
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
