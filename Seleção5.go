package main

import "fmt"

func main() {

	var num int
	fmt.Printf("digite um número: ")
	fmt.Scanln(&num)

	if num%3 == 0 {
		fmt.Printf("%d é múltiplo de 3: ", num)
	}
	if num%5 == 0 {
		fmt.Printf("%d é múltiplo de 5: ", num)
	} else if num%3 == 0 && num%5 == 0 {
		fmt.Printf("%d é múltiplo de 3 e de 5: ", num)
	}
}
