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
	// Declaración y asignación 
	p := Película{ 
		Nombre: "Die Hard",
		Puntaje: 9.5
		} 

	// Uso
	fmt.Print( p.Nombre, p.Puntaje) 
	
	// Modificación
	p.Nombre = "Die Hard 4"
    p.Puntaje = 7.9 
	// PROG_E OMIT
}
// MAIN_E OMIT

func DeclarationExample() {
	// DECL_S OMIT
	// Forma común
	a := Película{Nombre: "Annihilation", Puntaje: 8.1}
	// Para más legibilidad separado en lineas
	s := Película{
		Nombre:  "El Secreto de sus Ojos",
		Puntaje: 9.0,
	}
	// DECL_E OMIT
	fmt.Println(a, s)
}
