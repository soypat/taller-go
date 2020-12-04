//+build ignore OMIT

package main

import (
	"fmt"
)

type Pelicula struct {
	Nombre  string
	Año     int
	Puntaje float32
}

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	p := new(Pelicula)
	fmt.Printf("%+v\n", p)
	p.Nombre, p.Puntaje = "Die Hard", 9.5
	fmt.Printf("%+v", p)
	// PROG_E OMIT
}

// MAIN_E OMIT
