//+build ignore OMIT

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// PROG_S OMIT
	f, err := os.Open("archivo.txt")
	if err != nil {
		fmt.Println("No pude abrir el archivo:", err)
		// termina el programa indicandole al sistema operativo que hubo un error
		os.Exit(1)
	}
	buff, _ := ioutil.ReadAll(f)
	fmt.Println(string(buff), "\nArchivo leido!")
	// PROG_E OMIT
}
