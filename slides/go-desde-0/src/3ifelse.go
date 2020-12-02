//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var a = "pizza" 
	if a=="pizza" {
		fmt.Println("a es pizza 🍕")
	}
	// PROG_E OMIT
}
