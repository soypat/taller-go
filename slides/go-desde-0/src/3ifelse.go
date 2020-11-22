//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var b bool = true
	if b {
		fmt.Println("b es verdadero!")
	} else {
		fmt.Println("b es falso!")
	}
	// PROG_E OMIT
}
