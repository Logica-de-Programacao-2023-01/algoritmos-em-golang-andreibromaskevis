package main

import "fmt"

func main() {

	var num1, num2, num3 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&num3)

	soma := (num1 + num2 + num3)

	fmt.Print("Soma dos números: ", soma)

}
