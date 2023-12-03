package main

import (
	"testing"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

func TestPart1(t *testing.T) {

	want := 4361
	got := part1(util.ReadFile("test.txt"))

	t.Run("part1", func(t *testing.T) {
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
