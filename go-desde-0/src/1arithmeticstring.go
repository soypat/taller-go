//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var greeting string
	greeting = "Hola "
	greeting = greeting + "Roberto Bolaño"
	fmt.Println(greeting)
	// PROG_E OMIT
}
