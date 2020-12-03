//+build ignore OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// PROG_S OMIT
	// GOLINE_S OMIT
	go Hacer("pizza🍕") 
	// GOLINE_E OMIT
	Hacer("pan🍞")
	// PROG_E OMIT
}

// FUNC_S OMIT
func Hacer(cosa string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, cosa)
		time.Sleep(time.Millisecond * 500)
	}
}

// FUNC_E OMIT
