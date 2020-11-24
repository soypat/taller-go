//+build ignore OMIT
// PROG_S OMIT
package main

import "fmt"

func Sumar(x int, y int) int {
	return x + y
}
// PROG_E OMIT

// FUNC_S OMIT
func main() {
	suma := Sumar(1, 2)
	fmt.Println(suma)
}
// FUNC_E OMIT