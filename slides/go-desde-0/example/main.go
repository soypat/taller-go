package main

import "fmt"

func main(){
	n:=Fibonacci(1001)
	fmt.Println(n)
}

func Fibonacci(n int) (int) {
	a,b := 1,0
	for i:=0; i<n; i++{
		c := a + b
		a = b
		b = c
		if b<0{
			fmt.Println("El resultado se hizo negativo 🙃")
			return 0
		}
	}
	return b
}