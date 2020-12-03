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
func Hacer(cosa string, c chan string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 500)
		str := fmt.Sprintf("%d %s", i, cosa)
		c <- str
	}
	close(c) // Ahora cerramos el canal al terminar el trabajo!
}

// FUNC_E OMIT

// FUNC_DUMMY_S OMIT
func dummy(c chan string ){
	// ASK_S OMIT
	str, abierto <- c
	// ASK_E OMIT
	// CLOSE_S OMIT
	close(c)
	// CLOSE_E OMIT
}
// FUNC_DUMMY_E OMIT