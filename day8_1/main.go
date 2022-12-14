package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// var rexp *regexp.Regexp
var total uint32
var table [][]int16

var VIS int16 = 256

func main() {
	fmt.Println("Starting Day 8 - 1")
	fmt.Println("The time is", time.Now())

	// rexp = regexp.MustCompile(`(?P<size>\d+)\s(?P<name>[\w\.]+)`)
	file, err := os.Open("inputs/day8_input.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lineCount := 0
	table = make([][]int16, 0, 10)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			readInput(line, lineCount)
			lineCount++
			// if lineCount > 5 {
			// 	break
			// }
			fmt.Println("--------")
		} else {
			break
		}
	}
	doWork()
	fmt.Println("--------")

	total = 0
	// var total uint32 = 0
	fmt.Println(total)
}

func doWork() {
	height := len(table)
	width := len(table[0])
	fmt.Println("table height: ", height)
	fmt.Println("table width: ", width)

	//left to right
	for i := 1; i < len(table); i++ {
		var prevH int16 = -1
		row := table[i]
		for j := 0; j < len(row)-1; j++ {
			if j == 0 {
				prevH = row[j]
				continue
			}
			if prevH|VIS < (row[j] | VIS) {
				prevH = row[j]
				row[j] = row[j] | VIS
			}
		}
		// fmt.Println("row : ", row)
	}

	// right to left
	for i := 1; i < len(table); i++ {
		var prevH int16 = -1
		row := table[i]
		for j := width - 1; j > 0; j-- {
			if j == width-1 {
				prevH = row[j]
				continue
			}
			if prevH|VIS < (row[j] | VIS) {
				prevH = row[j]
				row[j] = row[j] | VIS
			}
		}
		// fmt.Println("row ",i, " :", row)
	}

	// top to bottom
	for i := 1; i < width; i++ {
		var prevH int16 = -1			
		for j := 0; j < height -1; j++ {
			if j == 0 {
				prevH = table[j][i]
				continue
			}
			if prevH|VIS < (table[j][i] | VIS) {
				prevH = table[j][i]
				table[j][i]= table[j][i] | VIS
			}
			// if i==1 {
			// 	fmt.Println("i:j ",i, ":", j,"  ", table[j][i])
			// }
		}
	}

		// bottom to top
		for i := 1; i < width; i++ {
			var prevH int16 = -1			
			for j := height -1; j >0 ; j-- {
				if j == 0 {
					prevH = table[j][i]
					continue
				}
				if prevH|VIS < table[j][i] | VIS {
					prevH = table[j][i]
					table[j][i]= table[j][i] | VIS
				}
			}
		}

		var count =0
		for i := 1; i < len(table)-1; i++ {
			row := table[i]
			fmt.Print(i, ":  ")
			
			for j := 1; j < len(row)-1; j++ {

				fmt.Print(row[j], " ")
				if(row[j] > 15) {
					count++
				}
			}
			fmt.Println()
		}

		added := 2*height +2*width-4
		fmt.Println("\ncount: " ,count)
		fmt.Println("count: " ,count+ added)
}

func readInput(line string, lineCount int) {
	runes := []rune(line)
	row := make([]int16, len(runes))

	for i := range runes {
		// fmt.Println(int16(runes[i] - 48))
		row[i] = int16(runes[i] - 48)
	}
	// fmt.Println("row : ", row)
	table = append(table, row)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toint(s string) int {
	value, err := strconv.Atoi(s)
	check(err)
	return value
}
