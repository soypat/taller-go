//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var creación, días, presente int
	creación, presente = 2009, 2020
	días = (presente - creación) * 365
	fmt.Println(días)
	// PROG_E OMIT
}
