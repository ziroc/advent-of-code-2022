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


func main() {
	fmt.Println("Starting Day 10 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day10.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lineCount := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			doWork(line, lineCount)
			lineCount++
			// if lineCount > 15 {
			// 	break
			// }

		} else {
			break
		}
	}
	fmt.Println("--------")
	fmt.Println(total)
}


func doWork(line string, count int) {
	fmt.Println(line, " reg: ", registerX, ", cycle: ", cycleCount + 1, " ", total)

	if strings.HasPrefix(line, "noop") {
		cycleCount++
		checkSignal()
	} else {
		cycleCount++
		checkSignal()

		cycleCount++
		checkSignal()
		split := strings.Split(line, " ")
		registerX += toint(split[1])
	}

}

func checkSignal() {
	if cycleCount == 20 {
		fmt.Println(cycleCount, "  ", registerX, " * ", cycleCount)
		total += registerX*cycleCount
	} else if (cycleCount-20)%40 == 0 {
		total += registerX*cycleCount
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
