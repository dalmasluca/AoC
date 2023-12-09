package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func puzzle2(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)

	var space string = " "
	_ = space
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
			scan[i] = slices.Insert(scan[i], 0, scan[i][0]-scan[i+1][0])
			//scan[i] = append(scan[i], scan[i][len(scan[i])-1]+scan[i+1][len(scan[i+1])-1])
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
		res += scan[0][0]
	}

	return res
}
