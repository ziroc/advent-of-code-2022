package main

import (
	"bufio"
	"fmt"
	. "image"
	u "advent22/utils"
	"os"
	"regexp"
	"time"
)

var yline map[Point]int

const Y = 2000000

var rexp *regexp.Regexp

func main() {
	fmt.Println("Starting Day 15 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day15.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	yline = map[Point]int{}
	rexp = regexp.MustCompile(`Sensor at x=(?P<sx>[\d-]+), y=(?P<sy>[\d-]+): closest beacon is at x=(?P<count>[\d-]+), y=(?P<count>[\d-]+)`)

	linecount := 0
	for fileScanner.Scan() {
		linecount++
		readInput(fileScanner.Text())
	}
	fmt.Println("answer ", len(yline))
}

func readInput(line string) {
	found := rexp.FindStringSubmatch(line)
	S := Point{u.Toint(found[1]), u.Toint(found[2])}
	B := Point{u.Toint(found[3]), u.Toint(found[4])}
	mrange := getMRange(S, B)

	if Y >= S.Y-mrange && Y <= S.Y+mrange {
		left := mrange - u.Abs(S.Y-Y)
		count := 2*left + 1
		for i := 0; i < count; i++ {
			p:=Point{S.X - left + i, Y}
			if p== B {
				continue
			}
			yline[p] = 1
		}
	}

}

func getMRange(S, B Point) int {
	return u.Abs(S.X-B.X) + u.Abs(S.Y-B.Y)
}
