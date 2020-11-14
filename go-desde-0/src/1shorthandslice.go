//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var beatles = [4]string{"John", "Paul", "George", "Ringo"}
	var beatleSlice = beatles[0:3]
	fmt.Println(beatleSlice)
	beatleSlice[1] = "Roger"
	fmt.Println(beatles)
	// PROG_E OMIT
}
