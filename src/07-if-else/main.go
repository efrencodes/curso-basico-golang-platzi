package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 2
	valor2 := 4

	if valor1 < valor2 {
		fmt.Println("Valor1 => ", valor1)
	} else {
		fmt.Println("Valor2 => ", valor2)
	}

	value, err := strconv.Atoi("1")
	if err != nil {
		// interrumpe el programa
		log.Fatal(err)
	}
	fmt.Println("value => ", value)
}
