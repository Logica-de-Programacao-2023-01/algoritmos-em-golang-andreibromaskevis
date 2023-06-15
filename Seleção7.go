package main

import "fmt"

func main() {

	var s float64
	var ns float64
	fmt.Print("Digite o salário do funcionário:")
	fmt.Scanln(&s)

	if s < 1000 {
		ns = s + s*0.1
		fmt.Println("Novo salário (+10%):", ns)
	} else {
		ns = s + s*0.05
		fmt.Println("Novo salário (+5%):", ns)
	}
}
