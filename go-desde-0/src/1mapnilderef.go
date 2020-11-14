//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var beatles map[string]string
	beatles["drummer"] = "ringo"
	fmt.Println(beatles)
	// PROG_E OMIT
}
