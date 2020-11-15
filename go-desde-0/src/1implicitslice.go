//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	// Variable name length is the same to visualize the length of different declaration types OMIT
	// Larga
	var beatlesLongg [4]string = [4]string{"John", "Paul", "George", "Ringo"}
	// Corta
	var beatlesShort = [4]string{"John", "Paul", "George", "Ringo"}
	// Implicita
	beatImplicit := [4]string{"John", "Paul", "George", "Ringo"}
	// Implicita + longitud inferida
	beatInferred := [...]string{"John", "Paul", "George", "Ringo"}

	fmt.Println(beatlesLongg)
	fmt.Println(beatlesShort)
	fmt.Println(beatImplicit)
	fmt.Println(beatInferred)
	// PROG_E OMIT
}
