package main

import (
	"bufio"
	"fmt"
	"os"
	. "image"
	"strconv"
	"strings"
	"time"
)

var pmap [][]int

const MX = 700
var start = Point{500, 0}

func main() {
	fmt.Println("Starting Day 14 - 2")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day14.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	pmap = make([][]int, 164)

	for i:= range pmap {
		row := make([]int, MX)
		pmap[i] = row
	}

	row := make([]int, MX)
	for i := 0; i < len(row); i++ {
		row[i] = 1
	}
	pmap[163] = row

	for fileScanner.Scan() {
		readInput(fileScanner.Text())
	}

	sandCount := 0
	for {
		sandCount++
		done := doWork()

		if done || sandCount > 100000 {
			break
		}
	}
	drawScreen()
	fmt.Println("answer ", sandCount)
}

func doWork() bool {
	sandPos := start

	for {
		result := moveY(&sandPos)

		switch result {
		case 1:
			continue
		case -2:
			return true
		}
		result = moveX(&sandPos)
		if result == -1 {
			pmap[sandPos.Y][sandPos.X] = 2
			return sandPos == start // true: reached the start, stop the work
		}
	}
}

func moveX(sandPos *Point) int {
	row := pmap[sandPos.Y+1]
	left := sandPos.X - 1
	right := sandPos.X + 1

	if left >= 0 && row[left] == 0 {
		sandPos.Y += 1
		sandPos.X = left
		return 1
	}
	if right <= MX-1 && row[right] == 0 {
		sandPos.Y += 1
		sandPos.X = right
		return 1
	}
	return -1
}

func moveY(sandPos *Point) int {
	if pmap[sandPos.Y+1][sandPos.X] == 0 {
		sandPos.Y += 1
		return 1
	}
	return -1
}

func readInput(line string) {
	points := strings.Split(line, " -> ")
	var prevP Point
	for i := range points {
		xy := strings.Split(points[i], ",")
		p := Point{toint(xy[0]), toint(xy[1])}
		if i == 0 {
			prevP = p
			continue
		}
		makeSolid(prevP, p)
		prevP = p
	}
}

func makeSolid(from, to Point) {
	if from.X == to.X {
		m := min(from.Y, to.Y)
		for i := 0; i <= abs(from.Y-to.Y); i++ {
			row := pmap[m+i]
			row[from.X] = 1
		}
	}
	if from.Y == to.Y {
		m := min(from.X, to.X)
		for i := 0; i <= abs(from.X-to.X); i++ {
			row := pmap[from.Y]			
			row[m+i] = 1
		}
	}

}

func drawScreen() {
	for y := 0; y < 164; y++ {
		row := pmap[y]
		if row == nil {
			row = make([]int, MX)
			pmap[y] = row
		}
		fmt.Println()
		for x := 500 - 162; x < 644; x++ {
			fmt.Print(pmap[y][x])
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func toint(s string) int {
	value, err := strconv.Atoi(s)
	check(err)
	return value
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
