package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

func puzzle2(name string, expensions int) int {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var sky [][]string
	var galaxies [][]int

	var rows []int
	var columns []int

	for scanner.Scan() {
		line := scanner.Text()
		sky = append(sky, strings.Split(line, ""))
		if !strings.Contains(line, "#") {
			rows = append(rows, len(sky)-1)
			// sky = append(sky, strings.Split(line, ""))
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
			columns = append(columns, k)
			/*for i := range sky {
				sky[i] = slices.Insert(sky[i], k, ".")
			}*/
			k++
		}
		k++
	}
	expensions--
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
			di := 0
			mi, ma := min(galaxies[galaxy][0], galaxies[othergalaxy][0]), max(galaxies[galaxy][0], galaxies[othergalaxy][0])
			for _, r := range rows {
				if r > mi && r < ma {
					di += expensions
				}
			}
			mi, ma = min(galaxies[galaxy][1], galaxies[othergalaxy][1]), max(galaxies[galaxy][1], galaxies[othergalaxy][1])
			for _, c := range columns {
				if c > mi && c < ma {
					di += expensions
				}
			}
			res += di
			res += int(math.Abs(float64(galaxies[galaxy][0]-galaxies[othergalaxy][0])) + math.Abs(float64(galaxies[galaxy][1]-galaxies[othergalaxy][1])))
		}
	}
	return res
}
