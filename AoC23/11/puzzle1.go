package main

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strings"
)

func puzzle1(name string) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var sky [][]string
	var galaxies [][]int

	for scanner.Scan() {
		line := scanner.Text()
		sky = append(sky, strings.Split(line, ""))
		if !strings.Contains(line, "#") {
			sky = append(sky, strings.Split(line, ""))
		}
	}
	k := 0
	for k < len(sky[0]) {
		yes := true
		for i := range sky {
			if sky[i][k] == "#" {
				yes = false
				break
			}
		}
		if yes {
			for i := range sky {
				sky[i] = slices.Insert(sky[i], k, ".")
			}
			k++
		}
		k++
	}

	for i := range sky {
		for j := range sky[i] {
			if sky[i][j] == "#" {
				galaxies = append(galaxies, []int{i, j})
			}
			//			fmt.Printf("%v", sky[i][j])
		}
		//		fmt.Printf("\n")
	}
	res := 0
	for galaxy := range galaxies {
		for othergalaxy := galaxy + 1; othergalaxy < len(galaxies); othergalaxy++ {
			res += int(math.Abs(float64(galaxies[galaxy][0]-galaxies[othergalaxy][0])) + math.Abs(float64(galaxies[galaxy][1]-galaxies[othergalaxy][1])))
		}
	}
	return res
}
