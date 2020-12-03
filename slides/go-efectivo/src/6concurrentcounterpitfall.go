//+build ignore OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// PROG_S OMIT
	// GOLINE_S OMIT
	go Contar("ovejas") 
	// GOLINE_E OMIT
	Contar("calamares")
	// PROG_E OMIT
}

// FUNC_S OMIT
func Contar(cosa string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, cosa)
		time.Sleep(time.Millisecond * 500)
	}
}

// FUNC_E OMIT
