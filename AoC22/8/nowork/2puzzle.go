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
    left, right, top, bottom int
    scenicScore int
}

var baseTree treeMapping
var minTree treeMapping

func makeline(s string) []int {
	var line []int
	sl := strings.Split(s, "")
	for i := range sl {
		number, _ := strconv.Atoi(sl[i])
		line = append(line, number)
	}
	return line
}

func segVisible(number int, line []int, dir bool) bool {
	// if dir is true the function knows that is the right side
    if !dir {
        for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
            line[i], line[j] = line[j], line[i]
        }
    }

    baseTree.right++
    baseTree.left++
    for i := range line {
		if dir {
            baseTree.right++
        }else{
            baseTree.left++
        }
        if line[i] >= number {
			return false
		}
	}
	return true
}

func colVisible(num int, matrix [][]int, line int, col int) bool {
	top, bottom := true, true
	for i := range matrix[:line] {
        baseTree.top++
		if matrix[i][col] >= num {
			top = false
		}
	}
	for i := line + 1; i < len(matrix); i++ {
        baseTree.bottom++
		if matrix[i][col] >= num {
			bottom = false
		}
	}
	return top || bottom
}

func isVisible(number int, matrix [][]int, line int, coloumn int) int {
	if segVisible(number, matrix[line][:coloumn],false) || segVisible(number, matrix[line][coloumn+1:],true) || colVisible(number, matrix, line, coloumn) {
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
    baseTree = treeMapping{0,0,0,0,0,0,0}
    minTree = baseTree
    for i := 1; i < len(matrix)-1; i++ {
        for j := 1; j < len(matrix[i]) - 1; j++ {
            baseTree = treeMapping{i,j,0,0,0,0,0}
            Visible += isVisible(matrix[i][j],matrix,i,j)
	        baseTree.scenicScore = baseTree.top * baseTree.right * baseTree.left *baseTree.bottom
            fmt.Println(baseTree)
            if baseTree.scenicScore > minTree.scenicScore {
                minTree = baseTree
            }
        }
    }
    Visible += (len(matrix)*4 - 4)
    fmt.Println(Visible)
    fmt.Println(minTree)
}
