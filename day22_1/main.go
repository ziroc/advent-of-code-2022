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
var curDir, X, Y int // 0 =right

func main() {
	fmt.Println("Starting Day 22 - 1")
	defer u.TimeTrack(time.Now(), "MAIN")
	file, err := os.Open("inputs/day22.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	mmap = make([][]int, 0, 8)

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

	for x := 0; x < len(mmap[0]); x++ {
		if mmap[0][x] == 1 {
			X = x
			break
		}
	}
	Y = 0

	// fmt.Println(path)
	// for y := 0; y < len(mmap); y++ {
	// 	fmt.Println(mmap[y])
	// }

	fmt.Println("SOLVE")
	fmt.Println("maxW ", maxW)
	fmt.Println("start: ", X, Y)
	runeP := []rune(path)
	for i := 0; i < len(runeP); i++ {
		n := runeP[i]
		if unicode.IsLetter(n) {
			turn(string(n))
		} else {
			if i < len(runeP)-1 && !unicode.IsLetter(runeP[i+1]) {
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

func move(s string) {
	fmt.Println("move ", s, curDir)
	moveX := getXAdd()
	moveY := getYAdd()

	num := u.Toint(s)
	for i := 0; i < num; i++ {
		if moveX != 0 {
			if X+moveX == maxW || X+moveX < 0 || mmap[Y][X+moveX] == 0 {
				fmt.Println("before wrap: ", X, Y)
				if !wrapX(moveX) {
					break
				} else {
					fmt.Println("wrapped X: ", X, Y)
					continue
				}
			} else if mmap[Y][X+moveX] == 2 {
				break
			} else {
				X += moveX
			}
		} else if moveY != 0 {
			if Y+moveY == len(mmap) || Y+moveY < 0 || mmap[Y+moveY][X] == 0 {
				if !wrapY(moveY) {
					break
				} else {
					fmt.Println("wrapped Y: ", X, Y)
					continue
				}
			} else if mmap[Y+moveY][X] == 2 {
				break
			} else {
				Y += moveY
			}
		}
		// fmt.Println("new: ", X, Y)
	}
}

func wrapY(moveY int) bool {
	if moveY == 1 {
		for i := 0; i < len(mmap); i++ {
			if mmap[i][X] == 1 {
				Y = i
				return true
			}
			if mmap[i][X] == 2 {
				return false
			}
		}
	}
	if moveY == -1 {
		for i := len(mmap) - 1; i >= 0; i-- {
			if mmap[i][X] == 1 {
				Y = i
				return true
			}
			if mmap[i][X] == 2 {
				return false
			}
		}
	}
	return false
}

func wrapX(moveX int) bool {
	if moveX == 1 {
		for i := 0; i < maxW; i++ {
			if mmap[Y][i] == 1 {
				X = i
				return true
			}
			if mmap[Y][i] == 2 {
				return false
			}
		}
	}
	if moveX == -1 {
		for i := maxW - 1; i >= 0; i-- {
			if mmap[Y][i] == 1 {
				X = i
				return true
			}
			if mmap[Y][i] == 2 {
				return false
			}
		}
	}
	return false
}

func getXAdd() int {
	switch curDir {
	case 0:
		return 1
	case 2:
		return -1
	}
	return 0
}

func getYAdd() int {
	switch curDir {
	case 1:
		return 1
	case 3:
		return -1
	}
	return 0
}

func turn(dir string) {
	fmt.Println("turn ", dir, " cur: ", curDir)
	switch dir {
	case "R":
		curDir++
	case "L":
		curDir--
	}
	curDir = mod (curDir, 4)
	// fmt.Println(" ", curDir)
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func readInput(line string, linecount int) {
	row := make([]int, len(line))
	mmap = append(mmap, row)
	for i := 0; i < len(line); i++ {
		if line[i:i+1] == "." {
			row[i] = 1 //ok
		}
		if line[i:i+1] == "#" {
			row[i] = 2 // solid
		}
	}
	// fmt.Println(row)
	maxW = u.Max(maxW, len(line))
}
