package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zuckeyM-17/adventofcode/2023/go/util"
)

func main() {
	filename := "input.txt"

	fmt.Println("part1: ", part1(util.ReadFile(filename)))
	fmt.Println("part2: ", part2(util.ReadFile(filename)))
}

type Strength int

const (
	HighCard Strength = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	cards    []string
	bid      int
	strength Strength
}

func values(m map[string]int) []int {
	vs := []int{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func getStrengthWithoutJoker(cards []string) Strength {
	m := map[string]int{}
	for _, c := range cards {
		m[c]++
	}
	vs := values(m)
	sort.Sort(sort.Reverse(sort.IntSlice(vs)))
	switch true {
	case vs[0] == 5:
		return FiveOfAKind
	case vs[0] == 4:
		return FourOfAKind
	case vs[0] == 3 && vs[1] == 2:
		return FullHouse
	case vs[0] == 3:
		return ThreeOfAKind
	case vs[0] == 2 && vs[1] == 2:
		return TwoPairs
	case vs[0] == 2:
		return OnePair
	default:
		return HighCard
	}
}

func extractHand(s string, getStrength func([]string) Strength) hand {
	n := strings.Split(s, " ")
	cards := strings.Split(n[0], "")
	bid := util.Atoi(n[1])
	return hand{cards, bid, getStrength(cards)}
}

func compareHands(h1, h2 hand, cardRankMap map[string]int) bool {
	if h1.strength == h2.strength {
		return compareCards(h1, h2, cardRankMap)
	} else {
		return h1.strength < h2.strength
	}
}

func compareCards(h1, h2 hand, cardRankMap map[string]int) bool {
	for i, c1 := range h1.cards {
		if c1 == h2.cards[i] {
			continue
		}
		return cardRankMap[c1] < cardRankMap[h2.cards[i]]
	}
	return false
}

func part1(input string) int {
	hands := []hand{}
	for _, line := range util.SplitLines(input) {
		hands = append(hands, extractHand(line, getStrengthWithoutJoker))
	}

	m := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], m)
	})

	ans := 0
	for i, h := range hands {
		ans += h.bid * (i + 1)
	}

	return ans
}

func getStrengthWithJoker(cards []string) Strength {
	if strings.Join(cards, "") == "JJJJJ" {
		return FiveOfAKind
	}

	m := map[string]int{}
	for _, c := range cards {
		m[c]++
	}
	plus := 0
	if m["J"] > 0 {
		plus = m["J"]
		delete(m, "J")
	}

	vs := values(m)
	sort.Sort(sort.Reverse(sort.IntSlice(vs)))
	vs[0] += plus

	switch true {
	case vs[0] == 5:
		return FiveOfAKind
	case vs[0] == 4:
		return FourOfAKind
	case vs[0] == 3 && vs[1] == 2:
		return FullHouse
	case vs[0] == 3:
		return ThreeOfAKind
	case vs[0] == 2 && vs[1] == 2:
		return TwoPairs
	case vs[0] == 2:
		return OnePair
	default:
		return HighCard
	}
}

func part2(input string) int {
	hands := []hand{}
	for _, line := range util.SplitLines(input) {
		hands = append(hands, extractHand(line, getStrengthWithJoker))
	}

	m := map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], m)
	})

	ans := 0
	for i, h := range hands {
		ans += h.bid * (i + 1)
	}

	return ans
}
