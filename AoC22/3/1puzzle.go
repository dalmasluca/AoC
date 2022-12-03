package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func priority(s string) int {
    var priority int 
    meta := len(s)/2
    prima := s[:meta]
    seconda := s[meta:]
    presente := make(map[rune]bool)

    for _, runa := range prima {
        if strings.Contains(seconda, string(runa)) {
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
        total += priority(scanner.Text())
    }
    fmt.Println(total)
}
