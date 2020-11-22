//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var soyUnSlice []string
	// ahora sin errores
	soyUnSlice = append(soyUnSlice, "hola!")
	longitud := len(soyUnSlice)
	fmt.Println(soyUnSlice)
	fmt.Println(len(longitud))
	// PROG_E OMIT
	// should error!
}
