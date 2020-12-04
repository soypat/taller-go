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
	fstr := "La película %v se estrenó en %v y recibió un puntaje de %0.1f/10"
	return fmt.Sprintf(fstr, p.Nombre, p.Año, p.Puntaje)
}

// REFMETHOD_S OMIT
func (p *Pelicula) SetNombre(name string) {
	p.Nombre = name
}

//REFMETHOD_E OMIT
// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	p := Pelicula{Nombre: "Un Cuento Chino",Año: 2011, Puntaje: 7.3}
	fmt.Println(p.Resumen())
	p.SetNombre("El robo del siglo")
	fmt.Println(p.Resumen())
	// PROG_E OMIT
}

// MAIN_E OMIT
