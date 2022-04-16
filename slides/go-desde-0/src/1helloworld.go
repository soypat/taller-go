//go:build OMIT
// +build OMIT

package main

import "fmt"

// MAIN_S OMIT
func main() { // HLmain
	// START OMIT
	var s string // HLhello
	s = "Hello, 世界"
	fmt.Println(s) // HLhello
	// END OMIT
} // HLmain
// MAIN_E OMIT
