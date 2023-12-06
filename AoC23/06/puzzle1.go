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
	scanner.Scan()
	times := make([]int, 0)
	distances := make([]int, 0)
	for _, k := range strings.Fields(strings.Split(scanner.Text(), ":")[1]) {
		n, _ := strconv.Atoi(k)
		times = append(times, n)
	}
	scanner.Scan()
	for _, k := range strings.Fields(strings.Split(scanner.Text(), ":")[1]) {
		n, _ := strconv.Atoi(k)
		distances = append(distances, n)
	}
	res := 1
	for i := range times {
		count := 0
		for j := 0; j < times[i]-1; j++ {
			if (times[i]-j)*j > distances[i] {
				count++
			}
		}
		res *= count
	}

	return res
}
