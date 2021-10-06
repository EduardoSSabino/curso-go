package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um numero: ")
	fmt.Scanf("%d", &numero)
	if numero <= 5 {
		fmt.Println("Numermo menor que 5e menor que 10")
	} else if numero < 10 {
		fmt.Println("numero menor que 10 e maior que 5")
	} else {
		fmt.Println("Numero maior que 10")
	}
}
