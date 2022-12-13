package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
    posX, posY int
}

func main () {
    var mapArea [][]string
    var head, tail Knot
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        moves := scanner.Text()
        if moves[0] == 'R' {            
        }
        
    }
    
}

