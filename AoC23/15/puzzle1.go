package main

import (
	"bufio"
	"os"
	"strings"
)

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	res := 0

	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), ",")
		for _, e := range elements {
			p := 0
			for i := range e {
				p += int(e[i])
				p *= 17
				p %= 256
			}
			res += p
		}
	}

	return res
}
