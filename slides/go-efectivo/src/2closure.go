//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	b := 0
	mas1 := func() {
		b++
	}
	mas1()
	mas1()
	fmt.Println(b)
	// PROG_E OMIT
}

// FUNC_S OMIT
func ClosureActualWork() {
	// FPROG_S OMIT

	// FPROG_E OMIT
}

// FUNC_E OMIT
