package main

import (
	"bufio"
	"os"
	"strings"
)

func rotate(m [][]string) [][]string {
	var t [][]string
	for j := range m[0] {
		col := make([]string, 0)
		for i := len(m) - 1; i >= 0; i-- {
			col = append(col, m[i][j])
		}
		t = append(t, col)
	}
	return t
}

func rollRocks(m [][]string) [][]string {
	for j := len(m[0]) - 2; j >= 0; j-- {
		for i := range m {
			if m[i][j] == "O" {
				var k int
				for k = j + 1; k < len(m[0]) && m[i][k] != "#" && m[i][k] != "O"; k++ {
				}
				m[i][j] = "."
				m[i][k-1] = "O"
			}
		}
	}
	return m
}

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var m [][]string

	for scanner.Scan() {
		m = append(m, strings.Split(scanner.Text(), ""))
	}

	m = rollRocks(rotate(m))
	res := 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] == "O" {
				res += (j + 1)
			}
		}
	}
	return res
}
