//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
)

func main() {
	// DECL_S OMIT
	var s string
	// DECL_E OMIT
	// ASSIGN_S OMIT
	var s string = "Hello world"
	var s = "Hello World"
	s := "Hello World"
	// ASSIGN_E OMIT
	// USE_S OMIT
	s = "Hola Mundo"
	// USE_E OMIT
	fmt.Println(s)
}

// FUNC_S OMIT
func Sumar(x int, y int) int {
	return x + y
}

// FUNC_E OMIT
