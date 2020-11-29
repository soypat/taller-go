//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	// Declaración y Asignación
	var beatles = []string{"John", "Paul", "George", "Ringo"} 

	// Uso
	fmt.Println(beatles[0], beatles[1:4]) 

	// Modificación
	beatles[0] = "Palito Ortega" 
	beatles = append(beatles, "Walrus")
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
