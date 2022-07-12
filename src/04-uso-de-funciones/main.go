package main

import "fmt"

// Función con parametros
func sum(x, y int) {
	total := 0
	total = x + y
	fmt.Printf("%d + %d => %d\n", x, y, total)
}

// Función que indica el tipo de dato que retorna
func subtraction(x, y int) int {
	return x - y
}

// Función con retorno multiples
func areaRectangle(a, b int) (area int, parameter int) {
	parameter = 2 * (a * b)
	area = a * b
	return
}
func main() {
	sum(10, 40)
	subtraction := subtraction(12, 5)
	fmt.Printf("La resta es => %d\n", subtraction)
	area, parameter := areaRectangle(10, 40)
	// cuando no queremos usar un parametro.
	_, parameter1 := areaRectangle(50, 40)
	fmt.Println("Area => ", area)
	fmt.Println("parameter => ", parameter)
	fmt.Println("parameter1 => ", parameter1)
}
