package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Starting Day 3 - 2")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("day3_input.txt")
	check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var prioSum int
	count := 0;
	var group [3] string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			group[count] = line;
			if(count <2) {
				count++
			} else {
				count = 0
				runes := []rune(line)
				size := len(runes)
				for i := 0; i < size; i++ {
					char :=line[i:i+1]
					if strings.Contains(group[0], char) && strings.Contains(group[1], char) {
						fmt.Println("found ", char)
						prioSum += getPrio(runes[i])
						break;
					}
	
				}
			}
			// fmt.Println(count)
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
