//+build OMIT
package main

import (
	"fmt"
	"time"
)

// MAIN_S OMIT
func main() {
	// PROG_S OMIT
	go slowFunction("rapido", 2)
	slowFunction("lento", 5)
	// PROG_E OMIT
}

// MAIN_E OMIT
// FUNC_S OMIT
// Esta función tardará `jobsize` segundos en correr
func slowFunction(job string, jobsize int) {
	fmt.Printf("job %s [size:%d] started\n", job, jobsize)
	time.Sleep(time.Duration(jobsize) * time.Second)
	fmt.Printf("job %s finished\n", job)
}

// FUNC_E OMIT
