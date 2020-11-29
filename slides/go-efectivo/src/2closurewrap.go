//+build ignore OMIT

package main

import (
	"fmt"
)

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	fmt.Println(AlgunaFuncion())
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT
func AlgunaFuncion() (mainError error) {
	// FPROG_S OMIT
	defer func() {
		if mainError != nil {
			mainError = fmt.Errorf("in AlgunaFuncion: %w", mainError)
		}
	}()
	mainError = ImGonnaError("too bad")
	if mainError != nil {
		return mainError
	}
	return nil
	// FPROG_E OMIT
}

// FUNC_E OMIT
func ImGonnaError(s string) error {
	return fmt.Errorf("ImGonnaError did not expect %q", s)
}
