package main

import (
	"fmt"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

func main() {
	filename := "input.txt"

	fmt.Println("part1: ", part1(util.ReadFile(filename)))
	fmt.Println("part2: ", part2(util.ReadFile(filename)))
}

type Position struct {
	i, j int
}

type Pipe struct {
	p Position
	c rune
}

func findSecond(p Position, lines []string) []Position {
	type Candidate struct {
		p Position
		c []rune
	}

	candidates := []Candidate{
		{Position{-1, 0}, []rune{'7', 'F', '|'}},
		{Position{0, -1}, []rune{'L', 'F', '-'}},
		{Position{0, 1}, []rune{'J', '7', '-'}},
		{Position{1, 0}, []rune{'L', 'J', '|'}},
	}

	var next []Position
	for _, s := range candidates {
		i, j := p.i+s.p.i, p.j+s.p.j
		if i < 0 || j < 0 {
			continue
		} else if i >= len(lines) || j >= len(lines[0]) {
			continue
		}

		t := lines[i][j]
		for _, c := range s.c {
			if c == rune(t) {
				next = append(next, Position{i, j})
			}
		}
	}

	return next
}

func findNext(cur Pipe, lines []string) []Pipe {
	rule := map[rune][]Position{
		'|': {{1, 0}, {-1, 0}},
		'-': {{0, 1}, {0, -1}},
		'7': {{1, 0}, {0, -1}},
		'J': {{-1, 0}, {0, -1}},
		'L': {{-1, 0}, {0, 1}},
		'F': {{1, 0}, {0, 1}},
	}
	var next []Pipe
	for _, p := range rule[cur.c] {
		i, j := cur.p.i+p.i, cur.p.j+p.j
		next = append(next, Pipe{Position{i, j}, rune(lines[i][j])})
	}
	return next
}

func findPipes(input string) []Pipe {
	pipes := []Pipe{}
	lines := util.SplitLines(input)
	for i, line := range lines {
		for j, c := range line {
			if c == 'S' {
				pipes = append(pipes, Pipe{Position{i, j}, c})
				np := findSecond(pipes[0].p, lines)[0]
				pipes = append(pipes, Pipe{np, rune(lines[np.i][np.j])})
			}
		}
	}

	for pipes[len(pipes)-1].c != 'S' {
		cur, last := pipes[len(pipes)-1], pipes[len(pipes)-2]
		np := findNext(cur, lines)
		if np[0].p.i == last.p.i && np[0].p.j == last.p.j {
			cur = np[1]
		} else {
			cur = np[0]
		}
		pipes = append(pipes, cur)
	}
	return pipes
}

func part1(input string) int {
	return len(findPipes(input)) / 2
}

func countTrue(l []rune) int {
	cnt := 0
	for _, c := range l {
		if c != '.' && c != '-' {
			cnt++
		}
	}
	return cnt
}

func isOutside(p Position, lines [][]rune) bool {
	line := lines[p.i]
	if countTrue(line[:p.j]) == 0 || countTrue(line[p.j+1:]) == 0 {
		return true
	}

	cnt := 0
	for i := 0; i < p.i; i++ {
		if lines[i][p.j] != '.' {
			cnt++
		}
	}
	if cnt == 0 {
		return true
	}

	cnt = 0
	for i := p.i + 1; i < len(lines); i++ {
		if lines[i][p.j] != '.' {
			cnt++
		}
	}
	if cnt == 0 {
		return true
	}

	return false
}

func part2(input string) int {
	pipes := findPipes(input)
	lines := util.SplitLines(input)
	m := make([][]rune, len(lines))
	for i, line := range lines {
		m[i] = make([]rune, len(line))
		for j := range line {
			m[i][j] = '.'
		}
	}

	for _, p := range pipes {
		m[p.p.i][p.p.j] = p.c
	}

	cnt := 0
	for i, l := range m {
		for j, c := range l {
			if c != '.' {
				continue
			}

			if !isOutside(Position{i, j}, m) && countTrue(l[:j]) != countTrue(l[j+1:]) {
				cnt++
				m[i][j] = '@'
			}
		}
	}

	for _, l := range m {
		fmt.Println(string(l))
	}

	fmt.Println(cnt)

	return cnt
}
