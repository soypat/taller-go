//go:build ignore || OMIT
//go:build OMIT
//go:build OMIT
// +build OMIT
// +build OMIT OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	miDiv, miResta := DivResta(10, 4)
	fmt.Println(miDiv, miResta)
	// PROG_E OMIT
}

// FUNC_S OMIT
func DivResta(x int, y int) (div int, resta int) {
	div = x / y
	resta = x % y
	return
}

// FUNC_E OMIT
