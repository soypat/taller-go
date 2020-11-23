//+build ignore OMIT

package main

// IMPORT_S OMIT
import (
	"fmt"
	"time"
)

// IMPORT_E OMIT
// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	ch := make(chan string)
	go contar("ovejas", ch)

	for msj := range ch { // Deja de iterar cuando se cierra el canal
		fmt.Println(msj)
	}
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT
func contar(cosa string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- cosa
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

// FUNC_E OMIT