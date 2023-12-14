package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Equal(dst [][]string, emt [][]string) bool {
	for i := range dst {
		if !slices.Equal(dst[i], emt[i]) {
			return false
		}
	}
	return true
}

func In(all [][][]string, one [][]string) int {
	for i, k := range all {
		if Equal(k, one) {
			return i
		}
	}
	return -1
}

func puzzle2(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var m [][]string

	for scanner.Scan() {
		m = append(m, strings.Split(scanner.Text(), ""))
	}

	var history [][][]string
	history = append(history, m)
	var i int
	var l int
	for i = 0; i < 1000000000; i++ {
		m = rollRocks(rotate(rollRocks(rotate((rollRocks(rotate(rollRocks(rotate(m)))))))))
		if e := In(history, m); e != -1 {
			fmt.Printf("found it! after %v cycles!\n", i)
			fmt.Printf("e: %v\n", e)
			history = history[e:]
			l = e
			break
		}
		history = append(history, m)
	}
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%v", m[i][j])
		}
		fmt.Println()
	}
	fmt.Printf("history: %v\n", history)
	res := 0
	pos := (1000000000 - l) % len(history)
	afert := history[pos]

	for i := range afert {
		for j := range afert[i] {
			fmt.Printf("%v", afert[i][j])
			if afert[i][j] == "O" {
				res += (len(m) - i)
			}
		}
		fmt.Println()
	}
	return res
}
