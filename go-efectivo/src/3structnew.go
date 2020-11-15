//+build ignore OMIT

package main

import (
	"fmt"
)

type Película struct {
	Nombre  string
	Puntaje float32
}

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	p := new(Película)
	fmt.Printf("%+v\n", p)
	p.Nombre, p.Puntaje = "Die Hard", 9.5
	fmt.Printf("%+v", p)
	// PROG_E OMIT
}

// MAIN_E OMIT
