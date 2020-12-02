//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	a := 0
	suma1 := func(n int) int {
		a += n
		return a
	}
	b := suma1(a)
	fmt.Println(b)
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
