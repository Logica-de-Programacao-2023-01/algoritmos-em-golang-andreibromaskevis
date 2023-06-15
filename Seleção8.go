package main

import "fmt"

func main() {

	var num int
	var diadaSemana string
	fmt.Println("Digite um número entre 1 e 7: ")
	fmt.Scanln(&num)

	if num == 1 {
		diadaSemana = "Domingo"
	} else if num == 2 {
		diadaSemana = "Segunda-feira"
	} else if num == 3 {
		diadaSemana = "Terça-feira"
	} else if num == 4 {
		diadaSemana = "Quarta-feira"
	} else if num == 5 {
		diadaSemana = "Quinta-feira"
	} else if num == 6 {
		diadaSemana = "Sexta-feira"
	} else if num == 7 {
		diadaSemana = "Sábado"
	} else {
		diadaSemana = "Inválido"
	}
	fmt.Print("Dia da semana correspondente: ", diadaSemana)
}
