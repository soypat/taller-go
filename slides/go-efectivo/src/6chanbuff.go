//+build ignore OMIT

package main

// IMPORT_S OMIT
import (
	"fmt"
)

// IMPORT_E OMIT
// MAIN_S OMIT
// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	// MAKE_S OMIT
	ch := make(chan string, 2)
	// MAKE_E OMIT
	ch <- "🏀"
	ch <- "🏀"
	b1 := <-ch
	b2 := <-ch
	fmt.Println(b1, b2)
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT

// FUNC_E OMIT
