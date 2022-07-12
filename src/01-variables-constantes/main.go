package main

import "fmt"

func main() {
	// Declaración de constantes.
	const PI float64 = 3.14
	const PI2 = 3.1415
	fmt.Println("PI => ", PI)
	fmt.Println("PI2 => ", PI2)

	// Declaración de variables.
	// := es la forma corta de declarar y asignar valor a una variable.
	name := "Efren"
	base := 12
	var altura int = 299
	var area int
	fmt.Println("name => ", name)
	fmt.Println("base => ", base)
	fmt.Println("altura => ", altura)
	fmt.Println("area => ", area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("a => ", a)
	fmt.Println("b => ", b)
	fmt.Println("c => ", c)
	fmt.Println("d => ", d)

	// Asignación múltiple en una linea
	edad, nombre, apellido := 22, "Efren", "Martinez"
	fmt.Println("edad => ", edad)
	fmt.Println("nombre => ", nombre)
	fmt.Println("apellido => ", apellido)
}
