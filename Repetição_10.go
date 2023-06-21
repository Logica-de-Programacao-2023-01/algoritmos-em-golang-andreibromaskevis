package main

import "fmt"

func main() {
	var numero int
  var maior int

	fmt.Println("Digite varios números (digite 0 para parar):")
	fmt.Print("Número: ")
	fmt.Scanln(&numero)
	maior = numero
	
	for numero != 0 {
		if numero > maior {
			maior = numero
		}
		fmt.Print("Número: ")
		fmt.Scanln(&numero)
	}
	fmt.Println("O maior número é:", maior)
}
 
