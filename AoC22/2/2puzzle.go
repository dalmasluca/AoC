package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scores(enemy, me rune) int {
    return int(me - 'X') * 3 + ((int(enemy - 'A' + (me - 'Y')) + 3) % 3) + 1
}

func main() {
	var totalScore int 
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	    sliceofString := strings.Split(scanner.Text(), " ")
        totalScore += scores( rune(sliceofString[0][0]), rune(sliceofString[1][0]))
    }
	fmt.Println(totalScore)
}
