package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type slot struct {
	label string
	lens  int
}

type box struct {
	slots []slot
}

func In(b box, label string) int {
	for i, s := range b.slots {
		if s.label == label {
			return i
		}
	}
	return -1
}

func hash(s string) int {
	k := 0
	for i := range s {
		k += int(s[i])
		k *= 17
		k %= 256
	}
	return k
}

func puzzle2(name string) int64 {
	in, _ := os.Open(name)
	scanner := bufio.NewScanner(in)
	var boxes []box = make([]box, 256)
	var res int64

	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), ",")
		for _, e := range elements {
			k := 0
			if strings.Contains(e, "=") {
				k = hash(e[:len(e)-2])
				if i := In(boxes[k], e[:len(e)-2]); i != -1 {
					boxes[k].slots[i].lens, _ = strconv.Atoi(string(e[len(e)-1]))
				} else {
					n, _ := strconv.Atoi(string(e[len(e)-1]))
					boxes[k].slots = append(boxes[k].slots, slot{label: e[:len(e)-2], lens: n})
				}
			} else {
				k = hash(e[:len(e)-1])
				i := In(boxes[k], e[:len(e)-1])
				if i != -1 {
					boxes[k].slots = slices.Delete(boxes[k].slots, i, i+1)
				}
			}
		}
	}

	for b := range boxes {
		for l := range boxes[b].slots {
			res += int64((1 + b) * (l + 1) * (boxes[b].slots[l].lens))
		}
	}

	return res
}
