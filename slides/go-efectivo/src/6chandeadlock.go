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
		msj := <-ch
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
}

// FUNC_E OMIT
