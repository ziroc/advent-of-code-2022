package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

var all []any

var paircount int
var total int

func main() {
	fmt.Println("Starting Day 13 - 2")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day13.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	all = make([]any, 0, 300)
	var lines [2]string

	for fileScanner.Scan() {
		lines[0] = fileScanner.Text()
		fileScanner.Scan()
		lines[1] = fileScanner.Text()
		fileScanner.Scan()
		paircount++
		if isCorrect(lines) {
			total += paircount
		}

	}
	var el1, el2 any
	err = json.Unmarshal([]byte("[[2]]"), &el1)
	check(err)
	err = json.Unmarshal([]byte("[[6]]"), &el2)
	check(err)
	all = append(all, el1, el2)

	var res = 1
	sort.Slice(all, func(i, j int) bool {
		return compare(all[i], all[j]) < 0
	})
	for i := range all {
		packet, _ := json.Marshal(all[i])
		if string(packet) == "[[2]]" ||
			string(packet) == "[[6]]" {
			res = res * (i + 1)
		}

	}
	fmt.Println(paircount)
	fmt.Println("part 1 answer: ", total)
	fmt.Println("part 2 answer: ", res)

}

func isCorrect(lines [2]string) bool {
	var dat, dat2 any
	err := json.Unmarshal([]byte(lines[0]), &dat)
	check(err)
	err = json.Unmarshal([]byte(lines[1]), &dat2)
	check(err)

	all = append(all, dat, dat2)
	res := compare(dat, dat2)
	
	return res < 0
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
