package utils

import (
	"fmt"
	i "image"
	"strconv"
	"time"
)

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Toint(s string) int {
	value, err := strconv.Atoi(s)
	Check(err)
	return value
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Init2DArray[T any](rows, cols int, initialValue T) [][]T {
	arr := make([][]T, cols)
	for i := 0; i < cols; i++ {
		arr[i] = make([]T, rows)
		for j := 0; j < rows; j++ {
			arr[i][j] = initialValue
		}
	}
	return arr
}

func Make3DArray[T any](x, y, z int) [][][]T {
	arr := make([][][]T, x)
	for i := 0; i < x; i++ {
		arr[i] = make([][]T, y)
		for j := 0; j < y; j++ {
			arr[i][j] = make([]T, z)
			// for k := 0; k < height; k++ {
			// 	arr[i][j][k] = initialValue
			// }
		}
	}
	return arr
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

var dirs8 map[string]i.Point = map[string]i.Point{"n": {0, -1}, "ne": {1, -1}, "e": {1, 0},
	"se": {1, 1}, "s": {0, 1}, "sw": {-1, 1}, "w": {-1, 0}, "nw": {-1, -1}}

var cardinals map[int]i.Point = map[int]i.Point{
	0: {1, 0}, 1: {0, 1}, 2: {-1, 0}, 3: {0, -1} } // right, down, left, up
