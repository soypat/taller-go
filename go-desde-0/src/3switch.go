//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	i := 2

	switch i {
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	case 4:
		fmt.Println("Cuatro")
	default:
		fmt.Println("No conozco ese número")
	}
	// PROG_E OMIT
}

/* OPERATORS_S OMIT
Comparación  y Lógicos | Matemáticos       |  Bits               | Pointers
>   <   ==             |  +   -    *   /   |  &  |   ^  &^       |  *   &
>=  <=  !=             |  +=  -=   *=  /=  |  &= |=  ^=          |  Canales
||  &&   !             |  %        %=      |  <<  >>  <<=  >>=   |  <-  ->
*/ // OPERATORS_E OMIT
