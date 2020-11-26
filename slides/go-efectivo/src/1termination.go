//+build non-slide OMIT

package main

import (
	"fmt"
	"os"
)

// MAIN_S OMIT
func main() { // HLmain
	// PROG_S OMIT
	// Se usa la libreria os para escribir a stdout para que si el usuario comenta os.Exit, no crashee OMIT
	defer fmt.Println("Terminando el programa!")
	fmt.Println("Comenzando el programa!")

	// En contexto de main() termina ejecución llamando el defer stack de main()
	return

	// panic ejecuta el defer stack y termina todas las funciones
	panic("Oh no! Un Error!")

	// Termina inmediatamente.  0: Todo OK  1-255: Hubo un error
	os.Exit(0)

	// Llegar al final de main() == return
	fmt.Println("Se llegó al final del programa!")
	// PROG_E OMIT
} // HLmain
// MAIN_E OMIT
