package main

import (
	"fmt"
	"os"
)

func main() {
	r1 := puzzle1(os.Args[1])
	fmt.Printf("r1: %v\n", r1)
	if r1 == 288 {
		fmt.Printf("puzzle1: %v\n", puzzle1(os.Args[2]))
	}
	r2 := puzzle2(os.Args[1])
	if r2 == 71503 {
		fmt.Printf("puzzle2: %v\n", puzzle2(os.Args[2]))
	}
}
