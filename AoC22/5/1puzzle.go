package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createCrates(scanner *bufio.Scanner) [][]string {
	// leggo la prima riga per sapere quando lungo devo creare la slice
	scanner.Scan()
	crates := make([][]string, (len(scanner.Text())/4)+1)
	s := scanner.Text()
	// popolo la prima riga della slice con il testo letto
	for i := 1; i < len(s); i += 4 {
		if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
			crates[i/4] = append(crates[i/4], string(s[i]))
		}
	}
	//finisco la lettura dal file
	for scanner.Scan() {
		// ciclo le singole righe del file
		for i := 1; i < len(scanner.Text()); i += 4 {
			if rune(scanner.Text()[i]) >= 'A' && rune(scanner.Text()[i]) <= 'Z' {
				crates[i/4] = append(crates[i/4], string(scanner.Text()[i]))
			}
		}
	}
	// raddrizzo le slice
	for k := range crates {
		for i, j := 0, len(crates[k])-1; i < j; i, j = i+1, j-1 {
			crates[k][i], crates[k][j] = crates[k][j], crates[k][i]
		}
	}
	return crates
}

func moveCrates(crates *[][]string, scanner *bufio.Scanner) {
	for scanner.Scan() {
		var sn []int
		s := scanner.Text()
		ss := strings.Split(s, " ")
		for _, stringa := range ss {
			number, err := strconv.Atoi(stringa)
			if err == nil {
				sn = append(sn, number)
			}
		}
		for i := 0; i < sn[0]; i++ {
			if len((*crates)[sn[1]-1]) > 0 {
				(*crates)[sn[2]-1] = append((*crates)[sn[2]-1], (*crates)[sn[1]-1][len((*crates)[sn[1]-1])-1])
				(*crates)[sn[1]-1] = (*crates)[sn[1]-1][:len((*crates)[sn[1]-1])-1]
			} else {
				(*crates)[sn[2]-1] = append((*crates)[sn[2]-1], (*crates)[sn[1]-1]...)
				(*crates)[sn[1]-1] = []string{}
			}
		}
	}
}

func main() {
	file, err := os.Open("crates.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	crates := createCrates(scanner)

	defer file.Close()

	file, err = os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner = bufio.NewScanner(file)
	moveCrates(&crates, scanner)
    defer file.Close()
	for i := 0; i < len(crates); i++ {
		fmt.Print(crates[i][len(crates[i])-1])
	}
}
