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
    var reg int = 2
    var somma int
    var conta int
    var check int = 20
    var image string
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")
        for _, cmd := range s {
            conta++
            if (conta%40 >= reg - 1) && (conta%40 <= reg + 1) {
                image += "#"
            }else{
                image += "."
            }
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
    for i := range image{
        if i%40 == 0 {
            fmt.Println()
        }
        fmt.Print(string(image[i]))
    }
    fmt.Println()
    fmt.Println(somma)
}
