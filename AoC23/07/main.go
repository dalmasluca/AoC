package main

import (
	"fmt"
	"os"
)

func main() {
	test1 := puzzle1(os.Args[1])
	fmt.Printf("r1: %v\n", test1)
	if test1 == 6440 {
		fmt.Printf("puzzle1: %v\n", puzzle1(os.Args[2]))
	}
	test2 := puzzle2(os.Args[1])
	fmt.Printf("test2: %v\n", test2)
	if test2 == 5905 {
		fmt.Printf("puzzle2: %v\n", puzzle2(os.Args[2]))
	}
}
