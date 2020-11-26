//+build ignore OMIT

package main

import "fmt"

type Juego struct {
	Nombre  string
	Puntaje float32
	Año     uint16
}

func (j Juego) Resumen() string { return fmt.Sprintf("%v [%v] %0.1f/10", j.Nombre, j.Año, j.Puntaje) }

func (j *Juego) NuevoPuntaje(p float32) {
	j.Puntaje = p
}

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	p := Juego{"Counter Strike", 7.5, 1999}
	fmt.Println(p.Resumen())
	p.NuevoPuntaje(9)
	fmt.Println(p.Resumen())
	// PROG_E OMIT
}

// MAIN_E OMIT
