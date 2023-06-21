package main

import "fmt"

func main() {
	var num int

  fmt.Println("Digite um número: ")
	fmt.Scanln(&num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d vezes %d é igual a %d\n", num, i, num*i)

	}

}
