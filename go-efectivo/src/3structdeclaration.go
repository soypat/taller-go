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
	var p Película
	fmt.Printf("%+v\n", p)
	p.Nombre, p.Puntaje = "Die Hard", 9.5
	fmt.Printf("%+v", p)
	// PROG_E OMIT
}

// MAIN_E OMIT

func DeclarationExample() {
	// DECL_S OMIT
	// Forma común
	a := Película{Name: "Annihilation", Puntaje: 8.1}
	// Campos en el orden que es definido el tipo
	// Recomendable solo para structs con pocos campos
	r := Película{"The Room", 10.0}
	// Para más legibilidad separado en lineas
	s := Película{
		Name:    "El Secreto de sus Ojos",
		Puntaje: 9.0,
	}
	// DECL_E OMIT
	fmt.Println(a, r, s)
}
