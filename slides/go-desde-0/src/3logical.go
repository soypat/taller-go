//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	parcial1, parcial2, final := 4, 6, 5

	if parcial1 >= 4 && parcial2 >= 4 {
		fmt.Print("Cursada aprobada y final ")
		if final >= 4 {
			fmt.Println("aprobado!")
		} else {
			fmt.Println("desaprobado :(")
		}

	} else if parcial1 >= 4 || parcial2 >= 4 {
		fmt.Println("Recupera parcial")

	} else {
		fmt.Println("Recursa")
	}
	// PROG_E OMIT
}

/* OPERATORS_S OMIT
Comparación  y Lógicos | Matemáticos       |  Bits               | Pointers
>   <   ==             |  +   -    *   /   |  &  |   ^  &^       |  *   &
>=  <=  !=             |  +=  -=   *=  /=  |  &= |=  ^=          |  Canales
||  &&   !             |  %        %=      |  <<  >>  <<=  >>=   |  <-  ->
*/ // OPERATORS_E OMIT
