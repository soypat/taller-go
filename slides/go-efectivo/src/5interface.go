package main

import "fmt"

// INTEFACE_S OMIT
type (
	Stacker interface {
		Push(string)
		Pop() string
	}
	Stack []string
)

// INTERFACE_E OMIT
func (a *Stack) Push(s string) { *a = append(*a, s) }
func (a *Stack) Pop() (x string) {
	if len(*a) >= 1 {
		x, *a = (*a)[len(*a)-1], (*a)[:len(*a)-1]
	}
	return
}

// STRUCT_E OMIT
// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	var s Stacker = &Stack{"QWERTY"}
	s.Push("DVORAK")
	fmt.Println(s)
	fmt.Println(s.Pop(), s.Pop())
	fmt.Println(s)
	// PROG_E OMIT
}

// MAIN_E OMIT
