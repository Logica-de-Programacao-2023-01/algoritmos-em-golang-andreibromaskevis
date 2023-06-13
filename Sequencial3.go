package main

import "fmt"

func main() {

	var peso, altura float64
	fmt.Print("Digite seu peso: ")
	fmt.Scanln(&peso)
	fmt.Print("Digite sua altura: ")
	fmt.Scanln(&altura)

	imc := peso / (altura * altura)

	fmt.Print("IMC:", imc)

}
