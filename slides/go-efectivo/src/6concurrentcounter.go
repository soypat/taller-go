//+build ignore OMIT

package main

import (
	"fmt"
	"time"
)

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	wait := true

	go func() {
		contar("ovejas")
		wait = false
	}()

	go contar("calamares")
	for wait { // solo espera
	}
	// PROG_E OMIT
}

// MAIN_E OMIT

// FUNC_S OMIT
func contar(cosa string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, cosa)
		time.Sleep(time.Millisecond * 500)
	}
}

// FUNC_E OMIT
