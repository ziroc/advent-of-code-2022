package main

import (
	u "advent22/utils"
	"bufio"
	"fmt"
	i "image"
	"os"
	"time"
)

var emap [][]int
var elves map[i.Point]*i.Point = make(map[i.Point]*i.Point)
var example bool = false
var maxW = 74
var bufferLength = 60
var file string = "inputs/day23.txt"
var curPropDir = 0 // 0=N,1=S, 2=W, 3=E
var NO = i.Point{-1, -1}
var proposals map[i.Point]i.Point

func main() {
	fmt.Println("Starting Day 22 - 2")
	if example {
		file = "inputs/day23ex.txt"
		maxW = 7
	}
	defer u.TimeTrack(time.Now(), "MAIN")
	file, err := os.Open(file)
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	emap = make([][]int, 0, bufferLength)
	for i := 0; i < bufferLength; i++ {
		emap = append(emap, make([]int, 2*bufferLength+maxW)) //add buffer top
	}

	linecount := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		readInput(line, linecount)
		linecount++
	}

	for i := 0; i < bufferLength; i++ {
		emap = append(emap, make([]int, 2*bufferLength+maxW)) //add buffer bottom
	}

	cc:=0
	var moved bool
	for r := 0; r < 10000; r++ {
		cc++
		moved = false
		proposals = make(map[i.Point]i.Point)
		for _, v := range elves {
			findProp(*v)
		}
		for prop, elf := range proposals {
			if elf == NO {
				continue
			}
			moved = true
			emap[elf.Y][elf.X] = 0
			emap[prop.Y][prop.X] = 1
			e1 := elves[elf]
			delete(elves, elf)
			e1.X = prop.X
			e1.Y = prop.Y
			elves[*e1] = e1
		}
		if !moved {
			fmt.Println("GREAT, answer: ", cc)
			break
		}

		moveProp()
	}
}

func findProp(elfLoc i.Point) {
	empty := true
	for _, dir := range dirs {
		n := elfLoc.Add(dir)
		if emap[n.Y][n.X] != 0 {
			empty = false
			break
		}
	}
	if empty {
		return
	}

	cur := curPropDir
	for j := 0; j < 4; j++ {
		ok := true
		for _, dd := range cardinals[cur] {
			n := elfLoc.Add(dirs[dd])
			if emap[n.Y][n.X] != 0 {
				ok = false
				break
			}
		}
		if ok {
			nd := dirs[cardinals[cur][0]]
			nn := elfLoc.Add(nd)
			_, ok := proposals[nn]
			if ok {
				proposals[nn] = NO
				return
			}
			proposals[nn] = elfLoc
			break
		}
		cur = nextProp(cur)
	}
}

var dirs map[string]i.Point = map[string]i.Point{"n": {0, -1}, "ne": {1, -1}, "e": {1, 0},
	"se": {1, 1}, "s": {0, 1}, "sw": {-1, 1}, "w": {-1, 0}, "nw": {-1, -1}}

var cardinals map[int][]string = map[int][]string{
	0: {"n", "nw", "ne"},
	1: {"s", "se", "sw"},
	2: {"w", "sw", "nw"},
	3: {"e", "ne", "se"},
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func moveProp() {
	curPropDir++
	curPropDir = mod(curPropDir, 4)
}

func nextProp(cur int) int {
	cur++
	return mod(cur, 4)
}

func readInput(line string, linecount int) {
	row := make([]int, len(line)+2*bufferLength)
	emap = append(emap, row)
	for l := 0; l < len(line); l++ {
		if line[l:l+1] == "." {
			row[l+bufferLength] = 0 // empty
		}
		if line[l:l+1] == "#" {
			e := i.Point{l + bufferLength, linecount + bufferLength}
			elves[e] = &e
			row[l+bufferLength] = 1 // elf
		}
	}
}
