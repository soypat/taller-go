//+build ignore OMIT

package main

// IMPORT_S OMIT
import (
	"fmt"
)

// IMPORT_E OMIT
// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	// Declaro un canal con un buffer de 1 mensaje
	ch := make(chan string, 1)
	ch <- "📦"

	s := <-ch
	fmt.Println(s)
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT

// FUNC_E OMIT
