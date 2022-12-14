package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Starting Day 2")
	fmt.Println("The time is", time.Now())

	var score int

	f, err := os.Open("day2_input.txt")
	check(err)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(len(line))
		if len(line) > 0 {
			// split := strings.Split(line, " ")

			score += getScore(line)
		} else {
			fmt.Println("--------")
		}
	}

	f.Close()
	fmt.Println(score)
}

func atoi(s string) int {
	value, err := strconv.Atoi(s)
	check(err)
	return value
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getScore(moves string) int {
	ret := 0
	switch moves {
	case "A X":
		ret = 1 + 3
	case "A Y":
		ret = 2 + 6
	case "A Z":
		ret = 3 + 0
	case "B X":
		ret = 1 + 0
	case "B Y":
		ret = 2 + 3
	case "B Z":
		ret = 3 + 6
	case "C X":
		ret = 1 + 6
	case "C Y":
		ret = 2 + 0
	case "C Z":
		ret = 3 + 3

	}
	return ret
}

// func getScore(theirMove string, yourMove string) int{
// 	if(theirMove== "A" && yourMove == "X")
// 		`return 3+1

// 	return -1
// }
