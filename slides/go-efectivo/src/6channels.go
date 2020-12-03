//+build ignore OMIT

package main

func main() {
	// CHANMAKE_S OMIT
	ch := make(chan string)
	// CHANMAKE_E OMIT
	// CHANSEND_S OMIT
	ch <- "pizza🍕"   
	// CHANSEND_E OMIT
	// CHANREC_S OMIT
	s := <-ch // ahora s vale "pizza🍕"
	// CHANREC_E OMIT
}