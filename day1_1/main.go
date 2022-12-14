package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Starting")
	fmt.Println("The time is", time.Now())

	// food := []int {}
	var max int

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
				fmt.Println(max)
				if max < cal {
					max = cal					
				} else {
					fmt.Println("shit")					
				}
				cal = 0
			} else {
				fmt.Println("WTF")
			}
		}

	}

	f.Close()
	fmt.Println(max)
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
