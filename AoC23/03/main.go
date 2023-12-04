package main

import (
	"fmt"
	"os"
)

func main() {
	if puzzle1(os.Args[1]) == 4361 {
		fmt.Printf("puzzle1: %v\n", puzzle1(os.Args[2]))
	}
	p2 := puzzle2(os.Args[1])
	fmt.Printf("p2: %v\n", p2)
	if p2 == 467835 {
		fmt.Printf("puzzle2: %v\n", puzzle2(os.Args[2]))
	}
}
