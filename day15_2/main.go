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

var sensors map[Point]int
const tun = 4000000
var rexp *regexp.Regexp

func main() {
	fmt.Println("Starting Day 15 - 2")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day15.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sensors = map[Point]int{}

	rexp = regexp.MustCompile(`Sensor at x=(?P<sx>[\d-]+), y=(?P<sy>[\d-]+): closest beacon is at x=(?P<count>[\d-]+), y=(?P<count>[\d-]+)`)

	linecount := 0
	for fileScanner.Scan() {
		linecount++
		readInput(fileScanner.Text())
	}
	doWork()
}

func doWork() bool {
	for y := 0; y < tun+1; y++ {
		for x := 0; x < tun+1; x++ {
			inside := false
			for sensor, mrange := range sensors {
				if getMRange(sensor, Point{x, y}) <= mrange { //insde sensor range
					inside = true
					x = sensor.X + mrange- u.Abs(sensor.Y-y) // getRightMostX of sensor range
					break
				}
			}
			if !inside {
				fmt.Println("found: ", x, ", ", y, " : ", x*tun+y)
				return true
			}
		}
	}
	return false
}

func readInput(line string) {
	found := rexp.FindStringSubmatch(line)
	S := Point{u.Toint(found[1]), u.Toint(found[2])}
	B := Point{u.Toint(found[3]), u.Toint(found[4])}
	mrange := getMRange(S, B)

	sensors[S] = mrange
}

func getMRange(S, B Point) int {
	return u.Abs(S.X-B.X) + u.Abs(S.Y-B.Y)
}
