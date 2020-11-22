//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// PROG_E OMIT
}
