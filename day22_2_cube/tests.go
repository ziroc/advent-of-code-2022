package main

import "fmt"

func test1() { //1 to 6
	X = 56
	Y = 0
	curDir = 3
	fmt.Println("START TEST 1: ", X, Y)
	wrapX(true)
	fmt.Println("test res: ", X, Y)
	if X != 0 || Y != 156 || curDir != 0 {
		panic("TEST failed!")
	}
}

func test1R() { // 6 to 1
	X = 0
	Y = 156
	curDir = 2
	fmt.Println("START TEST 1R: ", X, Y)
	wrapX(false)
	fmt.Println("test 1R res: ", X, Y, curDir)
	if X != 56 || Y != 0 || curDir != 1 {
		panic("TEST failed!")
	}
}

func test2() { //2 to 6
	X = 106
	Y = 0
	curDir = 3
	fmt.Println("START TEST 2: ", X, Y)
	wrapX(true)
	fmt.Println("test res: ", X, Y)
	if X != 6 || Y != 199 || curDir != 3 {
		panic("TEST failed!")
	}
}

func test3() { // 2 to 4
	X = 149
	Y = 5
	curDir = 0
	fmt.Println("START TEST 3: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 3 res: ", X, Y, curDir)
	if X != 99 || Y != 144 || curDir != 2 {
		panic("TEST failed!")
	}
}
func test3R() { // 4 to 2
	X = 99
	Y = 144
	curDir = 0
	fmt.Println("START TEST 3R: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 3 res: ", X, Y, curDir)
	if X != 149 || Y != 5 || curDir != 2 {
		panic("TEST failed!")
	}
}

func test4() { // 3 to 2
	X = 99
	Y = 55
	curDir = 0
	fmt.Println("START TEST 4: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 4 res: ", X, Y, curDir)
	if X != 105 || Y != 49 || curDir != 3 {
		panic("TEST failed!")
	}
}

func test4R() { // 2 to 3
	X = 105
	Y = 49
	curDir = 1
	fmt.Println("START TEST 4R: ", X, Y, curDir)
	wrapX(true)
	fmt.Println("test 4R res: ", X, Y, curDir)
	if X != 99 || Y != 55 || curDir != 2 {
		panic("TEST failed!")
	}
}

func test5() { // 2 to 4
	X = 149
	Y = 8
	curDir = 0
	fmt.Println("START TEST 5: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 5 res: ", X, Y, curDir)
	if X != 99 || Y != 141 || curDir != 2 {
		panic("TEST failed!")
	}
}
func test5R() { // 4 to 2
	X = 99
	Y = 141
	curDir = 0
	fmt.Println("START TEST 5: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 5 res: ", X, Y, curDir)
	if X != 149 || Y != 8 || curDir != 2 {
		panic("TEST failed!")
	}
}

func test6() { // 4 to 6
	X = 55
	Y = 149
	curDir = 1
	fmt.Println("START TEST 6: ", X, Y, curDir)
	wrapX(true)
	fmt.Println("test 6 res: ", X, Y, curDir)
	if X != 49 || Y != 155 || curDir != 2 {
		panic("TEST failed!")
	}
}

func test6R() { // 4 to 6
	X = 49
	Y = 155
	curDir = 0
	fmt.Println("START TEST 6R: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 6R res: ", X, Y, curDir)
	if X != 55 || Y != 149|| curDir != 3 {
		panic("TEST failed!")
	}
}

func test7() { // 3 to 5
	X = 50
	Y = 71
	curDir = 2
	fmt.Println("START TEST 7: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 7 res: ", X, Y, curDir)
	if X != 21 || Y != 100 || curDir != 1 {
		panic("TEST failed!")
	}
}

func test7R() { // 5 to 3
	X = 21
	Y = 100
	curDir = 3
	fmt.Println("START TEST 7R: ", X, Y, curDir)
	wrapX(true)
	fmt.Println("test 7R res: ", X, Y, curDir)
	if X != 50 || Y != 71 || curDir != 0 {
		panic("TEST failed!")
	}
}


func test8() { // 1 to 5
	X = 50
	Y = 8
	curDir = 2
	fmt.Println("START TEST 8: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 8 res: ", X, Y, curDir)
	if X != 0 || Y != 141 || curDir != 0 {
		panic("TEST failed!")
	}
}


func test8R() { // 5 to 1
	X = 0
	Y = 141
	curDir = 2
	fmt.Println("START TEST 8: ", X, Y, curDir)
	wrapX(false)
	fmt.Println("test 8 res: ", X, Y, curDir)
	if X != 50 || Y != 8 || curDir != 0 {
		panic("TEST failed!")
	}
}

func test() {
	test1()
	test1R()
	test2()
	test3()
	test3R()
	test4()
	test4R()
	test5()
	test5R()
	test6()
	test6R()
	test7()
	test7R()
	test8()
	test8R()
}
