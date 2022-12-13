package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main () {
    var cycle []int
    var reg int = 1
    var somma int
    var conta int
    var check int = 20
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")
        for _, cmd := range s {
            conta++
            if conta == check {
                check += 40
                somma += conta * reg
            }
            number, err := strconv.Atoi(cmd)
            if err == nil {
                reg += number
            }
            cycle = append(cycle, reg)
        }
    }
    fmt.Println(somma)
}
