package main

import (
	u "advent22/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var cubes [][][]int

var real map[Cube]struct{}
var seenOutside map[Cube]struct{}
var exists = struct{}{}
var totalSides int = 0
var maxSize = 21

func main() {
	fmt.Println("Starting Day 18 - 2")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day18.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	cubes = u.Make3DArray[int](maxSize, maxSize, maxSize)
	real = make(map[Cube]struct{})
	seenOutside = make(map[Cube]struct{})
	linecount := 0
	for fileScanner.Scan() {
		linecount++
		readInput(fileScanner.Text())
	}
	fmt.Println("cubes count ", linecount)
	fmt.Println("p1 answer", totalSides)

	findOutside()
	countedSides := 0
	for realCube, _ := range real {
		for _, dir := range dirs {
			n, _ := realCube.Add(dir)
			_, seenB := seenOutside[n]
			if seenB {
				countedSides++
			}
		}
	}
	fmt.Println("p2 answer", countedSides)
}

func readInput(line string) {
	split := strings.Split(line, ",")

	c := Cube{u.Toint(split[0]), u.Toint(split[1]), u.Toint(split[2])}
	cubes[c.x][c.y][c.z] = 1
	real[c] = exists
	sidesCount := 6

	for _, v := range dirs {
		n, ok := c.Add(v)
		if !ok {
			continue
		}
		if cubes[n.x][n.y][n.z] == 1 {
			sidesCount -= 2
		}
	}
	totalSides += sidesCount
}

func findOutside() (removeSides int) {

	stack := []Cube{{0, 0, 0}}

	for len(stack) > 0 {
		c := stack[0]
		stack = stack[1:]
		_, seenB := seenOutside[c]
		_, isReal := real[c]
		if valid2(c) && !seenB && !isReal {
			seenOutside[c] = exists
			for _, v := range dirs {
				n, _ := c.Add(v)
				stack = append(stack, n)
			}
		}
	}

	fmt.Println("seen count ", len(seenOutside))
	return removeSides
}

type Cube struct {
	x, y, z int
}

func (c Cube) Add(toAdd Cube) (Cube, bool) {
	nn := Cube{c.x + toAdd.x, c.y + toAdd.y, c.z + toAdd.z}
	return nn, valid(nn)
}

func valid(nn Cube) bool {
	return !(nn.x < 0 || nn.y < 0 || nn.z < 0 ||
		nn.x > 19 || nn.y > 19 || nn.z > 19)
}

func valid2(nn Cube) bool {
	return !(nn.x < -1 || nn.y < -1 || nn.z < -1 ||
		nn.x >= maxSize || nn.y >= maxSize || nn.z >= maxSize)
}

var dirs []Cube = []Cube{{0, 0, 1}, {0, 0, -1}, {1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}}
