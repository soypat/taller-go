//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	// Declaración y Asignación
	memes := map[string]int{"carpincho": 2020, "over 9000": 2006, "doge": 2005}

	// Uso
	fmt.Println(memes)
	fmt.Println(memes["carpincho"])
	año, ok := memes["not here"]
	fmt.Println(año, ok)

	// Modificación
	memes["loss"] = 2008
	delete(memes, "carpincho")
	// PROG_E OMIT
}
