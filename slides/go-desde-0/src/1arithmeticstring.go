//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	saludo := "Hola "
	nombre := "Roberto Bolaño"
	saludo = saludo + nombre
	fmt.Println(saludo)
	// PROG_E OMIT
}
