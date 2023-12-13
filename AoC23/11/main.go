package main

import (
	"fmt"
	"os"
)

func main() {
	test1 := puzzle1(os.Args[1])
	fmt.Printf("test1: %v\n", test1)
	if test1 == 374 {
		fmt.Printf("puzzle1: %v\n", puzzle1(os.Args[2]))
	}
	test2 := puzzle2(os.Args[1], 2)
	test3 := puzzle2(os.Args[1], 10)
	test4 := puzzle2(os.Args[1], 100)
	fmt.Printf("test2: %v\n", test2)
	if test2 == 374 && test3 == 1030 && test4 == 8410 {
		fmt.Printf("puzzle2: %v\n", puzzle2(os.Args[2], 1000000))
	}
}
