package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var all []any

var pmap [][]int

var paircount int
var linecount int
var total, minX, maxX, maxY int

func main() {
	fmt.Println("Starting Day 14 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day14.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	all = make([]any, 0, 300)
	// var lines [2]string
	minX = 1000
	pmap = make([][]int, 162)

	for fileScanner.Scan() {
		linecount++
		readInput(fileScanner.Text())
	}

	// drawScreen()
	sandCount := 0
	for {
		sandCount++
		stop := doWork()

		if stop || sandCount >10000 {
			break
		}
	}
	// fmt.Println(pmap[501][32])
	drawScreen()
	fmt.Println("answer ", sandCount)
	// fmt.Println("part 1 answer: ", total)

}

func doWork() bool {
	sandPos := Point{500, 0}

	for { 
		res := moveY(&sandPos)

		switch res  {
			case 1: {
				// fmt.Println("moved Y ", sandPos)
				continue
			}
			case -2: return true
		}
		res = moveX(&sandPos)
		switch res  {
			case -1:  {
				pmap[sandPos.Y][sandPos.X] = 2
				return false
			}
			case -2: return true
		}

	}
}

func moveX(sandPos *Point) int{
	if sandPos.Y == 161 {
		return -2
	}
	row := pmap[sandPos.Y+1]
	if row == nil {
		row = make([]int, 579)
		pmap[sandPos.Y+1] = row
	}
	left := sandPos.X-1
	right :=sandPos.X+1

	if left >= 0 && row[left] ==0 {
		sandPos.Y +=1 
		sandPos.X = left
		return 1
	}
	if right <= 578 && row[right] == 0 {
		sandPos.Y +=1 
		sandPos.X = right
		return 1
	}
	return -1
}

func moveY(sandPos *Point) int {
	if sandPos.Y == 161 {
		return -2
	}

	row := pmap[sandPos.Y+1]
	if row == nil {
		row = make([]int, 579)
		pmap[sandPos.Y+1] = row
	}
	if row[sandPos.X] == 0 {
		sandPos.Y +=1 
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
		// if p.X > maxX {
		// 	maxX = p.X
		// }
		// if p.Y > maxY {
		// 	maxY = p.Y
		// }
		// if p.X < minX {
		// 	minX = p.X
		// }
	}
}

func makeSolid(from, to Point) {
	// if from.Y == 117 || to.Y == 117 {
	// 	fmt.Println(from, " to ", to)
	// }
	if from.X == to.X {
		m := min(from.Y, to.Y)
		for i := 0; i <= abs(from.Y-to.Y); i++ {
			row := pmap[m+i]
			if row == nil {
				row = make([]int, 579)
				pmap[m+i] = row
			}
			pmap[m+i][from.X] = 1

		}
	}
	if from.Y == to.Y {
		m := min(from.X, to.X)
		for i := 0; i <= abs(from.X-to.X); i++ {
			row := pmap[from.Y]
			if row == nil {
				row = make([]int, 579)
				pmap[from.Y] = row
			}
			pmap[from.Y][m+i] = 1

		}
	}

}

func drawScreen() {
	for y := 0; y < 162; y++ {
		row := pmap[y]
		if row == nil {
			row = make([]int, 579)
			pmap[y] = row
		}
		fmt.Println()
		for x := 473; x < 579; x++ {
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

type Point struct {
	X int
	Y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
