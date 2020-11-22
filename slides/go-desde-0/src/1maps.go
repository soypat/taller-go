//+build ignore OMIT

package main

import "fmt"

func main() {
	// PROG_S OMIT
	var memes = make(map[string]int)
	memes["carpincho"] = 2020
	memes["over 9000"] = 2006
	memes["doge"] = 2005
	memes["loss"] = 2008
	fmt.Println(memes["doge"])
	fmt.Println(memes)
	// PROG_E OMIT
	// DEL_S OMIT
	delete(memes, "loss")
	// DEL_E OMIT
	// OK_S OMIT
	año, ok := memes["not bad"]
	// OK_E OMIT
	// so the code compiles:
	año++
	ok = !ok
}
