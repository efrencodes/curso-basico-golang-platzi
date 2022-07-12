package main

import "fmt"

func main() {
	// For incremental
	var count = 10
	for i := 1; i <= count; i++ {
		fmt.Println("n => ", i)
	}
	// For while
	counter := 1
	for counter <= 10 {
		fmt.Println("n => ", counter)
		counter++
	}
	// For forever
	counterForever := 1
	for {
		fmt.Println("n => ", counterForever)
		counterForever++
	}
}
