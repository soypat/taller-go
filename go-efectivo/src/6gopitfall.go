//+build OMIT
package main

import "fmt"

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	go fmt.Println("Concurrente")
	fmt.Println("Común")
	// PROG_E OMIT
}

// MAIN_E OMIT
