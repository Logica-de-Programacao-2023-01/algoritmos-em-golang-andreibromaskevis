package main

import "fmt"

func main() {

	var num1 int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&num1)
	var dobro int = (num1 * 2)
	var triplo int = (num1 * 3)
	var quadruplo int = (num1 * 4)
	fmt.Println("Dobro:", dobro)
	fmt.Println("Triplo:", triplo)
	fmt.Println("Quadruplo:", quadruplo)

}
