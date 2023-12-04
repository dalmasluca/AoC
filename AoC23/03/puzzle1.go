package main

import (
	"bufio"
	"os"
	"strconv"
)

func puzzle1() {
	in, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	prec := []rune(scanner.Text())

	for scanner.Scan() {
		succ := []rune(scanner.Text())
		for i, k := range succ {
			if n, err := strconv.Atoi(k); err != nil && k != '.' {
				//check update
				if i != 0 && i != len(succ) {
    	  	n, err := 

        }
			}
		}
		prec = succ
	}
}
