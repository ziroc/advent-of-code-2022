package main

import (
	u "advent22/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var valves []Valve
var valveIndexMap map[string]int
var costs [][]int
var neighbours map[string][]string

var rexp *regexp.Regexp

func main() {
	fmt.Println("Starting Day 16 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day16.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	valves = make([]Valve, 0, 10)
	valveIndexMap = map[string]int{}
	neighbours = map[string][]string{}

	rexp = regexp.MustCompile(`Valve (?P<valve>[A-Z]+) has flow rate=(?P<sy>[\d]+); tunnels? leads? to valves? (?P<valves>[A-Z\s,]+)`)

	linecount := 0
	for fileScanner.Scan() {
		linecount++
		readInput(fileScanner.Text())
	}
	valveCount := len(valves)

	costs = u.Init2DArray(valveCount, valveCount, 99)

	for vname, nbarray := range neighbours {
		for _, v1 := range nbarray {
			costs[valveIndexMap[vname]][valveIndexMap[v1]] = 1
		}
	}
	// fmt.Println(costs)
	for k := 0; k < valveCount; k++ {
		for i := 0; i < valveCount; i++ {
			for j := 0; j < valveCount; j++ {
				if costs[i][j] > costs[i][k]+costs[k][j] {
					costs[i][j] = costs[i][k] + costs[k][j]
				}
			}
		}
	}
	// fmt.Println(costs)
	rest := []int{}
	for i, v := range valves {
		if v.releaseRate > 0 {
			rest = append(rest, i)
			// fmt.Println(v, i)
		}
	}
	pressure := dfs(35, rest, 30)
	fmt.Println("res", pressure)
}

func dfs(cur int, rest []int, time int) (pressure int) {
	maxP := 0
	for i := 0; i < len(rest); i++ {
		to := rest[i]
		if to == -1 {
			continue
		}

		if costs[cur][to] < time {
			cc := make([]int, len(rest))
			copy(cc, rest)
			cc[i] = -1
			p := valves[to].releaseRate*(time-1-costs[cur][to]) + dfs(to, cc, time-1-costs[cur][to])
			if p > maxP {
				maxP = p
			}
		}
	}
	return maxP
}

func readInput(line string) {
	found := rexp.FindStringSubmatch(line)
	split := strings.Split(strings.ReplaceAll(found[3], " ", ""), ",")
	valve := Valve{name: found[1], releaseRate: u.Toint(found[2])}
	neighbours[valve.name] = split
	valves = append(valves, valve)
	valveIndexMap[valve.name] = len(valves) - 1
}

type Valve struct {
	name        string
	releaseRate int
}
