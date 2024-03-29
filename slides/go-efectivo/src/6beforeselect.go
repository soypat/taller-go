//go:build ignore || OMIT
//go:build OMIT
// +build OMIT OMIT

package main

// IMPORT_S OMIT
import (
	"fmt"
	"log" // OMIT
	"time"
)

// IMPORT_E OMIT

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	c1 := make(chan string)
	c2 := make(chan string)

	go PizzeroRapido(c1) // pizza cada 500ms
	go PanaderoLento(c2) // pizza cada 2000ms

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
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

func PanaderoLento(c chan string) {
	t := 2000
	for {
		c <- fmt.Sprintf("🍞 cada %vms", t)
		time.Sleep(time.Millisecond * time.Duration(t))
	}
}

// FUNC_E OMIT

// TIMEOUT_S OMIT

func init() {
	go main()
	time.Sleep(30 * time.Second)
	log.Fatal("timeout after 30 seconds")
}

// TIMEOUT_E OMIT
