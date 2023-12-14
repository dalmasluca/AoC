package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func checkRows(m [][]string) int {
	for i := 0; i < len(m)-1; i++ {
		yes := true
		//fmt.Printf("i: %v\n", i)
		for j := 0; j < min(i+1, len(m)-i-1); j++ {
			//fmt.Printf("m[%v]: %v\n", i-j, m[i-j])
			//fmt.Printf("m[%v]: %v\n", i+j+1, m[i+j+1])
			if !slices.Equal(m[i-j], m[i+j+1]) {
				yes = false
				break
			}
		}
		if yes {
			return i + 1
		}
	}
	return 0
}

func transpost(m [][]string) [][]string {
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

func checkColumns(m [][]string) int {
	t := transpost(m)
	return checkRows(t)
}

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)

	var m [][]string

	res := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			r1 := checkRows(m)
			r2 := checkColumns(m)
			res += (r1 * 100)
			res += r2
			m = make([][]string, 0)
		} else {
			m = append(m, strings.Split(scanner.Text(), ""))
		}
	}

	return res
}
