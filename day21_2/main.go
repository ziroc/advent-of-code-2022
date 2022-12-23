package main

import (
	u "advent22/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var staticM map[string]int64
var allM map[string]Monkey
var rexp *regexp.Regexp

func main() {
	fmt.Println("Starting Day 21 - 2")
	defer u.TimeTrack(time.Now(), "MAIN")
	file, err := os.Open("inputs/day21.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	allM = make(map[string]Monkey)
	staticM = make(map[string]int64)
	rexp = regexp.MustCompile(`(?P<m1>[a-z]+) (?P<op>[\+\-*/]) (?P<m2>[a-z]+)`)

	linecount := 0
	for fileScanner.Scan() {
		readInput(fileScanner.Text(), linecount)
		linecount++
	}

	var h int64 = 1000
	part2 := dfs(allM["root"].dep2)

	for {
		staticM["humn"] = h
		result := dfs(allM["root"].dep1)
		diff := result - part2
		if diff == 0 {
			fmt.Println("ANSWER ", h)
			break
		}

		if diff >= 100 {
			diff = diff / 100
		} else {
			diff = 1
		}

		if result < part2 {
			h-=diff
		} else {
			h += diff
		}
	}
}

func dfs(curId string) int64 {
	res, ok := staticM[curId]
	if ok {
		return int64(res)
	}
	m := allM[curId]
	return m.Op(dfs(m.dep1), dfs(m.dep2))
}

func readInput(line string, linecount int) {
	split := strings.Split(line, ": ")
	value, err := strconv.Atoi(split[1])

	if err == nil {
		staticM[split[0]] = int64(value)
	} else {
		found := rexp.FindStringSubmatch(split[1])
		Op := map[string]func(v1, v2 int64) int64{
			"+": func(v1, v2 int64) int64 { return v1 + v2 },
			"-": func(v1, v2 int64) int64 { return v1 - v2 },
			"*": func(v1, v2 int64) int64 { return v1 * v2 },
			"/": func(v1, v2 int64) int64 { return v1 / v2 },
		}[found[2]]

		m := Monkey{split[0], found[1], found[3], Op}
		allM[m.id] = m
	}
}
type Monkey struct {
	id         string
	dep1, dep2 string
	Op         func(v1, v2 int64) int64
}
