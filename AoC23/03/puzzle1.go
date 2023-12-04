package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isnearSim(m [][]string, x, y int) bool {
	right := rune(m[x][min(y+1, len(m[0])-1)][0])
	left := rune(m[x][max(y-1, 0)][0])
	up := rune(m[max(x-1, 0)][y][0])
	dw := rune(m[min(x+1, len(m)-1)][y][0])
	uplft := rune(m[max(x-1, 0)][max(y-1, 0)][0])
	uprgt := rune(m[max(x-1, 0)][min(y+1, len(m[0])-1)][0])
	dwrgt := rune(m[min(x+1, len(m)-1)][min(y+1, len(m[0])-1)][0])
	dwlft := rune(m[min(x+1, len(m)-1)][max(y-1, 0)][0])

	if (left != '.' && !unicode.IsDigit(left)) ||
		(right != '.' && !unicode.IsDigit(right)) ||
		(up != '.' && !unicode.IsDigit(up)) ||
		(dw != '.' && !unicode.IsDigit(dw)) || (dwlft != '.' && !unicode.IsDigit(dwlft)) || (dwrgt != '.' && !unicode.IsDigit(dwrgt)) ||
		(uplft != '.' && !unicode.IsDigit(uplft)) || (uprgt != '.' && !unicode.IsDigit(uprgt)) {
		return true
	}
	return false
}

func puzzle1(name string) int {
	file, _ := os.ReadFile(name)
	lines := strings.Split(string(file), "\n")
	m := make([][]string, len(lines)-1)

	for i := range lines {
		if i != len(lines)-1 {
			m[i] = strings.Split(lines[i], "")
		}
	}

	buffer := ""
	valid := false

	tot := 0

	i, j := 0, 0
	for ; (i/len(m[0])) != len(m)-1 || j != len(m[0])-1; i, j = i+1, (j+1)%len(m[0]) {

		x := i / len(m)

		if unicode.IsDigit([]rune(m[x][j])[0]) {
			buffer += m[x][j]
			valid = valid || isnearSim(m, x, j)
		} else if buffer != "" && valid {
			n, _ := strconv.Atoi(buffer)
			tot += n
			buffer = ""
			valid = false
		} else {
			buffer = ""
			valid = false
		}

	}
	return tot
}
