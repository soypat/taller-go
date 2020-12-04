//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	mas1 := func(n int) int {
		n++
		return n
	}
	b := mas1(2)
	fmt.Println(b)
	// PROG_E OMIT
}

// FUNC_S OMIT
func ClosureActualWork() {
	// FPROG_S OMIT
	b := "no, no sé. "
	mistring := "ya tu sabe"
	b += mistring
	fmt.Println(b)
	// FPROG_E OMIT
}

// FUNC_E OMIT
