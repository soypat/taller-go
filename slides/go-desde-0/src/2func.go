//go:build OMIT
// +build OMIT

// PROG_S OMIT
package main

import "fmt"

func Sumar(x int, y int) int {
	return x + y
}

// PROG_E OMIT

// FUNC_S OMIT
func main() {
	suma1 := Sumar(1, 2)
	suma2 := Sumar(30, 41)
	fmt.Println(suma1, suma2)
}

// FUNC_E OMIT
