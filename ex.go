package main

import "fmt"

func main() {
	var (
		nome  string
		idade int
		peso  float64
	)
	fmt.Println("Qual é o seu nome? ")
	fmt.Scan(nome)
	fmt.Println("Qual é a sua idade? ")
	fmt.Scan(idade)
	fmt.Println("Qual é o seu peso? ")
	fmt.Scan(peso)
}
