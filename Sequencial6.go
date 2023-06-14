package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Digite o seu salário: ")
	fmt.Scanln(&salario)

	aumento := salario * 0.15
	salarionovo := salario + aumento

	fmt.Print("Novo salario: ", salarionovo)
}
