//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	nota := 5
	if nota < 4 {
		fmt.Println("Desaprobado")
	} else if nota == 10 {
		fmt.Println("Excelente! Un diez!")
	} else if nota >= 4 {
		fmt.Println("Aprobado!")
	} else {
		fmt.Println("Qué?")
	}
	// PROG_E OMIT
}
