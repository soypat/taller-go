//+build OMIT
package main

import (
	"fmt"
	"time"
)

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	go fmt.Println("Concurrente")
	fmt.Println("Común")
	// El sleep pausa la ejecución y permite que se ejecute la gorutina
	time.Sleep(time.Millisecond)
	// PROG_E OMIT
}

// MAIN_E OMIT
/* FUNC_S OMIT
go funcion(arg1, arg2)
*/ // FUNC_E OMIT
