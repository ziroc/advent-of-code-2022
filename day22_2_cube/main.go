package main

import (
	u "advent22/utils"
	"bufio"
	"fmt"
	"os"
	"time"
	"unicode"
)

var mmap [][]int
var path string
var maxW = 0
var sideSize = 50
var curDir, Y int // 0 =right
var X = -1
var portals []PortalPair
var runTests = false
var moveCount = 0

type PortalPair struct { // `sideSize` long
	p1, p2 Portal
	rotate string // L/R
	flip   bool
}
type Portal struct {
	topleftX, topleftY int
	isHorizontal       bool
}

func main() {
	fmt.Println("Starting Day 22 - 2")
	defer u.TimeTrack(time.Now(), "MAIN")
	file, err := os.Open("inputs/day22.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	mmap = make([][]int, 0, 8)

	p1to6 := PortalPair{Portal{50, 0, true}, Portal{0, 150, false}, "R", false}
	p1to5 := PortalPair{Portal{50, 0, false}, Portal{0, 100, false}, "RR", true}
	p2to6 := PortalPair{Portal{100, 0, true}, Portal{0, 199, true}, "", false}
	p2to3 := PortalPair{Portal{100, 49, true}, Portal{99, 50, false}, "R", false}
	p2to4 := PortalPair{Portal{149, 0, false}, Portal{99, 100, false}, "RR", true}
	p3to5 := PortalPair{Portal{50, 50, false}, Portal{0, 100, true}, "L", false}
	p4to6 := PortalPair{Portal{50, 149, true}, Portal{49, 150, false}, "R", false}
	portals = []PortalPair{p1to6, p1to5, p2to6, p2to3, p2to4, p3to5, p4to6}

	linecount := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			fileScanner.Scan()
			path = fileScanner.Text()
			break
		}
		readInput(line, linecount)
		linecount++
	}

	for y := 0; y < len(mmap); y++ {
		row := mmap[y]
		if len(row) < maxW {
			row = append(row, make([]int, maxW-len(row))...)
			mmap[y] = row
		}
	}

	if runTests {
		test()
		return
	}

	fmt.Println("start: ", X, Y)
	runeP := []rune(path)
	for i := 0; i < len(runeP); i++ {
		moveCount++
		n := runeP[i]
		if unicode.IsLetter(n) {
			turn(string(n), false)
		} else {
			if i < len(runeP)-1 && !unicode.IsLetter(runeP[i+1]) { //double digit move count
				move(string(runeP[i : i+2]))
				i++
			} else {
				move(string(runeP[i : i+1]))
			}
		}
	}
	fmt.Println("final: ", X, Y)
	fmt.Println("answ: ", 1000*(Y+1)+4*(X+1)+curDir)
}

func move(moveCount string) {
	for i := 0; i < u.Toint(moveCount); i++ {
		moveX, moveY := getNewXY()
		if moveX != 0 { // moving horizontaly
			if X+moveX == maxW || X+moveX < 0 || mmap[Y][X+moveX] == 0 {
				if !wrapEdge(false) {
					break
				} else {
					continue
				}
			}
		} else { // moving vertically
			if Y+moveY == len(mmap) || Y+moveY < 0 || mmap[Y+moveY][X] == 0 {
				if !wrapEdge(true) {
					break
				} else {
					continue
				}
			}
		}
		if mmap[Y+moveY][X+moveX] == 2 {
			break
		}
		X += moveX
		Y += moveY
	}
}

func wrapEdge(horizontal bool) bool {
	pp, p := findPortal(horizontal)
	var newX, newY int
	if !p.isHorizontal { // vertical
		newX = p.topleftX
		if pp.flip {
			newY = p.topleftY + sideSize - 1 - Y%sideSize
		} else {
			newY = p.topleftY + X%sideSize
		}
	} else { // horizontal
		newY = p.topleftY
		if pp.p1.isHorizontal && pp.p2.isHorizontal {
			newX = p.topleftX + X%sideSize
		} else if pp.flip {
			newX = p.topleftX + sideSize - X%sideSize
		} else {
			newX = p.topleftX + Y%sideSize
		}
	}

	if mmap[newY][newX] == 1 {
		X = newX
		Y = newY
		for i := 0; i < len(pp.rotate); i++ {
			if p == pp.p1 {
				turn(pp.rotate[i:i+1], true)
			} else {
				turn(pp.rotate[i:i+1], false)
			}
		}
		return true
	}
	if mmap[newY][newX] == 2 {
		return false
	}
	return false
}

func findPortal(horizontal bool) (pr PortalPair, goingTo Portal) {
	for _, p := range portals {
		if horizontal && horizontal == p.p1.isHorizontal {
			if X >= p.p1.topleftX && X <= p.p1.topleftX+sideSize-1 && Y == p.p1.topleftY {
				return p, p.p2
			}
		}
		if horizontal && horizontal == p.p2.isHorizontal {
			if X >= p.p2.topleftX && X <= p.p2.topleftX+sideSize-1 && Y == p.p2.topleftY {
				return p, p.p1
			}
		}
		if !horizontal && horizontal == p.p1.isHorizontal {
			if X == p.p1.topleftX && Y >= p.p1.topleftY && Y <= p.p1.topleftY+sideSize-1 {
				return p, p.p2
			}
		}
		if !horizontal && horizontal == p.p2.isHorizontal {
			if X == p.p2.topleftX && Y >= p.p2.topleftY && Y <= p.p2.topleftY+sideSize-1 {
				return p, p.p1
			}
		}
	}
	panic("couldnt find portal " + fmt.Sprint(X) + " " + fmt.Sprint(Y) + " looking for " + fmt.Sprint(horizontal))
}

var dirs map[int][]int = map[int][]int{0: {1, 0}, 1: {0, 1}, 2: {-1, 0}, 3: {0, -1}}

func getNewXY() (xx, yy int) {
	return dirs[curDir][0], dirs[curDir][1]
}

func turn(dir string, reverse bool) {
	switch dir {
	case "R":
		curDir++
	case "L":
		curDir--
	}
	if reverse { // turn actually in the opposite of the new direction
		curDir += 2
	}
	curDir = mod(curDir, 4)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func readInput(line string, linecount int) {
	row := make([]int, len(line))
	mmap = append(mmap, row)
	for i := 0; i < len(line); i++ {
		if line[i:i+1] == "." {
			row[i] = 1 //ok
			if linecount == 0 && X == -1 {
				X = i
			}
		}
		if line[i:i+1] == "#" {
			row[i] = 2 // solid
		}
	}
	maxW = u.Max(maxW, len(line))
}
