//+build ignore OMIT

package main

import (
	"fmt"
)

func main() {
	// PROG_S OMIT
	var s = "Hola, world"
	s += "!!!!!!"
	var i = 20
	var f = 30.2
	fmt.Println(i, f, s)
	// PROG_E OMIT
}

// FUNC_S OMIT
func Sumar(x int, y int) int {
	return x + y
}

// FUNC_E OMIT