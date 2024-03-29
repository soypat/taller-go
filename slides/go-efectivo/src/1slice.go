//go:build ignore || OMIT
//go:build OMIT
// +build OMIT OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	// Declaración y Asignación
	beatles := []string{"John", "Paul", "George", "Ringo"}

	// Uso
	fmt.Println(beatles[0])
	fmt.Println(beatles[1:4])

	// Modificación
	beatles[0] = "Niki"
	beatles = append(beatles, "Nathy")
	fmt.Println(beatles)
	// PROG_E OMIT
}

/* IDX_S OMIT
beatles[0:5]
beatles[0:]
beatles[:5]
beatles[:]
beatles[:len(beatles)]
*/ // IDX_E OMIT
