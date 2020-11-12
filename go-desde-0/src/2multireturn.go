//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	miDiv, miResta := DivResta(10, 4)
	fmt.Println(miDiv, miResta)
	// PROG_E OMIT
}

// FUNC_S OMIT
func DivResta(x int, y int) (float64, int) {
	div := float64(x) / float64(y)
	resta := x % y
	return div, resta
}

// FUNC_E OMIT
