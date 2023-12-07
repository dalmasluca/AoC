package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var seeds []int
	var mapping []int = make([]int, 3)
	var check []bool
	scanner.Scan()

	for _, k := range strings.Split(strings.Split(scanner.Text(), ":")[1], " ") {
		n, err := strconv.Atoi(k)
		if err == nil {
			seeds = append(seeds, n)
		}
	}
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ":") {
			check = make([]bool, len(seeds))
		} else {
			for i, k := range strings.Split(scanner.Text(), " ") {
				mapping[i], _ = strconv.Atoi(k)
			}
			for i := range seeds {
				if seeds[i] >= mapping[1] && seeds[i] < mapping[1]+mapping[2] && !check[i] {
					seeds[i] = seeds[i] - mapping[1] + mapping[0]
					check[i] = true
				}
			}
		}
	}

	return slices.Min(seeds)
}
