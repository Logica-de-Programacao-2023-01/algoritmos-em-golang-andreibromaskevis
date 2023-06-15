package main

import "fmt"

func main() {

	var altura float64
	var pesoideal float64
	var sexo string
	fmt.Println("Digite sua altura: ")
	fmt.Scanln(&altura)
	fmt.Println("Digite o sexo (Masculino ou Feminino): ")
	fmt.Scanln(&sexo)

	if sexo == "Masculino" {
		pesoideal = (72.7 * altura) - 58
	} else if sexo == "Feminino" {
		pesoideal = (62.1 * altura) - 44.7
	} else {
		fmt.Println("Inválido.")
	}
	var peso float64
	fmt.Println("Digite seu peso: ")
	fmt.Scanln(&peso)

	if peso < pesoideal {
		fmt.Printf("Você está abaixo do peso.")
	} else if peso > pesoideal {
		fmt.Printf("Você está acima do peso.")
	} else {
		fmt.Printf("Você está dentro do peso.")
	}
}
