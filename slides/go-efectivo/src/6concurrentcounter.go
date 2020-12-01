//+build ignore OMIT

package main

import (
	"fmt"
	"time"
)

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	waitOvejas, waitCalam := true, true

	go func() {
		Contar("ovejas")
		waitOvejas = false
	}()

	go func() {
		Contar("calamares")
		waitCalam = false
	}()

	for waitOvejas || waitCalam {
		// solo espera
	}
	// PROG_E OMIT
}

// MAIN_E OMIT

// FUNC_S OMIT
func Contar(cosa string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, cosa)
		time.Sleep(time.Millisecond * 500)
	}
}

// FUNC_E OMIT
