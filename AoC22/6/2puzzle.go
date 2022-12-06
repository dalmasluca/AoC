package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        for i:=0; i<len(scanner.Text())-13; i++ {
            sub :=  scanner.Text()[i:i+13]
            vero := true
            for j,k := range sub {
                sub2 := sub[:j] + sub[j+1:]
                if strings.Contains(sub2,string(k)) {
                    vero = false
                }
            }
            if vero {
                fmt.Print(i+14)
                break
            }
        }
    }
}
