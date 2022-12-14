package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Empty struct{}

var head Point
var tails []Point

var visited map[Point]Empty

func main() {
	fmt.Println("Starting Day 9 - 2")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day9.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lineCount := 0
	visited = map[Point]Empty{}
	tails = make([]Point, 9)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			readInput(line, lineCount)
			lineCount++
			// if lineCount > 15 {
			// 	break
			// }

		} else {
			break
		}
	}
	fmt.Println("--------")
	fmt.Println(len(visited))
}

func readInput(line string, lineCount int) {

	split := strings.Split(line, " ")
	doWork(split[0], toint(split[1]))

}

func doWork(dir string, steps int) {
	fmt.Print()
	for i := 0; i < steps; i++ {
		move(dir, &head)
		// var lastdir = dir
		for i := range tails {
			if i == 0 {
				if notTouch(head, tails[i]) {
					keepUp(head, &tails[i])
				}
			} else {
				if notTouch(tails[i-1], tails[i]) {
					keepUp(tails[i-1], &tails[i])
				}
			}
		}
		visited[tails[8]] = Empty{}
	}
}

func move(dir string, p *Point) {
	switch dir {
	case "U":
		p.Y++
	case "D":
		p.Y--
	case "R":
		p.X++
	case "L":
		p.X--
	}
}

func keepUp(lead Point, tail *Point) {
	var tempdir string
	vert, tempdir := vertical(lead, tail)
	if vert {
		move(tempdir, tail)
		if lead.X == tail.X {
		} else if lead.X < tail.X {
			tail.X--
		} else {
			tail.X++
		}
	} else {
		move(tempdir, tail)
		if lead.Y == tail.Y {
		} else if lead.Y < tail.Y {
			tail.Y--
		} else {
			tail.Y++
		}
	}
}

func vertical(lead Point, tail *Point) (bool, string) {
	if (lead.Y - tail.Y) > 1 {
		return true, "U"
	}
	if (lead.Y - tail.Y) < -1 {
		return true, "D"
	}
	if (lead.X - tail.X) > 1 {
		return false, "R"
	}
	if (lead.X - tail.X) < -1 {
		return false, "L"
	}
	fmt.Println("shit")
	return false, ""
}

func notTouch(lead Point, tail Point) bool {
	if absDiffInt(lead.X, tail.X) > 1 {
		return true
	}
	if absDiffInt(lead.Y, tail.Y) > 1 {
		return true
	}
	return false
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
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
