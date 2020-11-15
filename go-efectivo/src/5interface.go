//+build ignore OMIT

package main

import "fmt"

// INTEFACE_S OMIT
type Lister interface {
	Push(string)
	Pop() string
}

// INTERFACE_E OMIT
// STRUCT_S OMIT
type StringList []string

func (a StringList) Push(s string) { a = append(a, s) }

func (a StringList) Pop() (x string) {
	if len(a) >= 1 {
		x, a = a[len(a)-1], a[:len(a)-1]
	}
	return
}

// STRUCT_E OMIT
// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	var s Lister
	s = StringList{"QWERTY"}
	s.Push("DVORAK")
	fmt.Println(s)
	// PROG_E OMIT
}

// MAIN_E OMIT
