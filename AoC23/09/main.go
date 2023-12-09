package main

import (
	"fmt"
	"os"
)

func main() {
	test1 := puzzle1(os.Args[1])
	fmt.Printf("test1: %v\n", test1)
	if test1 == 114 {
		fmt.Printf("puzzle1: %v\n", puzzle1(os.Args[2]))
	} else {
		return
	}
	test2 := puzzle2(os.Args[1])
	fmt.Printf("test2: %v\n", test2)
	if test2 == 2 {
		fmt.Printf("puzzle2: %v\n", puzzle2(os.Args[2]))
	}
}
