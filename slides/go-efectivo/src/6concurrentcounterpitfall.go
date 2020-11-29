//+build ignore OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// PROG_S OMIT
	go contar("ovejas")
	go contar("calamares")
	// PROG_E OMIT
}

// FUNC_S OMIT
func contar(cosa string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, cosa)
		time.Sleep(time.Millisecond * 500)
	}
}

// FUNC_E OMIT