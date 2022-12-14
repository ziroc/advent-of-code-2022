package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var cycleCount = 0
var registerX = 1
var total = 0

var screen [240]rune

func main() {
	fmt.Println("Starting Day 10 - 2")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day10.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			doWork(line)

		} else {
			break
		}
	}
	fmt.Println("--------")
	fmt.Println(total)
	drawScreen()
}

func doWork(line string) {
	if strings.HasPrefix(line, "noop") {
		cycleCount++
		draw()
		checkSignal()
	} else {
		cycleCount++
		draw()
		checkSignal()

		cycleCount++
		draw()
		checkSignal()
		split := strings.Split(line, " ")
		registerX += toint(split[1])
	}

}

func draw() {
	var pixel = (cycleCount - 1) % 40
	var row = (cycleCount - 1) / 40
	// fmt.Println("drawin pixel: ", pixel, " row: ", row)

	if pixel >= registerX-1 && pixel <= registerX+1 {
		screen[pixel+row*40] = 35 // #
	} else {
		screen[pixel+row*40] = 46 // .
	}

}

func checkSignal() {
	if cycleCount == 20 {
		total += registerX * cycleCount
	} else if (cycleCount-20)%40 == 0 {
		total += registerX * cycleCount
	}
}

func drawScreen() {
	for i := 0; i < 240; i++ {
		if i%40 == 0 {
			fmt.Println()
		}
		fmt.Print(string(screen[i]))
	}
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
