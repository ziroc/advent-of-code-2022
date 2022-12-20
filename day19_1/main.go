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
func main() {
	fmt.Println("Starting Day 19 - 1")
	defer timeTrack(time.Now(), "MAIN")
	t0 := time.Now()
	file, err := os.Open("inputs/day19.txt")
	u.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	rexp = regexp.MustCompile(`Blueprint (?P<bn>[\d]+): Each ore robot costs (?P<oreROre>[\d]+) ore. Each clay robot costs (?P<clayROre>[\d]+) ore. Each obsidian robot costs (?P<obsROre>[\d]+) ore and (?P<obsRClay>[\d]+) clay. Each geode robot costs (?P<geoROre>[\d]+) ore and (?P<geoRobs>[\d]+) obsidian.`)

	linecount := 0
	for fileScanner.Scan() {
		linecount++
		readInput(fileScanner.Text())
	}
	timeTrack(t0, "read input")
	// build ore robot
	// build clay robot
	// build obs robot
	// build geode robot
	// wait
	r := Resources{0, 0, 0, 0, 1, 0, 0, 0}
	res := 0
	for bpIndex := range bp {
		// if bpIndex >= 1 {
		// 	break
		// }
		maxOre = getMax(bpIndex)
		fmt.Println("calc for bp ", bpIndex, bp[bpIndex], "max ore R:", maxOre)
		t:=time.Now()
		bp[bpIndex].maxGeo = dfs(r, bpIndex, 24)
		timeTrack(t, "calc "+strconv.Itoa(bpIndex) ) 
		fmt.Println("finished calc: ", bp[bpIndex].maxGeo)
		res += bp[bpIndex].maxGeo * (bpIndex + 1)
	}
	fmt.Println("p1 answer", res)
	fmt.Println("p2 answer")
}

func getMax(bpIndex int) int {
	t:=0
	if t< bp[bpIndex].oreROre {
		t = bp[bpIndex].oreROre
	}
	if t< bp[bpIndex].clayROre {
		t = bp[bpIndex].clayROre
	}
	if t< bp[bpIndex].obsROre {
		t = bp[bpIndex].obsROre
	}
	if t< bp[bpIndex].geoROre {
		t = bp[bpIndex].geoROre
	}
	return t
}


func readInput(line string) {
	found := rexp.FindStringSubmatch(line)
	b := Blueprint{u.Toint(found[2]), u.Toint(found[3]), u.Toint(found[4]), u.Toint(found[5]), u.Toint(found[6]), u.Toint(found[7]), 0}
	bp = append(bp, b)
	fmt.Println(b)
}

func dfs(r Resources, bpIndex int, time int) (geodeCracked int) {
	// if r.geode >= 9 {
	// 	fmt.Println(time, r)
	// }
	currentMax := 0
	if time <= 0 {
		return r.geode
	}

	if rr, ok := buildGeodeR(r, bpIndex); ok {

		currentMax = dfs(awardRes(rr, bpIndex), bpIndex, time-1)
		if currentMax > geodeCracked {
			geodeCracked = currentMax
		}
		return geodeCracked
	}
	if time == 1 {
		currentMax = dfs(awardRes(r, bpIndex), bpIndex, time-1)
		if currentMax > geodeCracked {
			geodeCracked = currentMax
		}
		return geodeCracked
	}

	if rr, ok := buildOreR(r, bpIndex); r.oreR <=maxOre && ok {
		// fmt.Println("can build ore robot")
		currentMax = dfs(awardRes(rr, bpIndex), bpIndex, time-1)
		if currentMax > geodeCracked {
			geodeCracked = currentMax
		}
	}
	if rr, ok := buildClayR(r, bpIndex); ok {
		// fmt.Println("can build clay robot")
		currentMax = dfs(awardRes(rr, bpIndex), bpIndex, time-1)
		if currentMax > geodeCracked {
			geodeCracked = currentMax
		}
	}
	if rr, ok := buildObsR(r, bpIndex); ok {
		currentMax = dfs(awardRes(rr, bpIndex), bpIndex, time-1)
		if currentMax > geodeCracked {
			geodeCracked = currentMax
		}
	}

	// fmt.Println("wait")
	currentMax = dfs(awardRes(r, bpIndex), bpIndex, time-1)
	if currentMax > geodeCracked {
		geodeCracked = currentMax

	}

	return geodeCracked
}

func buildGeodeR(r Resources, bpIndex int) (Resources, bool) {
	if r.ore >= bp[bpIndex].geoROre && r.obsidian >= bp[bpIndex].geoRObs {
		r.ore -= bp[bpIndex].geoROre
		r.obsidian -= bp[bpIndex].geoRObs
		r.geode--
		r.geoR++
		return r, true
	}
	return r, false
}

func buildObsR(r Resources, bpIndex int) (Resources, bool) {
	if r.ore >= bp[bpIndex].obsROre && r.clay >= bp[bpIndex].obsRClay {
		r.ore -= bp[bpIndex].obsROre
		r.clay -= bp[bpIndex].obsRClay
		r.obsidian--
		r.obsR++
		return r, true
	}
	return r, false
}

func buildClayR(r Resources, bpIndex int) (Resources, bool) {
	if r.ore >= bp[bpIndex].clayROre {
		r.ore -= bp[bpIndex].clayROre
		r.clay--
		r.ClayR++
		return r, true
	}
	return r, false
}

func buildOreR(r Resources, bpIndex int) (Resources, bool) {
	if r.ore >= bp[bpIndex].oreROre {
		r.ore -= bp[bpIndex].oreROre
		r.ore--
		r.oreR++
		return r, true
	}
	return r, false
}

func awardRes(r Resources, bpIndex int) Resources {
	r.ore += r.oreR
	r.clay += r.ClayR
	r.obsidian += r.obsR
	r.geode += r.geoR
	return r
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

// func valid(nn Cube) bool {
// 	return !(nn.x < 0 || nn.y < 0 || nn.z < 0 ||
// 		nn.x > 19 || nn.y > 19 || nn.z > 19)
// }

// var dirs []Cube = []Cube{{0, 0, 1}, {0, 0, -1}, {1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
   fmt.Printf("%s took %s\n", name, elapsed)
}