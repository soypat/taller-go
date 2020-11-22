//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	beatles := []string{"John", "Paul", "George", "Ringo"}
	beatleSlice := beatles[0:3]

	fmt.Println(beatles)
	beatleSlice[1] = "Roger"
	fmt.Println(beatles)
	// PROG_E OMIT
}
