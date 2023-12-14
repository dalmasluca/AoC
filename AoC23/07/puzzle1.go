package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	bit   int
}

var score = "AKQJT98765432"

func isFive(cards string) (bool, string) {
	first := cards[0]
	for i := range cards {
		if cards[i] != first {
			return false, ""
		}
	}
	return true, string(first)
}

func isFour(cards string) bool {
	m := make(map[rune]int, 0)
	for _, card := range cards {
		m[card]++
	}
	for _, v := range m {
		if v == 4 {
			return true
		}
	}
	return false
}

func isFull(cards string) bool {
	m := make(map[rune]int, 0)
	for _, card := range cards {
		m[card]++
	}
	if len(m) != 2 {
		return false
	}
	return true
}

func isThree(cards string) bool {
	m := make(map[rune]int, 0)
	for _, card := range cards {
		m[card]++
	}
	for _, v := range m {
		if v == 3 {
			return true
		}
	}
	return false
}

func firstbigger(set1, set2 string) bool {
	for i := 0; i < len(set1); i++ {
		if set1[i] != set2[i] {
			return strings.Index(score, string(set1[i])) < strings.Index(score, string(set2[i]))
		}
	}
	return false
}

func twoPair(cards string) bool {
	m := make(map[rune]int, 0)
	for _, card := range cards {
		m[card]++
	}
	count := 0
	for _, v := range m {
		if v == 2 {
			count++
		}
	}
	return count == 2
}

func onePair(cards string) bool {
	m := make(map[rune]int, 0)
	for _, card := range cards {
		m[card]++
	}
	pairs := 0
	for _, v := range m {
		if v == 2 {
			pairs++
		}
	}
	return pairs == 1
}

func higercard(cards string) int {
	var max int = len(score)
	for _, card := range cards {
		i := strings.Index(score, string(card))
		if i < max {
			max = i
		}
	}
	return max
}

func compare(set1, set2 string) bool {
	b, r := isFive(set1)
	b1, r1 := isFive(set2)
	if b || b1 {
		if b && b1 {
			return strings.Index(score, r) < strings.Index(score, r1)
		}
		return b
	}
	b = isFour(set1)
	b1 = isFour(set2)
	if b || b1 {
		if b && b1 {
			return firstbigger(set1, set2)
		}
		return b
	}
	b = isFull(set1)
	b1 = isFull(set2)
	if b || b1 {
		if b && b1 {
			return firstbigger(set1, set2)
		}
		return b
	}
	b = isThree(set1)
	b1 = isThree(set2)
	if b || b1 {
		if b && b1 {
			return firstbigger(set1, set2)
		}
		return b
	}
	b = twoPair(set1)
	b1 = twoPair(set2)
	if b || b1 {
		if b && b1 {
			return firstbigger(set1, set2)
		}
		return b
	}
	b = onePair(set1)
	b1 = onePair(set2)
	if b || b1 {
		if b && b1 {
			return firstbigger(set1, set2)
		}
		return b
	}
	/*
		h1 := higercard(set1)
		h2 := higercard(set2)
		if h1 == h2 {*/
	return firstbigger(set1, set2)
	/*}
	return h1 > h2*/
}

func puzzle1(name string) int64 {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var hands []hand

	for scanner.Scan() {
		sl := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(sl[1])
		hand := hand{cards: string(sl[0]), bit: n}
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return !compare(hands[i].cards, hands[j].cards)
	})
	var res int64

	for i, hand := range hands {
		res += int64((i + 1) * hand.bit)
	}

	return res
}
