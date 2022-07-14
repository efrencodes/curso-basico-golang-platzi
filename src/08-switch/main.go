package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	switch today.Weekday() {
	case 1:
		fmt.Println("Hoy es Lunes")
	case 2:
		fmt.Println("Hoy es Martes")
	case 3:
		fmt.Println("Hoy es Miercoles.")
	}

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es un número par")
	default:
		fmt.Println("Es un número impar")
	}

	value := 12
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No hay condición")
	}
}
