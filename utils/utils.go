package utils

import "strconv"

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