package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Digite um número :")
	fmt.Scan(&num1)

	num_suc := num1 + 1
	num_ant := num1 - 1

	fmt.Printf("O seu sucessor é %d, e o seu antecessor é %d", num_suc, num_ant)
}
