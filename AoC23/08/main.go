package main

import (
	"fmt"
	"os"
)

type graph = map[string][]string

func main() {
	test1 := puzzle1(os.Args[1])
	test2 := puzzle1(os.Args[2])
	fmt.Printf("test1: %v\n", test1)
	fmt.Printf("test2: %v\n", test2)
	if test1 == 2 && test2 == 6 {
		fmt.Printf("puzzle1: %v\n", puzzle1(os.Args[3]))
	}
	test3 := puzzle2(os.Args[4])
	fmt.Printf("test3: %v\n", test3)
	if test3 == 6 {
		fmt.Printf("puzzle2: %v\n", puzzle2(os.Args[5]))
	}
}
