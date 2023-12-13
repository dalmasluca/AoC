package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var matrix [][]string
	var start []int
	var next []int = make([]int, 2)
	var prec []int = make([]int, 2)

	for i := 0; scanner.Scan(); i++ {
		if strings.Contains(scanner.Text(), "S") {
			start = []int{i, strings.Index(scanner.Text(), "S")}
		}
		matrix = append(matrix, strings.Split(scanner.Text(), ""))
	}
	copy(next, start)
	copy(prec, start)
	res := 1
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
		res++
	}

	return res / 2
}
