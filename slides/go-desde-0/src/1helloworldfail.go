//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// START OMIT
	var s string
	s = "Hello, 世界"
	s = 1
	fmt.Println(s)
	// END OMIT
}
