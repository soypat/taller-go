//+build non-slide OMIT

package main

import "fmt"

// MAIN_S OMIT
func main() { // HLmain
	// PROG_S OMIT
	s := "世界"                  // HLhello
	fmt.Printf("Hello, %s", s) // HLhello
	// PROG_E OMIT
} // HLmain
// MAIN_E OMIT

/* FMT_S OMIT
%v    valor: formato nativo              %d    entero: base 10
%+v   valor: nativo c/ campos de struct  %b    entero: base 2
%#v   valor: sintaxis de Go              %f    float:  decimal sin exponente
%T    tipo:  sintaxis de Go              %e    float:  decimal con exponente
%%    Un simbolo porcentaje              %9.2f float:  decimal ancho 9, precision 2
*/ // FMT_E OMIT
