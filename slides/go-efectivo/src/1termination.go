//+build non-slide OMIT

package main

import (
	"os"
)

// MAIN_S OMIT
func main() { // HLmain
	// PROG_S OMIT
	// Se usa la libreria os para escribir a stdout para que si el usuario comenta os.Exit, no crashee OMIT
	defer os.Stdout.WriteString("Terminando el programa!\n")
	os.Stdout.WriteString("Comenzando el programa!\n")
	// En contexto de main() termina ejecución llamando el defer stack de main()
	return
	// panic ejecuta el defer stack y devuelve control a la función circundante
	// esto ocurre recursivamente hasta terminar el programa o si se encuentra
	// con una llamada a recover()
	panic("Oh no! Un Error!")
	// Termina inmediatamente.  0: Todo OK  1-255: Hubo un error
	os.Exit(0)
	// Llegar al final de main() == return
	os.Stderr.WriteString("Se llegó al final del programa!\n")
	// PROG_E OMIT
} // HLmain
// MAIN_E OMIT
