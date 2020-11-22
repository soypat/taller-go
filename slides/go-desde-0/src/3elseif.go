//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var nota int = 5
	if nota < 4 {
		fmt.Println("Desaprobado")
	} else if nota == 10 {
		fmt.Println("Excelente! Un diez!")
	} else if nota >= 4 {
		fmt.Println("Aprobado!")
	}
	// PROG_E OMIT
}

/* OPERATORS_S OMIT
Comparación  y Lógicos | Matemáticos       |  Bits               | Pointers
>   <   ==             |  +   -    *   /   |  &  |   ^  &^       |  *   &
>=  <=  !=             |  +=  -=   *=  /=  |  &= |=  ^=          |  Canales
||  &&   !             |  %        %=      |  <<  >>  <<=  >>=   |  <-  ->
*/ // OPERATORS_E OMIT
