package main

import "fmt"

func main() {
	var diasdeTrabalho int
	var valordadiaria float64
	fmt.Print("Digite o número de dias trabalhados: ")
	fmt.Scan(&diasdeTrabalho)
	fmt.Print("Digite o valor da diária: ")
	fmt.Scan(&valordadiaria)

	salario := float64(diasdeTrabalho) * valordadiaria

	fmt.Print("O salário do funcionário é: ", salario)
}
