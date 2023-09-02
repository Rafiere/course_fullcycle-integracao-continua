package main

import "fmt"

func main() {

	fmt.Println("Hi!")

	fmt.Println(Soma(10, 10))
}

func Soma(a int, b int) int {
	return a + b
}
