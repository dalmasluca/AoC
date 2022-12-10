package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dir struct {
	name   string
	father *Dir
	child  map[string]*Dir
	file   []int
}

type wSlice struct {
    sl []int
}

type wDir struct {
    dir *Dir
}

type Filesystem struct {
	start *Dir
	curr  *Dir
}

func changedirectory(name string, dir *Dir) *Dir {
	if name == ".." {
		return dir.father
	}
	return dir.child[name]
}

func newfilesystem(name string, FS *Filesystem) {
	dir := new(Dir)
	dir.name = name
	dir.child = make(map[string]*Dir)
	FS.start = dir
	FS.curr = dir
}

func mkadir(dir *Dir, name string, father *Dir) {
	dir.name = name
	dir.child = make(map[string]*Dir)
	dir.father = father
}

func analyzeline(line []string, FS *Filesystem) {
	if line[0] == "$" {
		if line[1] == "cd" {
			if FS.start == nil {
				newfilesystem(line[2], FS)
			} else {
				FS.curr = changedirectory(line[2], FS.curr)
			}
		}
	} else {
		number, err := strconv.Atoi(line[0])
		if err == nil {
			FS.curr.file = append(FS.curr.file, number)
		} else {
			FS.curr.child[line[1]] = new(Dir)
			mkadir(FS.curr.child[line[1]], line[1], FS.curr)
		}
	}
}

func printFS(dir *Dir, spazio int) {
	for i := 0; i < spazio; i++ {
		fmt.Print("\t")
	}
	fmt.Print("- ", dir.name, "( dir )\n")
	for k := range dir.child {
		printFS(dir.child[k], spazio+1)
	}
	for i := range dir.file {
		fmt.Print(strings.Repeat("\t", spazio+1))
		fmt.Print("- ", dir.file[i], " (file size)\n")
	}
}

func totalsize(dir *Dir) int {
	var filesize int
	for i := range dir.file {
		filesize += dir.file[i]
	}
	for k := range dir.child {
		filesize += totalsize(dir.child[k])
	}
	return filesize
}

func allsize(wsl *wSlice, dir *Dir ) {
    wsl.sl = append(wsl.sl, totalsize(dir))
    for k := range dir.child {
        allsize(wsl, dir.child[k])
    }
}

func main() {
	spaceneeded := 30000000
	machinespace := 70000000
	FS := new(Filesystem)
    wsl := new(wSlice)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		analyzeline(line, FS)
	}
	printFS(FS.start, 0)
    spaziodaliberare := spaceneeded - (machinespace - totalsize(FS.start))
    allsize(wsl, FS.start)
    fmt.Println(wsl.sl)
    sort.Ints(wsl.sl)
    fmt.Println(wsl.sl)
    for i := range wsl.sl {
        if wsl.sl[i] > spaziodaliberare {
            fmt.Println(wsl.sl[i])
            break
        }
    }
    fmt.Println(spaziodaliberare)
}
