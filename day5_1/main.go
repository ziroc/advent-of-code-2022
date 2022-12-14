package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)


var columns int
var stacks []stack
var rexp *regexp.Regexp;


func main() {
	fmt.Println("Starting Day 5 - 1")
	fmt.Println("The time is", time.Now())
	
	rexp = regexp.MustCompile(`move (?P<count>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	file, err := os.Open("day5_input.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	count := 0
	inputReadDone := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			count++
			if inputReadDone == 0 {
				inputReadDone = readInput(line)
			} else {
				doWork(line, &count)
			}
		} else {
			fmt.Println("--------")
			if count < inputReadDone + 3 {
				continue
			} else {
				break
			}
		}
	}

	for i := range stacks {
		fmt.Print(string(stacks[i].Peek()))
	}
}

func readInput(line string) int {
	runes := []rune(line)
	if columns == 0 {
		columns = (len(runes) + 1) / 4
		fmt.Println(columns)
		stacks = make([]stack, columns)
		for i := range stacks {
			stacks[i] = make(stack, 0)
		}
	}

	if strings.HasPrefix(line, " 1") {
		return columns
	}

	for i := 0; i < len(runes); {
		if runes[i] == '[' {
			item := runes[i+1]
			stacks[i/4] = stacks[i/4].PushFront(item)
		}
		if i >= len(line)-3 {
			break
		}
		i = i + 4
	}
	return 0
}

func doWork(line string, count *int) {

	found := rexp.FindStringSubmatch(line)

	howmany := toint(found[1])
	from := toint(found[2])-1
	to := toint(found[3])-1

	for i := 0; i < howmany; i++ {
		ff, el := stacks[from].Pop()
		stacks[from] = ff;
		stacks[to] = stacks[to].Push(el)
	}
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

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() (rune) {
	return s[len(s)-1]
}

func (s stack) PushFront(v rune) stack {
	return append([]rune{v}, s...)
}
