package main

import (
	"fmt"
	"math/cmpx"
)

func add(x, y int) int {
	return x+y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java  = true, false, "no!"

var (
	ToBe 	bool 		= false
	MaxInt 	unit64		= 1<<64-1
	z		complex128	= cmpx.Sqrt(-5+12i)
)

func main() {
    fmt.Println(add(42,13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i, j int = 1, 2
	//can only use the := syntax inside functions
	k := 3
	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

}