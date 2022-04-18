//go:build OMIT

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// PROG_S OMIT
	fp, _ := os.Open("data/archivo.txt")
	defer fp.Close()
	bytes, _ := io.ReadAll(fp)
	fmt.Println(string(bytes))
	// PROG_E OMIT
}
