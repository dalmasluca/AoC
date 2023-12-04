package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func puzzle2(file string) int {
	in, _ := os.Open(file)
	scanner := bufio.NewScanner(in)
	copies := make([]int, 0)
	card := 0

	for scanner.Scan() {

		if card < len(copies) {
			copies[card]++
		} else {
			copies = append(copies, 1)
		}

		wins := strings.Fields(strings.Split(strings.Split(scanner.Text(), ": ")[1], "|")[0])
		games := strings.Fields(strings.Split(strings.Split(scanner.Text(), ": ")[1], "|")[1])

		count := 0

		for _, game := range games {
			for _, win := range wins {
				if game == win {
					count++
				}
			}
		}

		//fmt.Printf("copies: %v\n", copies)
		//fmt.Printf("La carta %v di cipie %v, ha vinto %v numeri.\n", card+1, copies[card], count)

		for i := card + 1; i <= card+count; i++ {
			if i < len(copies) {
				copies[i] += (1 * copies[card])
				//fmt.Printf("aggingo alla carta %v una copia\n", i+1)
			} else {
				copies = append(copies, (1 * copies[card]))
				//fmt.Printf("aggingo alla carta %v una copia\n", i+1)
			}
		}

		card++
	}

	//fmt.Printf("copies: %v\n", copies)

	tot := 0

	for _, k := range copies {
		tot += k
	}
	fmt.Printf("tot: %v\n", tot)
	return tot
}
