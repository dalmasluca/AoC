package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func puzzle2(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	var time int
	var distance int
	time, _ = strconv.Atoi(strings.Join(strings.Fields(strings.Split(scanner.Text(), ":")[1]), ""))
	scanner.Scan()
	distance, _ = strconv.Atoi(strings.Join(strings.Fields(strings.Split(scanner.Text(), ":")[1]), ""))

	count := 0
	for j := 0; j < time-1; j++ {
		if (time-j)*j > distance {
			count++
		}
	}

	return count
}
