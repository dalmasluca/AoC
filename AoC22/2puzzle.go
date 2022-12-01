package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func topisFull(a [3]int) int {
    for i := range a {
        if a[i] == 0 {
            return i
        }
    }
    return -1
}

func min(a [3]int) int {
    min := 0
    for i := range a {
        if a[min] >= a[i]{
            min = i
        }
    }
    return min
}

func main(){
    var ctElfo, numCal int
    var top3 [3]int
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        if scanner.Text() == "" {
            
            info, min := topisFull(top3), min(top3)

            if info != -1 { //se ce spazio nell'array allora inserisco nell'ultimo spazio il valore di calorie
                fmt.Println("l'array e ancora vuoto")
                top3[info] = numCal
                fmt.Println(top3)
            }else{ //altrimento cerco e sostituisco il valore minore
                if numCal > top3[min] {
                    top3[min] = numCal
                }
            }
            ctElfo++
            numCal = 0
        }else{
            i, _ := strconv.Atoi(scanner.Text())
            numCal += i
        }
    }
    fmt.Println(top3[0]+top3[1]+top3[2])
}
