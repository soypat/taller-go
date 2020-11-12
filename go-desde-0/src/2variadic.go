//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	s := Concatenar("Saludos", " al", " gran", " rey!")
	fmt.Println(s)
	// PROG_E OMIT
}

// FUNC_S OMIT
func Concatenar(palabras ...string) string {
	var oración string
	for _, palabra := range palabras {
		oración += palabra
	}
	return oración
}

// FUNC_E OMIT
