package main

import (
	"fmt"
	"strings"
    "bufio"
    "os"
)

func score(a, b string) int {
    switch a{
    case "A":
        switch b{
        case "X":
            return 1 + 3
        case "Y":
            return 2 + 6
        case "Z":
            return 3 + 0
        }
    case "B":
        switch b{
        case "X":
            return 1 + 0
        case "Y":
            return 2 + 3
        case "Z":
            return 3 + 6
        }
    case "C":
        switch b{
        case "X":
            return 1 + 6
        case "Y":
            return 2 + 0
        case "Z":
            return 3 + 3
        }
    }
    return 0
}

func main(){
    var totalScore int
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        game := strings.Split(scanner.Text(), " ")
        totalScore += score(game[0],game[1])
    }
    fmt.Println(totalScore)
}
