package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var gr graph = make(map[string][]string)

	scanner.Scan()
	sin := scanner.Text()
	sin = strings.ReplaceAll(strings.ReplaceAll(sin, "L", "0"), "R", "1")
	l := len(sin)
	var instructions []int = make([]int, l)
	for i, runa := range sin {
		instructions[i], _ = strconv.Atoi(string(runa))
	}
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		node := strings.Fields(strings.Split(line, "=")[0])[0]
		adjs := strings.Fields(strings.Trim(strings.ReplaceAll(strings.Split(line, "=")[1], ",", ""), "( )"))
		gr[node] = adjs
	}

	n := "AAA"
	res := 0

	for n != "ZZZ" {
		n = gr[n][instructions[res%l]]
		res++
	}

	return res
}
