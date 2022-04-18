// ALL_S OMIT
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("data/archivo.txt") // f es un *os.File
	b, _ := io.ReadAll(f)               // ReadAll acepta a f como argumento
	fmt.Printf("%s", b)
}

// ALL_E OMIT

// INTERFACE_S OMIT
// Dentro del paquete io
type Reader interface {
	Read(p []byte) (n int, err error)
}

// INTERFACE_E OMIT

// READALL_S OMIT
// Dentro del paquete ioutil
func ReadAll(r io.Reader) ([]byte, error) {
	return nil, nil // OMIT
	// ...
}

// READALL_E OMIT

// THINGS_S OMIT
func init() {
	os.Mkdir("data", os.ModeDevice)
	f, err := os.Create("data/archivo.txt")
	if err != nil {
		panic(err)
	}
	f.Write([]byte(`Contengo datos muy importantes:

	creado: 14/11/2020
	modificado: 22/11/2020
	contenido: Saludos de parte del Seminario de Go
	autores: Patricio Whittingslow, Agustín Canalis
	
	fin`))
}

// THINGS_E OMIT
