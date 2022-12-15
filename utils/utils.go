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