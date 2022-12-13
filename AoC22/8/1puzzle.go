package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type treeMapping struct {
    //position
    x,y int
    //view
    top, left, right, bottom int
}

func makeline(s string) []int {
	var line []int
	sl := strings.Split(s, "")
	for i := range sl {
		number, _ := strconv.Atoi(sl[i])
		line = append(line, number)
	}
	return line
}

func segVisible(number int, line []int) bool {
	for i := range line {
		if line[i] >= number {
			return false
		}
	}
	return true
}

func colVisible(num int, matrix [][]int, line int, col int) bool {
	top, bottom := true, true
	for i := range matrix[:line] {
		if matrix[i][col] >= num {
			top = false
		}
	}
	for i := line + 1; i < len(matrix); i++ {
		if matrix[i][col] >= num {
			bottom = false
		}
	}
	return top || bottom
}

func isVisible(number int, matrix [][]int, line int, coloumn int) int {
	if segVisible(number, matrix[line][:coloumn]) || segVisible(number, matrix[line][coloumn+1:]) || colVisible(number, matrix, line, coloumn) {
		return 1
	}
	return 0
}

func main() {
	var matrix [][]int
    var Visible int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		matrix = append(matrix, makeline(scanner.Text()))
	}

    for i:= range matrix {
        for j := range matrix[i] {
            fmt.Print(matrix[i][j], " ")
        }
        fmt.Println()
    }
    fmt.Println()
	for i := 1; i < len(matrix)-1; i++ {
        for j := 1; j < len(matrix[i]) - 1; j++ {
            Visible += isVisible(matrix[i][j],matrix,i,j)
	    }
	}
    Visible += (len(matrix)*4 - 4)
    fmt.Println(Visible)
}
