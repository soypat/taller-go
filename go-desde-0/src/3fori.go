//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	for i := 1; i <= 10; i++ {
		fmt.Println("i es ", i)
	}
	// PROG_E OMIT
}

/* OPERATORS_S OMIT
Comparación  y Lógicos | Matemáticos       |  Bits               | Pointers
>   <   ==             |  +   -    *   /   |  &  |   ^  &^       |  *   &
>=  <=  !=             |  +=  -=   *=  /=  |  &= |=  ^=          |  Canales
||  &&   !             |  %        %=      |  <<  >>  <<=  >>=   |  <-  ->
*/ // OPERATORS_E OMIT
