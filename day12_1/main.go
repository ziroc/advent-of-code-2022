package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"time"
)

var hmap [][]Point
var startPos Point
var end Point
var maxW int
var maxH int
type graph[Node comparable] map[Node][]Node
var gr graph[Point]
var count int

func main() {
	fmt.Println("Starting Day 12 - 1")
	fmt.Println("The time is", time.Now())

	file, err := os.Open("inputs/day12.txt")
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	hmap = make([][]Point, 0, 10)
	lineCount := 0
	gr = make(map[Point][]Point)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			readInput(line, lineCount)
			lineCount++
		} else {
			continue
		}
	}
	maxH = len(hmap)

	fmt.Println(startPos, " ", end)
	fmt.Println(&end)

	// _, dist, found := Path(startPos, end)

	path := FindPath[Point](gr, startPos, end, PathNeighborCost, PathEstimatedCost)
	fmt.Println("len: ", len(path)-1)
	 fmt.Println(path, " ", count)
}



func readInput(line string, lineCount int) {
	if maxW == 0 {
		maxW = (len(line))
	}

	runes := []rune(line)
	row := make([]Point, 0, maxW)
	for i := 0; i < len(runes); i++ {
		var p Point
		if runes[i] == 83 {
			p = Point{i, lineCount, 97}
			startPos = p
		} else if runes[i] == 69 {
			p = Point{i, lineCount, 122}
			end = p
		} else {
			p = Point{i, lineCount, int(runes[i])}
		}
		// nm[&p] = &node{pather: &p}
		row = append(row, p)
	}
	hmap = append(hmap, row)
}



// The Graph interface is the minimal interface a graph data structure
// must satisfy to be suitable for the A* algorithm.
type Graph[Node any] interface {
	// Neighbours returns the neighbour nodes of node n in the graph.
	Neighbours(n Node) []Node
}

// A CostFunc is a function that returns a cost for the transition
// from node a to node b.
type CostFunc[Node any] func(a, b Node) float64

// A Path is a sequence of nodes in a graph.
type Path[Node any] []Node

// newPath creates a new path with one start node. More nodes can be
// added with append().
func newPath[Node any](start Node) Path[Node] {
	return []Node{start}
}

// last returns the last node of path p. It is not removed from the path.
func (p Path[Node]) last() Node {
	return p[len(p)-1]
}

// cont creates a new path, which is a continuation of path p with the
// additional node n.
func (p Path[Node]) cont(n Node) Path[Node] {
	newPath := make([]Node, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, n)
	return newPath
}

// Cost calculates the total cost of path p by applying the cost function d
// to all path segments and returning the sum.
func (p Path[Node]) Cost(d CostFunc[Node]) (c float64) {
	for i := 1; i < len(p); i++ {
		c += d(p[i-1], p[i])
	}
	return c
}

// FindPath finds the shortest path between start and dest in graph g
// using the cost function d and the cost heuristic function h.
func FindPath[Node comparable](g Graph[Node], start, dest Node, d, h CostFunc[Node]) Path[Node] {
	closed := make(map[Node]bool)

	pq := &priorityQueue[Path[Node]]{}
	heap.Init(pq)
	heap.Push(pq, &item[Path[Node]]{value: newPath(start)})

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*item[Path[Node]]).value
		n := p.last()
		if closed[n] {
			continue
		}
		if n == dest {
			// Path found
			return p
		}
		closed[n] = true

		for _, nb := range g.Neighbours(n) {
			count++
			newPath := p.cont(nb)
			heap.Push(pq, &item[Path[Node]]{
				value:    newPath,
				priority: -(newPath.Cost(d) + h(nb, dest)),
			})
		}
	}

	// No path found
	return nil
}

type Point struct {
	X int
	Y int
	h int
}

func (g graph[Node]) Neighbours(np Point) []Point {
	return np.Neighbours()
}


func (from *Point) Neighbours() []Point {
	neighbors := []Point{}
	if from.X != 0 {
		p := hmap[from.Y][from.X-1]
		if (p.h-from.h) < 2 {
			neighbors = append(neighbors, p)
		} 
	}
	if from.X != maxW-1 {
		p := hmap[from.Y][from.X+1]
		if (p.h-from.h) < 2 {
			neighbors = append(neighbors, p)
		}
	}
	if from.Y != 0 {
		p := hmap[from.Y-1][from.X]
		if (p.h-from.h) < 2 {
			neighbors = append(neighbors, p)
		}
	}
	if from.Y != maxH-1 {
		p := hmap[from.Y+1][from.X]
		if (p.h-from.h) < 2 {
			neighbors = append(neighbors, p)
		}
	}
	fmt.Println(from , " - " ,neighbors)
	return neighbors
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func PathNeighborCost(from, to Point) float64 {
	if abs(from.h-to.h) > 1 {
		return 10000
	}

	return 1
}

func  PathEstimatedCost(from, to Point) float64 {
	absX := to.X - from.X
	if absX < 0 {
		absX = -absX
	}
	absY := to.Y - from.Y
	if absY < 0 {
		absY = -absY
	}
	r := float64(absX + absY)

	return r
}


type item[T any] struct {
	value    T       // The value of the item; arbitrary.
	priority float64 // The priority of the item in the queue.
}

// A priorityQueue implements heap.Interface and holds items.
type priorityQueue[T any] []*item[T]

func (pq priorityQueue[T]) Len() int { return len(pq) }

func (pq priorityQueue[T]) Less(i, j int) bool {
	// We want heap.Pop to give us the item with the highest,
	// not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue[T]) Push(x any) {
	*pq = append(*pq, x.(*item[T]))
}

func (pq *priorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
