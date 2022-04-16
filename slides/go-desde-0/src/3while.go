//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	i := 0
	for i <= 10 {
		fmt.Println("i es", i)
		i++
	}
	// PROG_E OMIT
}
