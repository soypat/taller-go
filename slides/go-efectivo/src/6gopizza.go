//+build ignore OMIT

package main

import (
	"fmt"
	"time"
)

// MAIN_S OMIT
func main() {
	puerta := make(chan string)
	go Delivery(puerta)
	pedido := <-puerta
	fmt.Println(pedido)
}

// MAIN_E OMIT
// FUNC_S OMIT
func Delivery(c chan string) {
	fmt.Println("Haciendo la pizza")
	time.Sleep(time.Second * 2)
	c <- "🍕"
}

// FUNC_E OMIT
