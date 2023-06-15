package main

import "fmt"

func main() {

	var produto float64

	fmt.Print("Digite o valor do produto:")
	fmt.Scanln(&produto)

	desconto := produto * 0.10
	precofinal := produto - desconto

	fmt.Print("Preço do produto: ", precofinal)
}
