//+build os OMIT

package main

import (
	"fmt"
	"sort"
)

// INTERFACE_S OMIT
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// INTERFACE_E OMIT

// USG_S OMIT
type Names []string

func (a Names) Len() int           { return len(a) }
func (a Names) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Names) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var nombres sort.Interface = Names{"Paco", "Yorsch", "Quique", "Cana", "Pato"}
	fmt.Println(nombres)
	sort.Sort(nombres)
	fmt.Println(nombres)
}

// USG_E OMIT

// SORTER_S OMIT
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// SORTER_E OMIT

// SORTSIGNATURE_S OMIT
func Sort(data Interface) {
	// el código que ordena va aqúi
}

// SORTSIGNATURE_E OMIT
