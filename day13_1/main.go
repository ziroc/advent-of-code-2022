package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var paircount int
var total int

func main() {
	fmt.Println("Starting Day 13 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day13.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lineCount := 0
	var line [2]string

	for fileScanner.Scan() {
		line[lineCount%2] = fileScanner.Text()
		if len(line[lineCount%2]) == 0 {
			continue
		}
		lineCount++

		if lineCount%2 == 0 {
			paircount++
			fmt.Println("++ checking ", paircount)
			res := readInput(line, lineCount)
			if res {
				fmt.Println("-- adding ", paircount)
				total += paircount
			}
		}
	}
	fmt.Println(paircount)
	fmt.Println(total)

}

func readInput(line [2]string, lineCount int) bool {
	byt := []byte(line[0])
	byt2 := []byte(line[1])
	var dat any
	var dat2 any
	err := json.Unmarshal(byt, &dat)
	check(err)
	err = json.Unmarshal(byt2, &dat2)
	check(err)

	for {
		res := compare(dat, dat2)

		if res > 0 || res < 0 {
			fmt.Println(res < 0)
			return res < 0
		}
	}
}

func compare(p1 any, p2 any) int {
	_, ok1 := p1.(float64)
	_, ok2 := p2.(float64)
	if ok1 && ok2 {
		return int(p1.(float64)) - int(p2.(float64))
	}

	if ok1 {
		p1 = []any{p1}
	}
	if ok2 {
		p2 = []any{p2}
	}

	if len(p1.([]any)) == 0 || len(p2.([]any)) == 0 {
		return len(p1.([]any)) - len(p2.([]any))
	}

	result := compare(p1.([]any)[0], p2.([]any)[0])
	if result == 0 {
		next1 := p1.([]any)[1:]
		next2 := p2.([]any)[1:]

		if len(next1) == 0 || len(next2) == 0 {
			return len(next1) - len(next2)
		}
		return compare(next1, next2)
	}

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
