package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var seeds []int
	var mapping [][]uint
	scanner.Scan()

	for _, k := range strings.Split(strings.Split(scanner.Text(), ":")[1], " ") {
		n, _ := strconv.Atoi(k)
		seeds = append(seeds, n)
	}

	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	count := 0
	for scanner.Text() != "" {
		mapping = append(mapping, make([]int, 3))
		for i, k := range strings.Split(scanner.Text(), " ") {
			mapping[count][i], _ = strconv.Atoi(k)
		}
		scanner.Scan()
		count++
	}

	fmt.Printf("mapping: %v\n", mapping)

	fmt.Printf("seeds: %v\n", seeds)

	return 0
}
