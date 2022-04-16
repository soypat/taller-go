//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	i := 2

	switch i {
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	case 4, 5:
		fmt.Println("Cuatro o cinco")
	default:
		fmt.Println("No conozco ese número")
	}
	// PROG_E OMIT
}
