//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	números := [5]int{1, 5, 4, 2, 3}
	for i, n := range números {
		fmt.Println("el número", n, " está en posición", i)
	}
	// PROG_E OMIT
}
