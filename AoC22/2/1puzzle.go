package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func scores(enemy, me rune) int {
	return int(math.Abs(float64((int(enemy-'A'-me+'X')-1%3)*3))) + int(me-'Z') + 1
}

func main() {
	var totalScore int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sliceofString := strings.Split(scanner.Text(), " ")
		totalScore += scores(rune(sliceofString[0][0]), rune(sliceofString[1][0]))
	}
	fmt.Println(totalScore)
}
