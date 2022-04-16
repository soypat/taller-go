//go:build OMIT
// +build OMIT

package main

import "fmt"

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	a := 2
	b := 4
	c := a + b
	fmt.Println(c)
	// PROG_E OMIT
}

// MAIN_E OMIT
