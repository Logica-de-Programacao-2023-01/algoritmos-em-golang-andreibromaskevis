package main

import "fmt"

func main() {

	var idade int
	fmt.Print("Digite sua idade em anos: ")
	fmt.Scanln(&idade)

	idade_dias := idade * 365

	fmt.Printf("%d anos = %d dias ", idade, idade_dias)

}
