//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	pedido := func() string {
		return "🍕"
	}( /*0 args*/ )

	fmt.Println(pedido)
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
