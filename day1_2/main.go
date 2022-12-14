package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Starting")
	fmt.Println("The time is", time.Now())

	food := []int {}

	f, err := os.Open("day1_input.txt")
	check(err)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var cal int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(len(line))
		if len(line) > 0 {
			cal = cal + atoi(line)
		} else {
			fmt.Println("--------")
			if cal > 0 {
				fmt.Println(cal)
				food = append(food, cal)
				cal = 0
			} else {
				fmt.Println("WTF")
			}
		}
	}

	f.Close()
	// sort.Sort(sort.IntSlice(food))
	sort.Ints(food)
	fmt.Println(food[0])
	fmt.Println(food[len(food)-1])

	var sum int;
	for i := len(food)-1; i >= len(food)-3; i-- {
		sum += food[i]
	 }
	 fmt.Println(sum)
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
