//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var a = "papas🍟"
	if a == "pizza" {
		fmt.Println("a es pizza🍕")
	} else {
		fmt.Println("a no es pizza 🍕😣")
	}
	// PROG_E OMIT
}
