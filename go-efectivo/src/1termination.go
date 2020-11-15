//+build non-slide OMIT

package main

import (
	"fmt"
	"os"
)

// MAIN_S OMIT
func main() { // HLmain
	// PROG_S OMIT
	defer fmt.Println("Terminando el programa!")
	fmt.Println("Comenzando el programa!")
	// En contexto de main() termina ejecución llamando el defer stack de main()
	return
	// panic ejecuta el defer stack y devuelve control a la función circundante
	// esto ocurre recursivamente hasta terminar el programa o si se encuentra
	// con una llamada a recover()
	panic("Oh no! Un Error!")
	// Termina inmediatamente.  0: Todo OK  1-255: Hubo un error
	os.Exit(0)
	// PROG_E OMIT
} // HLmain
// MAIN_E OMIT
