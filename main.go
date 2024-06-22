package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	a := 42
	fmt.Println(a)
	b, c, d, e, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, e, f)
	var h int = 44
	fmt.Println(h)
	adams := 43
	fmt.Println("42 as binary , &b \n ", adams)

	fmt.Println("My favorite number is", math.Pi)

	fmt.Println(add(42, 13))
	sayHellow()

	fmt.Println(split(20))
}

func sayHellow() {
	fmt.Println("hellow")
}
