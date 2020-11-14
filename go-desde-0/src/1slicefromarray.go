//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	beatles := [...]string{"John", "Paul", "George", "Ringo"}
	var beatleSlice []string = beatles[0:3]
	fmt.Println(beatleSlice)
	// PROG_E OMIT
}
