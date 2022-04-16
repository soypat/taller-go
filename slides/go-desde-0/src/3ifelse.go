//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	a := "pizza"
	if a == "pizza" {
		fmt.Println("a es pizza 🍕")
	}
	// PROG_E OMIT
}
