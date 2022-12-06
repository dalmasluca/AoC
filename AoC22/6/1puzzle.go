package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    s := ""
    for scanner.Scan() {
        for i, k := range scanner.Text() {
            if !strings.Contains(s,string(k)) {
                s += string(k)
            }else{
                s = string(k)
            }
            if len(s) == 4 {
                fmt.Println(i+1)
                break
            }
        }
    }
}
