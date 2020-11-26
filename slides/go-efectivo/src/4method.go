//+build ignore OMIT

package main

import "fmt"

type Juego struct {
	Nombre  string
	Puntaje float32
	Año     uint16
}

// Primero va el receptor de método (j Juego)
func (j Juego) Resumen() string {
	return fmt.Sprintf("%v [%v] %0.1f/10",
		j.Nombre, j.Año, j.Puntaje)
}

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	p := Juego{"Dark Souls", 8.6, 2011}
	fmt.Println(p.Resumen())
	// PROG_E OMIT
}

// MAIN_E OMIT
