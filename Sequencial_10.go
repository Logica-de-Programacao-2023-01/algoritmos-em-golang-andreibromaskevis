package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("Digite seu peso em quilos para a conversão :")
	fmt.Scanln(&peso)

	pesoLb := peso * 2.2046

	fmt.Print("Peso em libra: ", pesoLb)
}
