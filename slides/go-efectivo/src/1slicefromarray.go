//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var beatles [4]string = [4]string{"John", "Paul", "George", "Ringo"}
	var beatleSlice []string = beatles[0:3]
	fmt.Println(beatleSlice)
	// PROG_E OMIT
}
