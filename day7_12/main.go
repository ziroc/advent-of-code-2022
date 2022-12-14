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

var current *filenode
var root *filenode
var lastCommand string = ""
var rexp *regexp.Regexp
var total uint32
var currentTotal uint32 = 48748071
var needed uint32

func main() {
	fmt.Println("Starting Day 7 - 12")
	fmt.Println("The time is", time.Now())

	rexp = regexp.MustCompile(`(?P<size>\d+)\s(?P<name>[\w\.]+)`)
	file, err := os.Open("day7_input.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	count := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			readInput(line)
			count++
		} else {
			fmt.Println("--------")
			break
		}
	}
	fmt.Println("--------")

	needed  = 30000000 -(70000000-currentTotal)
	total = 0
	// var total uint32 = 0
	total = getSize(root)
	fmt.Println(total)
}

func getSize(node *filenode) uint32 {
	if !node.isdir {
		return node.size
	}

	var sum uint32 = 0

	for i := 0; i < len(node.children); i++ {
		sum += getSize(node.children[i])
	}

	if sum > needed {
		fmt.Println("sum ", sum)
	}

	return sum
}

func readInput(line string) {

	if strings.HasPrefix(line, "$ cd ..") {
		goToParent()
	} else if strings.HasPrefix(line, "$ cd ") {
		dirname := line[5:]
		// fmt.Println("dir: ", dirname)
		checkAndChangeDir(dirname)
	} else if strings.HasPrefix(line, "$ ls") {
		// fmt.Println("ls called ")
		lastCommand = "ls"
	} else {
		if strings.HasPrefix(line, "dir ") {
			// fmt.Println("dir : ", line)
		} else {
			fmt.Println(line)
			found := rexp.FindStringSubmatch(line)
			fmt.Println(found[1], " - ", found[2])
			addFile(found[2], uint32(toint(found[1])))
		}

	}

}

func goToParent() {
	current = current.parent
}

func checkAndChangeDir(dirname string) {
	if current == nil {
		current = &filenode{name: dirname, isdir: true}
		root = current
		return
	}
	for i := 0; i < len(current.children); i++ {
		if current.children[i].name == dirname {
			current = current.children[i]
			return
		}
	}
	newdir := &filenode{name: dirname, isdir: true, parent: current}
	current.children = append(current.children, newdir)
	current = newdir
	fmt.Println("switched dir to ", dirname)
}

func addFile(filename string, size uint32) {
	for i := 0; i < len(current.children); i++ {
		if current.children[i].name == filename {
			return
		}
	}
	newfile := &filenode{name: filename, isdir: false, parent: current, size: size}
	// total += size
	current.children = append(current.children, newfile)
	fmt.Println("added file: ", newfile)
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

type filenode struct {
	children []*filenode
	parent   *filenode
	name     string
	size     uint32
	isdir    bool
}
