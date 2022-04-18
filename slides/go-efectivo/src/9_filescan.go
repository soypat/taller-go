//go:build OMIT

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// PROG_S OMIT
	fp, _ := os.Open("data/archivo.txt")
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("linea:", line)
	}
	// PROG_E OMIT
}
