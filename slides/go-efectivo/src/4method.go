//+build ignore OMIT

package main

import "fmt"

type Pelicula struct {
	Nombre  string
	Año     int
	Puntaje float32
}

// Primero va el receptor de método (p Pelicula)
func (p Pelicula) Resumen() string {
	fstr := "La película %v se estrenó en %v \nRecibió un puntaje de %0.1f/10"
	return fmt.Sprintf(fstr, p.Nombre, p.Año, p.Puntaje)
}

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	p := Pelicula{
		Nombre:  "El Secreto de sus Ojos",
		Año:     2009,
		Puntaje: 8.2,
	}
	fmt.Println(p.Resumen())
	// PROG_E OMIT
}

// MAIN_E OMIT
