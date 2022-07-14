package main

import "fmt"

func main() {
	// Array vacio / sin elementos
	var array [3]int
	array[0] = 2
	array[1] = 4
	fmt.Println(array)
	// Array lleno / con elementos
	nombre := [3]string{"Efren", "Martinez", "Rodriguez"}
	fmt.Println(nombre[1])
	// longitud y capacidad
	fmt.Println(len(array))
	fmt.Println(cap(nombre))
}
