//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	mas1 := func(n int) int {
		n++
		return n
	}
	a := mas1(2)
	fmt.Println(a)
	// PROG_E OMIT
}

// FUNC_S OMIT
func ClosureActualWork() {
	// FPROG_S OMIT
	a := "no, no sé. "
	mistring := "ya tu sabe"
	a += mistring
	fmt.Println(a)
	// FPROG_E OMIT
}

// FUNC_E OMIT
