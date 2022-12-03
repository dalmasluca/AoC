package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func priority(s1,s2,s3 string) int {
    var priority int 
    presente := make(map[rune]bool)

    for _, runa := range s1 {
        if strings.Contains(s2, string(runa)) && strings.Contains(s3, string(runa)) {
            if runa <= 'z' && runa >= 'a' && !presente[runa] {
                priority += int(runa - 'a' + 1)
            }else if (!presente[runa]){
                priority += int(runa - 'A' + 'z' - 'a' + 2)
            }
            presente[runa] = true
        }
    }
    return priority
}

func main(){
    var total int
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        s1 := scanner.Text()
        scanner.Scan()
        s2 := scanner.Text()
        scanner.Scan()
        s3 := scanner.Text()

        total += priority(s1,s2,s3)
    }
    fmt.Println(total)
}
