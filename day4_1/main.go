package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Range struct{
	lower, upper int 
}

func main() {
	fmt.Println("Starting Day 4 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("day4_input.txt")
	check(err)
	defer file.Close();

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	count := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			doWork(line, &count)
		} else {
			fmt.Println("--------")
		}
	}

	fmt.Println(count)
}

func doWork(line string, count *int) {
	split := strings.Split(line, ",")

	firstArr := strings.Split(split[0],"-")
	secondArr := strings.Split(split[1],"-")
	first := Range{toint(firstArr[0]), toint(firstArr[1])}
	second := Range{toint(secondArr[0]), toint(secondArr[1])}
	if(first.lower <= second.lower && first.upper >= second.upper) {
		*count++
		return;
	} else if(first.lower >= second.lower && first.upper <= second.upper) {
		*count++
		return;
	} 
	fmt.Println(line, " ", first, " ", second)
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
