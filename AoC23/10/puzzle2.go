package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func In(all [][]int, one []int) bool {
	for _, k := range all {
		if slices.Equal(k, one) {
			return true
		}
	}
	return false
}

func puzzle2(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var matrix [][]string
	var start []int
	var next []int = make([]int, 2)
	var prec []int = make([]int, 2)

	var steps [][]int

	for i := 0; scanner.Scan(); i++ {
		if strings.Contains(scanner.Text(), "S") {
			start = []int{i, strings.Index(scanner.Text(), "S")}
		}
		matrix = append(matrix, strings.Split(scanner.Text(), ""))
	}
	steps = append(steps, start)
	copy(next, start)
	copy(prec, start)
	if strings.ContainsAny("|7F", matrix[max(next[0]-1, 0)][next[1]]) {
		next = []int{next[0] - 1, next[1]}
	} else if strings.ContainsAny("-7J", matrix[next[0]][next[1]+1]) {
		next = []int{next[0], next[1] + 1}
	} else if strings.ContainsAny("|JL", matrix[next[0]+1][next[1]]) {
		next = []int{next[0] + 1, next[1]}
	} else if strings.ContainsAny("-FL", matrix[next[0]][next[1]-1]) {
		next = []int{next[0], next[1] - 1}
	}

	for matrix[next[0]][next[1]] != "S" {
		steps = append(steps, []int{next[0], next[1]})
		switch matrix[next[0]][next[1]] {
		case "|":
			if prec[0] > next[0] {
				prec[0] = next[0]
				next[0]--
			} else {
				prec[0] = next[0]
				next[0]++
			}
		case "-":
			if prec[1] > next[1] {
				prec[1] = next[1]
				next[1]--
			} else {
				prec[1] = next[1]
				next[1]++
			}
		case "F":
			if prec[0] > next[0] {
				prec[0] = next[0]
				next[1]++
			} else {
				prec[1] = next[1]
				next[0]++
			}
		case "J":
			if prec[0] < next[0] {
				prec[0] = next[0]
				next[1]--
			} else {
				prec[1] = next[1]
				next[0]--
			}
		case "L":
			if prec[0] < next[0] {
				prec[0] = next[0]
				next[1]++
			} else {
				prec[1] = next[1]
				next[0]--
			}
		case "7":
			if prec[0] > next[0] {
				prec[0] = next[0]
				next[1]--
			} else {
				prec[1] = next[1]
				next[0]++
			}
		case ".":
			fmt.Println("Houston abbiamo un problema")
			break
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			var step []int = []int{i, j}
			if In(steps, step) {
			} else {
				matrix[i][j] = "."
			}
		}
	}
	res := 0

	for i := range matrix {
		for j := range matrix[i] {
			count := 0
			count2 := 0
			for k := 0; k < j; k++ {
				if matrix[i][j] == "." && strings.ContainsAny("|JL", matrix[i][k]) {
					count++
				}
			}
			for k := 0; k < i; k++ {
				if matrix[i][j] == "." && strings.ContainsAny("-FL", matrix[k][j]) {
					count2++
				}
			}
			if count%2 == 1 && count2%2 == 1 {
				res++
				matrix[i][j] = " "
			}
		}
	}

	return res
}
