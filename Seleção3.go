package main

import "fmt"

func main() {

	var num int
	fmt.Printf("digite um número: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("%d é par", num)
	} else {
		fmt.Printf("%d é ímpar", num)
	}
}
