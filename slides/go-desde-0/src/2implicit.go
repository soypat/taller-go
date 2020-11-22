//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	suma := 2 + 2
	suma++
	nombre := "Annie"
	nombre += " are you ok?"
	fmt.Println(suma)
	fmt.Println(nombre)
	// PROG_E OMIT
}

// FUNC_S OMIT
func Sumar(x int, y int) int {
	return x + y
}

// FUNC_E OMIT
