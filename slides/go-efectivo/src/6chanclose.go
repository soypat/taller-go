//+build ignore OMIT

package main

// IMPORT_S OMIT
import (
	"fmt"
	"time"
)

// IMPORT_E OMIT
// MAIN_S OMIT
// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	ch := make(chan string)
	go Contar("ovejas", ch)

	for {
		msj, abierto := <-ch
		if !abierto { // si el canal está cerrado salimos del for
			break
		}
		fmt.Println(msj)
	}
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT
func Contar(cosa string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- cosa
		time.Sleep(time.Millisecond * 500)
	}
	close(c) // Ahora cerramos el canal al terminar el trabajo!
}

// FUNC_E OMIT
