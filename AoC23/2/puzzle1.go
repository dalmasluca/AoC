package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func puzzle1() {
	in, _ := os.Open(os.Args[1])
	ss := bufio.NewScanner(in)
	game := 1
	somma := 0

	var m = map[string]int{"red": 12, "green": 13, "blue": 14}
	_ = m
	for ss.Scan() {
		sets := strings.Split(strings.Split(ss.Text(), ":")[1], ";")
		ok := true
		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				color := strings.Split(ball, " ")
				n, _ := strconv.Atoi(color[1])
				if n > m[color[2]] {
					ok = false
					break
				}
			}
		}
		if ok {
			somma += game
		}
		game += 1
	}
	fmt.Printf("somma: %v\n", somma)
}
