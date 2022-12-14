package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var cycleCount = 0
var registerX = 1
var total = 0

var monkeys []*Monkey

func main() {
	fmt.Println("Starting Day 11 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day11.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	monkeys = make([]*Monkey, 0, 7)
	lineCount := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			readInput(line, fileScanner)
			lineCount++
			// if lineCount > 15 {
			// 	break
			// }

		} else {
			continue
		}
	}
	for i := range monkeys {
		fmt.Println(i, " : ", monkeys[i])

	}
	fmt.Println("--------")
	for i := 0; i < 20; i++ {
		doWork(i)
		// fmt.Println(monkeys)
	}
	for i := range monkeys {
		fmt.Println(i, " inspect: ", monkeys[i].inspectCount)
	}
}

func readInput(line string, fs *bufio.Scanner) {
	if strings.HasPrefix(line, "Monkey") {
		m := new(Monkey)

		fs.Scan()
		line = fs.Text()
		ind := strings.Index(line, ":")
		items := strings.Split(line[ind+2:], ", ")
		for i := range items {
			m.items = append(m.items, toint(items[i]))
		}
		fs.Scan()
		line = fs.Text()
		ind = strings.Index(line, "old ")
		op := line[ind+4 : ind+5]
		m.operation = op
		m.opConst = line[ind+6:]

		fs.Scan()
		line = fs.Text()
		ind = strings.Index(line, " by ")
		m.testDiv = toint(line[ind+4:])

		fs.Scan()
		line = fs.Text()
		ind = strings.Index(line, "monkey")
		m.yesTo = toint(line[ind+7:])

		fs.Scan()
		line = fs.Text()
		ind = strings.Index(line, "monkey")
		m.noTo = toint(line[ind+7:])

		monkeys = append(monkeys, m)
		// fmt.Println(m)
	}
}

func doWork(round int) {
	fmt.Println("round ", round)
	var item int
	for i := range monkeys {
		monkey := monkeys[i]
		// fmt.Println("\nstart monkey: ", i)
		// fmt.Println(monkey)
		for len(monkey.items) > 0 {

			applyOp(monkey, 0)
			monkey.inspectCount++
			// fmt.Println("r:", round, "m:", i, " ins:", monkey.inspectCount)
			monkey.items[0] = monkey.items[0] / 3
			// fmt.Println(monkey.items[0])

			if monkey.items[0]%monkey.testDiv == 0 {
				monkey.items, item = monkey.items.PopOut()
				monkeys[monkey.yesTo].items = monkeys[monkey.yesTo].items.Push(item)
				// fmt.Println("yes", monkey)
				// fmt.Println( monkeys[monkey.yesTo])
			} else {
				// fmt.Println("!@#!@#",monkey.noTo)
				// fmt.Println( monkeys[monkey.noTo])

				monkey.items, item = monkey.items.PopOut()
				monkeys[monkey.noTo].items = monkeys[monkey.noTo].items.Push(item)
				// fmt.Println("no", monkey)
				// fmt.Println( monkeys[monkey.noTo])
			}
		}
	}

}

func applyOp(monkey *Monkey, i int) {
	var varToUse int
	if monkey.opConst == "old" {
		varToUse = monkey.items[i]
	} else {
		varToUse = toint(monkey.opConst)
	}

	switch monkey.operation {
	case "*":
		monkey.items[i] = monkey.items[i] * varToUse
	case "+":
		monkey.items[i] = monkey.items[i] + varToUse
	}
}

func toint(s string) int {
	value, err := strconv.Atoi(s)
	check(err)
	return value
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Monkey struct {
	items        stack
	operation    string
	opConst      string
	testDiv      int
	yesTo        int
	noTo         int
	inspectCount int
}

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) PopOut() (stack, int) {
	l := len(s)
	return s[1:l], s[0]
}

func (s stack) Peek() int {
	return s[len(s)-1]
}

func (s stack) PushFront(v int) stack {
	return append([]int{v}, s...)
}
