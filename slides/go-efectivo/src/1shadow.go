//+build OMIT
package main

import "fmt"

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	yo := "Ismael"
	{
		yo := "Fulano" // pruebe borrar el ":" y ver como cambia
		fmt.Println(yo)
	}
	fmt.Println(yo)
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT
// FUNC_E OMIT
