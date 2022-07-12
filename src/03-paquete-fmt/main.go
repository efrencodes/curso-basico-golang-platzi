package main

import "fmt"

func main () {
	helloMessage := "Hello"
	worldMessage := "World"
	fmt.Println(helloMessage, worldMessage)

	nombre := "efren"
	edad := 29
	// Para imprimir un tipo de dato string => %s
	// Para imprimir un tipo de dato entero => %d
	fmt.Printf("Soy %s y tengo %d años\n", nombre, edad)
	// Cuando no sabes el tipo de dato se imprimi de esta manera => %v
	fmt.Printf("Soy %v y tengo %v años\n", nombre, edad)

	// Guarda un string para despues imprimirla
	message := fmt.Sprintf("Soy %s y tengo %d años", nombre, edad)
	fmt.Println(message)

	// Saber el tipo de dato
	fmt.Printf("nombre %T\n", nombre)
	fmt.Printf("edad %T\n", edad)
}