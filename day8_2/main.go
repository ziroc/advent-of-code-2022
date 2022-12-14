package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var table [][]int16
var views [][]View

var VIS int16 = 256

func main() {
	fmt.Println("Starting Day 8 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day8_input.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lineCount := 0
	table = make([][]int16, 0, 10)
	views = make([][]View, 0, 10)
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

			views[i][j].left = 1
			for k := j - 1; k >= 0; k-- {
				if k == 0 {
					views[i][j].left = uint16(j)
				}
				if row[k]|VIS >= row[j]|VIS {
					views[i][j].left = uint16(j - k)
					break
				}
			}
		}
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
			views[i][j].right = 1
			for k := j + 1; k <= width-1; k++ {
				if k == width-1 {
					views[i][j].right = uint16(width - 1 - j)
				}
				if row[k]|VIS >= row[j]|VIS {
					views[i][j].right = uint16(k - j)
					break
				}
			}

		}
		// fmt.Println("row ",i, " :", row)
	}

	// top to bottom
	for i := 1; i < width; i++ {
		var prevH int16 = -1
		for j := 0; j < height-1; j++ {
			if j == 0 {
				prevH = table[j][i]
				continue
			}
			if prevH|VIS < (table[j][i] | VIS) {

				prevH = table[j][i]
				table[j][i] = table[j][i] | VIS
			}
			views[j][i].top = 1
			for k := j - 1; k >= 0; k-- {

				if k == 0 {
					views[j][i].top = uint16(j)
				}
				if table[k][i]|VIS >= table[j][i]|VIS {
					views[j][i].top = uint16(j - k)
					break
				}
			}
		}
	}

	// bottom to top
	for i := 1; i < width; i++ {
		var prevH int16 = -1
		for j := height - 1; j > 0; j-- {
			if j == 0 {
				prevH = table[j][i]
				continue
			}
			if prevH|VIS < table[j][i]|VIS {
				prevH = table[j][i]
				table[j][i] = table[j][i] | VIS
			}
			views[j][i].down = 1

			for k := j + 1; k <= height-1; k++ {
				if k == height-1 {
					views[j][i].down = uint16(height - 1 - j)
				}
				if table[k][i]|VIS >= table[j][i]|VIS {
					views[j][i].down = uint16(k - j)
					break
				}
			}
		}
	}

	var max = 0
	for i := 1; i < len(table)-1; i++ {
		row := table[i]

		for j := 1; j < len(row)-1; j++ {
			temp := mult(views[i][j])
			if max < temp {
				max = temp
			}
		}
	}

	fmt.Println("max: ", max)
}

func mult(view View) int {
	return int(view.left) * int(view.right) * int(view.top) * int(view.down)
}

func readInput(line string, lineCount int) {
	runes := []rune(line)
	row := make([]int16, len(runes))
	for i := range runes {
		row[i] = int16(runes[i] - 48)
	}
	views = append(views, make([]View, len(runes)))
	table = append(table, row)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type View struct {
	left  uint16
	right uint16
	top   uint16
	down  uint16
}
