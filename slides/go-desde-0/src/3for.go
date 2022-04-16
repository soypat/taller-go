//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	for i := 0; i <= 10; i++ {
		fmt.Println("i es", i)
	}
	// PROG_E OMIT
}
