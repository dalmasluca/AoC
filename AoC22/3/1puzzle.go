package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func priority(s string) int {
    meta := len(s)/2
    prima := s[:meta]
    seconda := s[meta:]

    for _, runa := range prima {
        if strings.Contains(seconda, string(runa)) {
            if runa <= 'z' && runa >= 'a' {
                return int(runa - 'a' + 1)
            }else{
                return int(runa - 'A' + 'z' - 'a' + 2)
            }
        }
    }
    return 0
}

func main(){
    var total int
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        total += priority(scanner.Text())
    }
    fmt.Println(total)
}
