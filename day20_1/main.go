package main

import (
	u "advent22/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

const SIZE = 5000
var zero = -2
var origin [SIZE]int

func main() {
	fmt.Println("Starting Day 20 - 1")
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

	var cc [SIZE]int =[SIZE]int{}
	for i := range cc {
		cc[i] = i
	}
	operatingSlice := cc[:]
	// fmt.Println(operatingSlice)
	
	fmt.Println("-------------")
	for i := 0; i < SIZE; i++ {
		apply(operatingSlice, i)
		// fmt.Println(operatingSlice)
	}

	/// finding ZERO
	currentPos := 0
	for i, item := range operatingSlice {
		if item == zero {
			currentPos = i
			break
		}
	}
	// hm,_:=findFuturePos(len(operatingSlice), currentPos+1, 1000)
	// hm2,_:=findFuturePos(len(operatingSlice), currentPos+1, 2000)
	// hm3,_:=findFuturePos(len(operatingSlice), currentPos+1, 1002)
	fmt.Println("zero pos: ", currentPos)
	hm := (1000 + currentPos) % SIZE
	hm2 := (2000 + currentPos) % SIZE
	hm3 := (3000 + currentPos) % SIZE
	fmt.Println("p1 answer", origin[ operatingSlice[hm]]+origin[operatingSlice[hm2]]+origin[operatingSlice[hm3]])

}

func apply(masiv []int, target int) {
	currentPos := curPos(masiv, target)
	fmt.Println("curpos" ,currentPos)
	realTargeg := origin[target]
	fut := findFuturePos(currentPos, realTargeg)
	fmt.Println("futrpos" ,fut)
	move(masiv, currentPos, fut)
}

func curPos (masiv []int, target int) (currentPos int) {
	for i, item := range masiv {
		if item == target {
			currentPos = i
			break
		}
	}
	return
}

func findFuturePos(currentPos, target int) (future int) {
	// fmt.Println(currentPos, target, length)
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
	// fmt.Println(positions)
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
	// c:=curPos(origin[:], num)
	if num ==0 {
		zero = linecount
	}
	// fmt.Println(num, linecount)
	origin[linecount] = num
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
