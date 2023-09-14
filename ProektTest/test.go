package main

import (
	"fmt"
)

func main() {
	x, y := 5, 5
	result := rr(x, y)
	fmt.Println(result)
	res := bes(x, y)
	fmt.Println(res)
}

func rr(x, y int) int {
	r := x * y
	fmt.Println(r)
	return r
}

func bes(x, y int) int {
	r := x / y
	fmt.Println(r)
	return r
}
