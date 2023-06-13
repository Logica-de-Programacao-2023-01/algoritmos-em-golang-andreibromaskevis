package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Println("Digite três números: ")
	fmt.Scan(&num1, &num2, &num3)

	mediap := (num1*2 + num2*3 + num3*5) / (2 + 3 + 5)

	fmt.Print("Média ponderada: ", mediap)
}
