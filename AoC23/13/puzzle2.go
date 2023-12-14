package main

import (
	"bufio"
	"os"
	"strings"
)

func EqualbyOne(s1 []string, s2 []string) int {
	count := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}

func checkRows2(m [][]string) int {
	for i := 0; i < len(m)-1; i++ {
		//fmt.Printf("i: %v\n", i)
		count := 0
		for j := 0; j < min(i+1, len(m)-i-1); j++ {
			//fmt.Printf("m[%v]: %v\n", i-j, m[i-j])
			//fmt.Printf("m[%v]: %v\n", i+j+1, m[i+j+1])
			count += EqualbyOne(m[i-j], m[i+j+1])
		}
		if count == 1 {
			return i + 1
		}
	}
	return 0
}

func checkColumns2(m [][]string) int {
	return checkRows2(transpost(m))
}

func puzzle2(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)

	var m [][]string

	res := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			/*
			         for i := range m {
			   				for j := range m[i] {
			   					fmt.Printf("%v", m[i][j])
			   				}
			   				fmt.Printf("\n")
			   			}
			*/
			//fmt.Println()
			r1 := checkRows2(m)
			//fmt.Printf("checkRows(m): %v\n", r1)
			r2 := checkColumns2(m)
			res += (r1 * 100) + r2
			m = make([][]string, 0)
		} else {
			m = append(m, strings.Split(scanner.Text(), ""))
		}
	}

	return res
}
