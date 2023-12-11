package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func allFinished(nodes []string) bool {
	for _, node := range nodes {
		if !strings.Contains(node, "Z") {
			return false
		}
	}
	return true
}

func gdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func puzzle2(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var gr graph = make(map[string][]string)

	scanner.Scan()
	sin := scanner.Text()
	sin = strings.ReplaceAll(strings.ReplaceAll(sin, "L", "0"), "R", "1")
	l := len(sin)
	var instructions []int = make([]int, l)
	var starts []string
	for i, runa := range sin {
		instructions[i], _ = strconv.Atoi(string(runa))
	}
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		node := strings.Fields(strings.Split(line, "=")[0])[0]
		adjs := strings.Fields(strings.Trim(strings.ReplaceAll(strings.Split(line, "=")[1], ",", ""), "( )"))
		if strings.Contains(node, "A") {
			starts = append(starts, node)
		}
		gr[node] = adjs
	}

	var times []int

	for i, start := range starts {
		times = append(times, 0)
		var time int = 0
		for !strings.Contains(start, "Z") {
			start = gr[start][instructions[time%l]]
			time++
		}
		times[i] = time
	}

	var g int = 1
	var mcm int

	for _, time := range times {
		a := gdc(g, time)
		mcm = (g * time) / a
		g = mcm
	}

	return mcm
}
