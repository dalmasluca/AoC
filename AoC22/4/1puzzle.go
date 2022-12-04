package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contain(s string) int {
    var sl []string
    var il []int
    sos := strings.Split(s, ",")

    for _, k :=range sos {
        sl = append(sl, strings.Split(k, "-")...)
    }
    
    for _, k := range sl {
        i,_ := strconv.Atoi(k)
        il = append(il, i)
    }

    fmt.Println(il)

    if (il[0]<=il[2] && il[1]>=il[3]) || (il[0]>=il[2] && il[1]<=il[3]){
        return 1
    }

    return 0
}

func main(){
    var totale int
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        totale += contain(scanner.Text())
    }
    fmt.Println(totale)
}
