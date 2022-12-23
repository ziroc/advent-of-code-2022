package main

import (
	u "advent22/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

const SIZE = 5000

var ZERO = -2
var KEY = 811589153
var origin [SIZE]int

func main() {
	fmt.Println("Starting Day 20 - 2")
	defer timeTrack(time.Now(), "MAIN")
	file, err := os.Open("inputs/day20.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	linecount := 0
	for fileScanner.Scan() && linecount < SIZE {
		readInput(fileScanner.Text(), linecount)
		linecount++
	}

	var cc [SIZE]int = [SIZE]int{}
	for i := range cc {
		cc[i] = i
	}
	operatingSlice := cc[:]

	for k := 0; k < 10; k++ {
		for i := 0; i < SIZE; i++ {
			apply(operatingSlice, i)
		}
	}
	/// finding ZERO
	currentPos := 0
	for i, item := range operatingSlice {
		if item == ZERO {
			currentPos = i
			break
		}
	}
	hm := (1000 + currentPos) % SIZE
	hm2 := (2000 + currentPos) % SIZE
	hm3 := (3000 + currentPos) % SIZE
	fmt.Println("p1 answer", origin[operatingSlice[hm]]+origin[operatingSlice[hm2]]+origin[operatingSlice[hm3]])
}

func apply(masiv []int, target int) {
	currentPos := curPos(masiv, target)
	realTarget := origin[target]
	fut := findFuturePos(currentPos, realTarget)
	move(masiv, currentPos, fut)
}

func curPos(masiv []int, target int) int {
	for i, item := range masiv {
		if item == target {
			return i
		}
	}
	return 0
}

func findFuturePos(currentPos, target int) (future int) {
	if target > 0 {
		temp := (currentPos + target) % (SIZE - 1)
		return temp
	} else { // < 0
		if temp := currentPos + target; temp <= 0 {
			temp = u.Abs(temp) % (SIZE - 1)
			return SIZE - temp - 1
		} else {
			return temp
		}
	}
}

func insertInt(array []int, value int, index int) []int {
	n := append(array[:index], append([]int{value}, array[index:]...)...)
	return n
}

func removeInt(array []int, index int) []int {
	n := append(array[:index], array[index+1:]...)
	return n
}

func move(array []int, srcIndex int, dstIndex int) []int {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

func readInput(line string, linecount int) {
	num := u.Toint(line)
	if num == 0 {
		ZERO = linecount
	}
	origin[linecount] = num * KEY
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
