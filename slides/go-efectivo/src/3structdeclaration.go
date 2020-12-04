//+build ignore OMIT

package main

import (
	"fmt"
)

type Pelicula struct {
	Nombre  string
	Año int
	Puntaje float32
}

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	// Declaración y asignación
	p := Pelicula{
		Nombre:  "El Secreto de sus Ojos",
		Año: 2009,
		Puntaje: 8.2,
	}
	// Uso
	fmt.Println(p.Nombre, p.Puntaje)
	// Modificación
	p.Nombre = "Die Hard 4"
	p.Puntaje = 7.9
	// PROG_E OMIT
}

// MAIN_E OMIT
func DeclarationExample() {
	// DECL_S OMIT
	// Forma común
	p1 := Pelicula{Nombre: "Un Cuento Chino",Año: 2011, Puntaje: 7.3}
	// Para más legibilidad separado en lineas
	p2 := Pelicula{
		Nombre:  "El Secreto de sus Ojos",
		Año: 2009,
		Puntaje: 8.2,
	}
	// DECL_E OMIT
	fmt.Println(p1, p2)
}
