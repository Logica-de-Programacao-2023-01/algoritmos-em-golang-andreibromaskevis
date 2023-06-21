package main

import "fmt"

func main() {

	var num, soma, contador float64
	fmt.Println("Digite varios números:")

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		soma += num
		contador++

	}
	resultado := soma / contador
	fmt.Printf("Média aritmética: %.2f ", resultado)

}
