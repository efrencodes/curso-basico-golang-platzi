package main

import "fmt"

func main() {
	// defer
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	// break y continue
	for i := 1; i < 10; i++ {
		fmt.Println(i)

		// continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// break
		if i == 7 {
			fmt.Println("Es 8")
			break
		}
	}
}
