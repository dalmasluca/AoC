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
            return 0 + 3
        case "Y":
            return 3 + 1
        case "Z":
            return 6 + 2
        }
    case "B":
        switch b{
        case "X":
            return 0 + 1
        case "Y":
            return 3 + 2
        case "Z":
            return 6 + 3
        }
    case "C":
        switch b{
        case "X":
            return 0 + 2
        case "Y":
            return 3 + 3
        case "Z":
            return 6 + 1
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
