package main

import (
	"bufio"
	"os"
)

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {

	}

	res := 0
	return res
}
