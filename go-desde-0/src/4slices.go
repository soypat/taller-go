//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	// el arreglo de siempre
	var soyUnArreglo [3]string
	// Oh! Un slice!
	var soyUnSlice []string

	soyUnSlice[0] = "hola!"
	soyUnArreglo[0] = "chau!"
	fmt.Println(soyUnSlice, soyUnArreglo)
	// PROG_E OMIT
	// should error!
}
