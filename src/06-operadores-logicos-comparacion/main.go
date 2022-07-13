package main

import "fmt"

func main() {
	valor1 := 10
	valor2 := 9

	fmt.Println("valor1 == valor 2 => ", valor1 == valor2)
	fmt.Println("valor1 != valor 2 => ", valor1 != valor2)
	fmt.Println("valor1 < valor 2 => ", valor1 < valor2)
	fmt.Println("valor1 > valor 2 => ", valor1 > valor2)
	fmt.Println("valor1 <= valor 2 => ", valor1 <= valor2)
	fmt.Println("valor1 >= valor 2 => ", valor1 >= valor2)

	fmt.Println("Operador AND", 1 > 0 && 2 > 0)
	fmt.Println("Operador OR", 2 < 0 || 1 > 0)
	myBool := true
	fmt.Println("Operador NOT", !myBool)
}
