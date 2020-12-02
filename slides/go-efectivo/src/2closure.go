//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	a := "no, no sé. "
	mistring := "ya tu sabe"
	miFuncion := func(arg1 string) string {
		a += arg1
		return a
	}
	fmt.Println(miFuncion(mistring))
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
