//+build OMIT
package main

import "fmt"

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	yo := "Ismael"
	{
		tu := "Fulano"
		fmt.Println(yo, tu)
	}
	fmt.Println(yo) // Se puede imprimir la variable "tu" aquí?
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT
// FUNC_E OMIT
