//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	suma := Sumar(1, 2)
	fmt.Println(suma)
	// PROG_E OMIT
}

// FUNC_S OMIT
func Sumar(x int, y int) int {
	return x + y
}

// FUNC_E OMIT

// FCFUNC_S OMIT
var función = func(a int) int {
	return a*20 + 5
}

// FCFUNC_E OMIT
