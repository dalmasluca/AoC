package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func puzzle1(file string) int {
	in, _ := os.Open(file)
	scanner := bufio.NewScanner(in)
	tot := 0

	for scanner.Scan() {

		wins := strings.Fields(strings.Split(strings.Split(scanner.Text(), ": ")[1], "|")[0])
		games := strings.Fields(strings.Split(strings.Split(scanner.Text(), ": ")[1], "|")[1])

		score := 0.5

		for _, game := range games {
			for _, win := range wins {
				if game == win {
					score *= 2
				}
			}
		}

		tot += int(score)
	}
	fmt.Printf("tot: %v\n", tot)
	return tot
}
