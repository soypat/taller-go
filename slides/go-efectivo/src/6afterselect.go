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
	// Declaro un canal con un buffer de 1 mensaje
	c1, c2 := make(chan string), make(chan string)

	go Mensajero(500, c1)
	go Mensajero(2000, c2)

	for {
		select {
		case msj := <-c1:
			fmt.Println(msj)
		case msj := <-c2:
			fmt.Println(msj)
		}
	}
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT
func Mensajero(t int, c chan string) {
	for {
		c <- fmt.Sprintf("cada %vms", t)
		time.Sleep(time.Millisecond * time.Duration(t))
	}
}

// FUNC_E OMIT
