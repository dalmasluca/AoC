package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func puzzle2() {
	scanner := bufio.NewScanner(os.Stdin)
	var somma int

	var m = make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line: %v\n", line)

		firstIndexDigit := strings.IndexFunc(line, unicode.IsDigit)
		lastIndexDigit := strings.LastIndexFunc(line, unicode.IsDigit)

		fmt.Printf("firstIndexDigit: %v\n", firstIndexDigit)
		fmt.Printf("lastIndexDigit: %v\n", lastIndexDigit)

		firstIndexWord := len(line)
		lastIndexWord := 0

		var firstNumber, lastNumber int

		for number := range m {
			fiw := strings.Index(line, number)
			liw := strings.LastIndex(line, number)

			if fiw != -1 && fiw < firstIndexWord {
				firstIndexWord = fiw
				firstNumber = m[number]
			}
			if liw > lastIndexWord {
				lastIndexWord = liw
				lastNumber = m[number]
			}

		}

		if firstIndexDigit < firstIndexWord && firstIndexDigit != -1 {
			firstNumber, _ = strconv.Atoi(string(line[firstIndexDigit]))
		}

		if lastIndexDigit >= lastIndexWord {
			lastNumber, _ = strconv.Atoi(string(line[lastIndexDigit]))
		}

		fmt.Printf("firstNumber: %v\n", firstNumber)
		fmt.Printf("lastNumber: %v\n", lastNumber)

		somma += (firstNumber*10 + lastNumber)
	}
	fmt.Printf("somma: %v\n", somma)
}
