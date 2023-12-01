package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `
	1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	`

	want := 142
	got := part1(input)

	t.Run("part1", func(t *testing.T) {
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	input := `
		two1nine
		eightwothree
		abcone2threexyz
		xtwone3four
		4nineeightseven2
		zoneight234
		7pqrstsixteen`

	want := 281
	get := part2(input)

	t.Run("part2", func(t *testing.T) {
		if get != want {
			t.Errorf("got %d want %d", get, want)
		}
	})
}
