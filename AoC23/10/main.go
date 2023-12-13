package main

import (
	"fmt"
	"os"
)

func main() {
	test1 := puzzle1(os.Args[1])
	test2 := puzzle1(os.Args[3])
	fmt.Printf("test1: %v\n", test1)
	fmt.Printf("test2: %v\n", test2)
	if test1 == 4 && test2 == 8 {
		fmt.Printf("puzzle1: %v\n", puzzle1(os.Args[2]))
	}
	test3 := puzzle2(os.Args[4])
	fmt.Printf("test3: %v\n", test3)
	if test3 == 10 {
		fmt.Printf("puzzle2: %v\n", puzzle2(os.Args[2]))
	}
}
