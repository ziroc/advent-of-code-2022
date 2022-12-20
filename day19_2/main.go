package main

import (
	u "advent22/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var rexp *regexp.Regexp
var bp []Blueprint
var maxOre = 0
var maxObs = 0
var maxClay = 0
var bestGeodes = 0
var TIME = 32

func main() {
	fmt.Println("Starting Day 19 - 2")
	defer timeTrack(time.Now(), "MAIN")
	t0 := time.Now()
	file, err := os.Open("inputs/day19ex.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	rexp = regexp.MustCompile(`Blueprint (?P<bn>[\d]+): Each ore robot costs (?P<oreROre>[\d]+) ore. Each clay robot costs (?P<clayROre>[\d]+) ore. Each obsidian robot costs (?P<obsROre>[\d]+) ore and (?P<obsRClay>[\d]+) clay. Each geode robot costs (?P<geoROre>[\d]+) ore and (?P<geoRobs>[\d]+) obsidian.`)

	linecount := 0
	for fileScanner.Scan() && linecount<3 {
		linecount++
		readInput(fileScanner.Text())
	}
	timeTrack(t0, "read input")

	r := Resources{0, 0, 0, 0, 1, 0, 0, 0}
	res := 1
	for bpIndex := range bp {
		bestGeodes = 0
		maxOre = getMaxOre(bpIndex)
		maxClay = bp[bpIndex].obsRClay
		maxObs = bp[bpIndex].geoRObs

		fmt.Println("calc for bp ", bpIndex, bp[bpIndex])
		t := time.Now()
		bp[bpIndex].maxGeo = dfs(r, bpIndex, TIME)
		timeTrack(t, "calc "+strconv.Itoa(bpIndex))
		fmt.Println("finished calc: ", bp[bpIndex].maxGeo)
		res *= bp[bpIndex].maxGeo
	}
	fmt.Println("p2 answer", res)
}

func getMaxOre(bpIndex int) int {
	t1 := u.Max(bp[bpIndex].oreROre, bp[bpIndex].clayROre)
	t2 := u.Max(bp[bpIndex].obsROre, bp[bpIndex].geoROre)
	return u.Max(t1, t2)
}

func readInput(line string) {
	found := rexp.FindStringSubmatch(line)
	b := Blueprint{u.Toint(found[2]), u.Toint(found[3]), u.Toint(found[4]), u.Toint(found[5]), u.Toint(found[6]), u.Toint(found[7]), 0}
	bp = append(bp, b)
	fmt.Println(b)
}

func dfs(r Resources, bpIndex int, time int) (geodeCracked int) {
	currentMax := 0

	if time <= 0 || bestGeodes > r.geode+time*r.geoR+time*(time-1)/2 {
		return 0
	}
	if r.geoR >= bp[bpIndex].geoROre && r.obsR >= bp[bpIndex].geoRObs {
		bla := time*r.geoR + time*(time-1)/2
		return bla
	}

	if rr, ok := buildGeodeR(r, bpIndex); ok {
		currentMax = u.Max(currentMax, dfs(rr, bpIndex, time-1)+r.geoR) // rr.geoR-1
	}

	if rr, ok := buildObsR(r, bpIndex); r.obsR < maxObs && ok {
		currentMax = u.Max(currentMax, dfs(rr, bpIndex, time-1)+rr.geoR)
	}

	if rr, ok := buildOreR(r, bpIndex); r.oreR < maxOre && ok {
		currentMax = u.Max(currentMax, dfs(rr, bpIndex, time-1)+rr.geoR)
	}

	if rr, ok := buildClayR(r, bpIndex); r.ClayR < maxClay && ok {
		currentMax = u.Max(currentMax, dfs(rr, bpIndex, time-1)+rr.geoR)
	}

	if r.ore < maxOre || r.clay < maxClay || r.obsidian < maxObs { //wait
		r.awardRes()
		currentMax = u.Max(currentMax, dfs(r, bpIndex, time-1)+r.geoR)
	}

	bestGeodes = u.Max(bestGeodes, currentMax)
	return currentMax
}

func buildGeodeR(r Resources, bpIndex int) (Resources, bool) {
	if r.ore >= bp[bpIndex].geoROre && r.obsidian >= bp[bpIndex].geoRObs {
		r.ore -= bp[bpIndex].geoROre
		r.obsidian -= bp[bpIndex].geoRObs
		r.awardRes()
		r.geoR++
		return r, true
	}
	return r, false
}

func buildObsR(r Resources, bpIndex int) (Resources, bool) {
	if r.ore >= bp[bpIndex].obsROre && r.clay >= bp[bpIndex].obsRClay {
		r.ore -= bp[bpIndex].obsROre
		r.clay -= bp[bpIndex].obsRClay
		r.awardRes()
		r.obsR++
		return r, true
	}
	return r, false
}

func buildClayR(r Resources, bpIndex int) (Resources, bool) {
	if r.ore >= bp[bpIndex].clayROre {
		r.ore -= bp[bpIndex].clayROre
		r.awardRes()
		r.ClayR++
		return r, true
	}
	return r, false
}

func buildOreR(r Resources, bpIndex int) (Resources, bool) {
	if r.ore -= bp[bpIndex].oreROre; r.ore >= 0 {
		r.awardRes()
		r.oreR++
		return r, true
	}
	return r, false
}

func (r *Resources) awardRes() {
	r.ore += r.oreR
	r.clay += r.ClayR
	r.obsidian += r.obsR
	r.geode += r.geoR
}

type Blueprint struct {
	oreROre, clayROre, obsROre, obsRClay, geoROre, geoRObs int
	maxGeo                                                 int
}
type Resources struct {
	ore, clay, obsidian, geode int
	oreR, ClayR, obsR, geoR    int
}

func (r Resources) Add(toAdd Resources) Resources {
	nn := Resources{r.ore + toAdd.ore, r.clay + toAdd.clay, r.obsidian + toAdd.obsidian,
		r.geode + toAdd.geode, r.oreR + toAdd.oreR, r.ClayR + toAdd.ClayR, r.obsR + toAdd.obsR, r.geoR + toAdd.geoR}
	return nn
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
