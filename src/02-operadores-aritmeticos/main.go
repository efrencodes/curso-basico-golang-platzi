package main

import "fmt"

func main() {
	x := 10
	y := 50

	result := x + y
	fmt.Println("SUMA => ", result)
	result = x - y
	fmt.Println("RESTA => ", result)
	result = x * y
	fmt.Println("MULTIPLICACIÓN => ", result)
	result = y / x
	fmt.Println("DIVISIÓN => ", result)
	result = y % x
	fmt.Println("MODULO => ", result)
	x++
	fmt.Println("INCREMENTAL => ", x)
	x--
	fmt.Println("DECREMENTAL => ", x)
}
