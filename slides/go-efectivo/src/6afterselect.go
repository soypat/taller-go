//+build ignore OMIT

package main

// IMPORT_S OMIT
import (
	"fmt"
	"log" // OMIT
	"time"
)

// IMPORT_E OMIT
// MAIN_S OMIT
// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	c1, c2 := make(chan string), make(chan string)

	go PizzeroRapido(c1)
	go PizzeroLento(c2)

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
func PizzeroRapido(c chan string) {
	t := 500
	for {
		c <- fmt.Sprintf("🍕 cada %vms", t)
		time.Sleep(time.Millisecond * time.Duration(t))
	}
}
func PizzeroLento(c chan string) {
	t := 2000
	for {
		c <- fmt.Sprintf("🍕 cada %vms", t)
		time.Sleep(time.Millisecond * time.Duration(t))
	}
}

// FUNC_E OMIT

// FUNC_E OMIT
// TIMEOUT_S OMIT

func init() {
	go main()
	time.Sleep(30 * time.Second)
	log.Fatal("timeout after 30 seconds")
}

// TIMEOUT_E OMIT
