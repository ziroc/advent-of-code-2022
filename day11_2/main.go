package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var monkeys []Monkey
var lcm int = 1

func main() {
	fmt.Println("Starting Day 11 - 2")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day11.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	monkeys = make([]Monkey, 0, 7)
	lineCount := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			readInput(line, fileScanner)
			lineCount++
		} else {
			continue
		}
	}
	for i := range monkeys {
		fmt.Println(i, " : ", &monkeys[i], " ", monkeys[i])

	}
	fmt.Println("--------")
	for i := 0; i < 10000; i++ {
		doWork(i)
	}

	var ii = make([]int,0, 5)
	for i := range monkeys {
		fmt.Println(i, " inspect: ", monkeys[i].inspectCount)
		ii = append(ii, monkeys[i].inspectCount)
	}

	sort.Ints(ii)
	fmt.Println(ii[7], " ", ii[6])

}

func readInput(line string, fs *bufio.Scanner) {
	if strings.HasPrefix(line, "Monkey") {
		m := new(Monkey)

		fs.Scan()
		line = fs.Text()
		ind := strings.Index(line, ":")
		items := strings.Split(line[ind+2:], ", ")
		for i := range items {
			m.items = append(m.items, toint64(items[i]))
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

		lcm = lcm * m.testDiv
		monkeys = append(monkeys, *m)
	}
}

func doWork(round int) {
	var item uint64
	for i := range monkeys {
		monkey := &monkeys[i]
		for len(monkey.items) > 0 {

			applyOp(monkey, 0)
			monkey.inspectCount++
			if monkey.items[0]%uint64(monkey.testDiv) == 0 {
				monkey.items, item = monkey.items.PopOut()
				monkeys[monkey.yesTo].items = monkeys[monkey.yesTo].items.Push(item)

			} else {
				monkey.items, item = monkey.items.PopOut()
				monkeys[monkey.noTo].items = monkeys[monkey.noTo].items.Push(item)
			}
		}
	}

}

func applyOp(monkey *Monkey, i int) {
	var varToUse uint64
	if monkey.opConst == "old" {
		varToUse = monkey.items[i]
	} else {
		varToUse = toint64(monkey.opConst)
	}

	switch monkey.operation {
	case "*":
		monkey.items[i] = monkey.items[i] * varToUse
	case "+":
		monkey.items[i] = monkey.items[i] + varToUse
	}
	monkey.items[i] = monkey.items[i] % uint64(lcm)
}

func toint(s string) int {
	value, err := strconv.Atoi(s)
	check(err)
	return value
}

func toint64(s string) uint64 {
	value, err := strconv.Atoi(s)
	check(err)
	return uint64(value)
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

type stack []uint64

func (s stack) Push(v uint64) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, uint64) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) PopOut() (stack, uint64) {
	l := len(s)
	return s[1:l], s[0]
}

func (s stack) Peek() uint64 {
	return s[len(s)-1]
}

func (s stack) PushFront(v uint64) stack {
	return append([]uint64{v}, s...)
}
