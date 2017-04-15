package main

import (
	"fmt"
	"time"
)

type appendFunc func([]int, int) []int

func main() {
	test("my version", appendInt2)
	test("version from book", appendInt)
	test("vanilla version", vanilla)
}

func test(what string, f appendFunc) {
	var x, y []int

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		y = f(x, i)
		//fmt.Printf("%d cap=%d      %v\n", i, cap(y), y)
		x = y
	}
	fmt.Println(what, "took\t\t", time.Since(start))
}

func vanilla(x []int, y int) []int {
	return append(x, y)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// there is room to grow
		z = x[:zlen]
	} else {
		// there is insufficient space. allocate a new array
		// grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendInt2(x []int, y int) []int {
	//var z []int
	if cap(x) == 0 {
		z := make([]int, 1)
		z[0] = y
		return z
	} else if cap(x) > len(x) {
		zlen := len(x) + 1
		z := x[:zlen]
		z[len(x)] = y
		return z
	} else {
		newlen := len(x) + 1
		newcap := cap(x) * 2
		z := make([]int, newlen, newcap)
		copy(z, x)
		z[cap(x)] = y
		return z
	}
}
