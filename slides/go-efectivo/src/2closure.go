//go:build ignore || OMIT
//go:build OMIT
// +build OMIT OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	a := 0
	mas1 := func() {
		a++
	}

	mas1()
	mas1()
	fmt.Println(a)
	// PROG_E OMIT
}

// FUNC_S OMIT
func ClosureActualWork() {
	// FPROG_S OMIT

	// FPROG_E OMIT
}

// FUNC_E OMIT
