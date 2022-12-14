package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

var columns int
var rexp *regexp.Regexp

func main() {
	fmt.Println("Starting Day 6 - 1")
	fmt.Println("The time is", time.Now())

	// rexp = regexp.MustCompile(`move (?P<count>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	file, err := os.Open("day6_input.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			doWork(line)
		} else {
			fmt.Println("--------")
			break
		}
	}

}



func doWork(line string) int {

	var last4 [4]rune
	count := 0
	runes := []rune(line)
	fmt.Println("total l: ", len(runes))
	for ; count < len(runes); count++ {

		// fmt.Println(last4)
		if count < 3 {
			last4[count] = runes[count]
			continue
		}
		last4[count%4] = runes[count]
		if diff(last4) {
			fmt.Println(last4)
			break
		}
	}
	fmt.Println(line[count-3:count+1])
	fmt.Println(count)
	return count
}

func diff(last4 [4]rune) bool {

	for i := 0; i < len(last4); i++ {
		for j := i+1; j < len(last4); j++ {
			if last4[i] == last4[j] {
				return false 
			}
		}
	}
	
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

