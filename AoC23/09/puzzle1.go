package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func allzeros(s []int) bool {
	for i := range s {
		if s[i] != 0 {
			return false
		}
	}
	return true
}

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)

	var space string = " "
	var scan [][]int = make([][]int, 1)
	res := 0

	for scanner.Scan() {
		scan = make([][]int, 1)
		for _, k := range strings.Fields(scanner.Text()) {
			n, _ := strconv.Atoi(k)
			scan[0] = append(scan[0], n)
		}
		for !allzeros(scan[len(scan)-1]) {
			scan = append(scan, make([]int, len(scan[len(scan)-1])-1))
			for i := 1; i < len(scan[len(scan)-2]); i++ {
				scan[len(scan)-1][i-1] = scan[len(scan)-2][i] - scan[len(scan)-2][i-1]
			}
		}
		for i := len(scan) - 2; i >= 0; i-- {
			scan[i] = append(scan[i], scan[i][len(scan[i])-1]+scan[i+1][len(scan[i+1])-1])
		}
		/*for i, k := range scan {
			fmt.Printf(strings.Repeat(space, i*3))
			for _, v := range k {
				fmt.Printf("%v %s", v, strings.Repeat(space, len(scan)-1))
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
		fmt.Printf("\n")
		fmt.Printf("\n")*/
		res += scan[0][len(scan[0])-1]
	}

	return res
}
