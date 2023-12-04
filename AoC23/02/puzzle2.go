package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func puzzle2() {
	in, _ := os.Open(os.Args[1])
	defer in.Close()
	ss := bufio.NewScanner(in)
	game := 1
	somma := 0
	somma1 := 0

	for ss.Scan() {
		sets := strings.Split(strings.Split(ss.Text(), ":")[1], ";")
		ok := true
		var s map[string]int = make(map[string]int)
		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				color := strings.Split(ball, " ")
				n, _ := strconv.Atoi(color[1])
				if n > s[color[2]] {
					s[color[2]] = n
				}
			}
		}
		p := 1
		for _, v := range s {
			p *= v
		}
		if ok {
			somma += game
		}
		somma1 += p
		game += 1
	}
	fmt.Printf("somma1: %v\n", somma1)
}
