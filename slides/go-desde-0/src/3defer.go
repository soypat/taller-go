//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"os"
)

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	defer fmt.Println("Fin de programa!")
	defer fmt.Println("Corro después que termina la función circundante")
	fmt.Println("Comienzo de programa!")
	// PROG_E OMIT
}

// MAIN_E OMIT

func open() {
	// EXAMPLE_S OMIT
	archivo, err := os.Open("miArchivo.txt")
	defer archivo.Close()
	// EXAMPLE_E OMIT
	fmt.Println(err.Error())
}
