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

const SIZE = 5000

var ZERO = -2
var KEY = 811589153
var staticM map[string]int
var allM map[string]Monkey
var rexp *regexp.Regexp

func main() {
	fmt.Println("Starting Day 21 - 1")
	defer timeTrack(time.Now(), "MAIN")
	file, err := os.Open("inputs/day21.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	allM=make(map[string]Monkey)
	staticM = make(map[string]int)
	rexp = regexp.MustCompile(`(?P<m1>[a-z]+) (?P<op>[\+\-*/]) (?P<m2>[a-z]+)`)

	linecount := 0
	for fileScanner.Scan() && linecount < SIZE {
		readInput(fileScanner.Text(), linecount)
		linecount++
	}

	fmt.Println(bfs("root"))
	fmt.Println("p1 answer")
}

func bfs (curId string) int {
	fmt.Println("cur" , curId)
	res, ok := staticM[curId]
	if ok {
		return res
	}
	m := allM[curId] 
	return m.Op( bfs(m.dep1), bfs(m.dep2))
}


func readInput(line string, linecount int) {
	split := strings.Split(line, ": ")
	value, err := strconv.Atoi(split[1])

	if err == nil {
		fmt.Println(split[0], " -- ", value)
		staticM[split[0]] = value
	} else {
		found := rexp.FindStringSubmatch(split[1])
		fmt.Println(split[0], " -- ", found[1], found[2], found[3])

		Op := map[string]func(v1, v2 int) int{
			"+": func(v1, v2 int) int { return v1 + v2 },
			"-": func(v1, v2 int) int { return v1 - v2 },
			"*": func(v1, v2 int) int { return v1 * v2 },
			"/": func(v1, v2 int) int { return v1 / v2 },
		}[found[2]]

		m := Monkey{split[0], found[1], found[3], Op}
		allM[m.id] = m
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

type Monkey struct {
	id         string
	dep1, dep2 string
	Op         func(v1, v2 int) int
	// nxt *Monkey
}
