//+build ignore OMIT

package main

import (
	"fmt"
	"time"
)

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	waitPizza, waitPan := true, true

	go func() {
		Hacer("pan🍞")
		waitPan = false
	}()

	go func() {
		Hacer("pizza🍕") 
		waitPizza = false
	}()

	for waitPan || waitPizza {
		// solo espera
	}
	// PROG_E OMIT
}

// MAIN_E OMIT

// FUNC_S OMIT
func Hacer(cosa string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, cosa)
		time.Sleep(time.Millisecond * 500)
	}
}
// FUNC_E OMIT
