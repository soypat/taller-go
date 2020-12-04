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
	go Hacer("pizza🍕", ch)
	for {
		msj := <-ch
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
}

// FUNC_E OMIT
