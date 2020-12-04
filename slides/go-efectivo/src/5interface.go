//+build os OMIT

// ALL_S OMIT
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("data/archivo.txt") // f es un *os.File
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f) // ReadAll acepta a f como argumento
	if err != nil {
		panic(err)
	}
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
