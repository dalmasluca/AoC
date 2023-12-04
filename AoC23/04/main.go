package main

import (
	"os"
)

func main() {
	if puzzle1(os.Args[1]) == 13 {
		puzzle1(os.Args[2])
	}
	if puzzle2(os.Args[1]) == 30 {
		puzzle2(os.Args[2])
	}
	// puzzle2()
}
