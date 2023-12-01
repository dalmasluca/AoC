package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func puzzle1() {
	scanner := bufio.NewScanner(os.Stdin)
	var somma int
	for scanner.Scan() {
		var n1, n2 int
		line := scanner.Text()
		for _, runa := range line {
			n, err := strconv.Atoi(string(runa))
			if err == nil {
				n1 = n
				break
			}
		}
		for _, runa := range Reverse(line) {
			n, err := strconv.Atoi(string(runa))
			if err == nil {
				n2 = n
				break
			}
		}
		somma += n1*10 + n2
	}
	fmt.Printf("somma: %v\n", somma)
}
