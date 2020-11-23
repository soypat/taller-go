//+build ignore OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		contar("ovejas")
		wg.Done()
	}()

	contar("calamares")
	wg.Wait()
	// PROG_E OMIT
}

// MAIN_E OMIT

// FUNC_S OMIT
func contar(cosa string) {
	for i := 0; i <= 7; i++ {
		fmt.Println(i, cosa)
		time.Sleep(time.Millisecond * 500)
	}
}

// FUNC_E OMIT
