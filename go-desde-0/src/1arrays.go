//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var beatles [4]string
	beatles[0] = "John"
	beatles[1] = "Paul"
	beatles[2] = "George"
	beatles[3] = "Ringo"
	fmt.Println(beatles)
	// PROG_E OMIT
}
