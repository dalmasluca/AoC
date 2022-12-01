package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main(){
    var numMax, numElfMax, ctElfo, numCal int
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        if scanner.Text() == "" {
            if numMax <= numCal {
                numMax = numCal
                numElfMax = ctElfo
                fmt.Println(numMax, numElfMax)
            }
            ctElfo++
            numCal = 0
        }else{
            i, _ := strconv.Atoi(scanner.Text())
            numCal += i
        }
    }
    fmt.Println(numElfMax)
}
