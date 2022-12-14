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

// var visited [] Point
var head Point
var tail Point
const t = true
var visited map[Point]Empty 
// var VIS int16 = 256

func main() {
	fmt.Println("Starting Day 9 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day9.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lineCount := 0
	visited =  map[Point]Empty{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			readInput(line, lineCount)
			lineCount++
			// if lineCount > 15 {
			// 	break
			// }
			fmt.Println("--------")
		} else {
			break
		}
	}
	fmt.Println("--------")
	fmt.Println(len(visited))
}


func readInput(line string, lineCount int) {

	split :=strings.Split(line, " ")

	doWork(split[0], toint(split[1]))

	// runes := []rune(line)
	// row := make([]int16, len(runes))
	// for i := range runes {
	// 	row[i] = int16(runes[i] - 48)
	// }
	// views = append(views, make([]View, len(runes)))
	// table = append(table, row)
}

func doWork(dir string, steps int) {
	for i := 0; i < steps; i++ {
		move(dir, &head)
		if(notTouch()) {
			// fmt.Println("not touching: ", head, "  ", tail)
			keepUp(dir);
		}
		visited[tail] = Empty{}
	}
	// fmt.Println(visited)

}

func move (dir string, p *Point) {
	switch dir {
		case "U": p.Y++
		case "D": p.Y--
		case "R": p.X++
		case "L": p.X-- 
		}
}

func keepUp(dir string) {

	if vertical(dir) {
		move(dir, &tail)
		if head.X == tail.X {			
		} else if(head.X < tail.X) {
			tail.X--
		} else { 
			tail.X++
		}
	} else {
		move(dir, &tail) 
		if head.Y == tail.Y {			
			} else if(head.Y < tail.Y) {
				tail.Y--
			} else { 
				tail.Y++
			}
	}
}

func vertical(dir string) bool{
	return dir == "U" || dir == "D"
}

func notTouch()bool {
	if absDiffInt(head.X, tail.X)> 1 {
		return true
	} 
	if absDiffInt(head.Y, tail.Y)> 1 {
		return true
	} 
	return false
}


func absDiffInt(x, y int16) int16 {
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
	X int16
	Y int16
}
