//+build ignore OMIT

package main

func main() {
	// PROG_S OMIT

	// PROG_E OMIT
}

/* CHAN_S OMIT
ch <- v    // Manda v al canal ch.
v := <-ch  // Recibe de ch
		   // y asigna el valor a v.

// creación de un canal de tipo int
ch := make(chan int)
*/ // CHAN_E OMIT
