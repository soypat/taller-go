//+build non-slide OMIT

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string = "Hello, 世界"
	fmt.Println(strings.Split(s))
	os.Mkdir()
}
