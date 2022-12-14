package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Starting Day 3 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("day3_input.txt")
	check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var prioSum int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// fmt.Println(len(line))
		if len(line) > 0 {
			runes := []rune(line)
			size := len(runes)
			secondConmp := line[size/2:]
			for i := 0; i < size/2; i++ {
				if strings.Contains(secondConmp, line[i:i+1]) {
					fmt.Println(line[0:size/2] + " " + secondConmp + " " + line[i:i+1])
					fmt.Println("prio ", int(runes[i]), " ->", getPrio(runes[i]))
					prioSum += getPrio(runes[i])
					break;
				}

			}

		} else {
			fmt.Println("--------")
		}
	}

	file.Close()
	fmt.Println(prioSum)
}

func getPrio(char rune) int {
	ascii := int(char)
	if ascii > 90 {
		return ascii - 96
	} else {
		return ascii - 38
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

