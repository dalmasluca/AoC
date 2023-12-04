package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isValidGear(m [][]string, x, y int) (bool, [][]int) {
	right := rune(m[x][min(y+1, len(m[0])-1)][0])
	left := rune(m[x][max(y-1, 0)][0])
	up := rune(m[max(x-1, 0)][y][0])
	dw := rune(m[min(x+1, len(m)-1)][y][0])
	uplft := rune(m[max(x-1, 0)][max(y-1, 0)][0])
	uprgt := rune(m[max(x-1, 0)][min(y+1, len(m[0])-1)][0])
	dwrgt := rune(m[min(x+1, len(m)-1)][min(y+1, len(m[0])-1)][0])
	dwlft := rune(m[min(x+1, len(m)-1)][max(y-1, 0)][0])

	count := 0
	pos := make([][]int, 0)

	if !unicode.IsDigit(up) {
		if unicode.IsDigit(uplft) {
			count++
			pos = append(pos, []int{max(x-1, 0), max(y-1, 0)})
		}
		if unicode.IsDigit(uprgt) {
			count++
			pos = append(pos, []int{max(x-1, 0), min(y+1, len(m[0])-1)})
		}
	} else {
		count++
		pos = append(pos, []int{max(x-1, 0), y})
	}

	if unicode.IsDigit(left) {
		count++
		pos = append(pos, []int{x, max(y-1, 0)})
	}

	if unicode.IsDigit(right) {
		count++
		pos = append(pos, []int{x, min(y+1, len(m[0])-1)})
	}

	if !unicode.IsDigit(dw) {
		if unicode.IsDigit(dwlft) {
			count++
			pos = append(pos, []int{min(x+1, len(m)-1), max(y-1, 0)})
		}
		if unicode.IsDigit(dwrgt) {
			count++
			pos = append(pos, []int{min(x+1, len(m)-1), min(y+1, len(m[0])-1)})
		}
	} else {
		count++
		pos = append(pos, []int{min(x+1, len(m)-1), y})
	}
	return count == 2, pos
}

func detectNum(m [][]string, pos [][]int) (int, int) {
	buffer1 := m[pos[0][0]][pos[0][1]]
	//fmt.Printf("pos: %v\n", pos)
	x := pos[0][0]
	y := pos[0][1] - 1
	for y >= 0 && unicode.IsDigit(rune(m[x][y][0])) {
		buffer1 = m[x][y] + buffer1
		y--
	}
	y = pos[0][1] + 1
	for y < len(m[0]) && unicode.IsDigit(rune(m[x][y][0])) {
		buffer1 += m[x][y]
		y++
	}

	n1, _ := strconv.Atoi(buffer1)

	buffer2 := m[pos[1][0]][pos[1][1]]
	x = pos[1][0]
	y = pos[1][1] - 1
	for y >= 0 && unicode.IsDigit(rune(m[x][y][0])) {
		buffer2 = m[x][y] + buffer2
		y--
	}
	y = pos[1][1] + 1
	for y < len(m[0]) && unicode.IsDigit(rune(m[x][y][0])) {
		buffer2 += m[x][y]
		y++
	}

	n2, _ := strconv.Atoi(buffer2)

	return n1, n2
}

func puzzle2(name string) int {

	file, _ := os.ReadFile(name)
	lines := strings.Split(string(file), "\n")

	m := make([][]string, len(lines)-1)

	for i := range lines {
		if i != len(lines)-1 {
			m[i] = strings.Split(lines[i], "")
		}
	}
	//fmt.Printf("m: %v\n", m)
	tot := 0
	i, j := 0, 0

	for ; (i/len(m[0])) != len(m)-1 || j != len(m[0])-1; i, j = i+1, (j+1)%len(m[0]) {
		x := i / len(m)

		if v, p := isValidGear(m, x, j); v && m[x][j] == "*" {
			n1, n2 := detectNum(m, p)
			//		fmt.Printf("p: %v\n", p)
			//		fmt.Printf("%v\n", m[x])
			//		fmt.Printf("n1: %v\n", n1)
			//		fmt.Printf("n2: %v\n", n2)
			tot += (n1 * n2)
		}
	}
	return tot
}
