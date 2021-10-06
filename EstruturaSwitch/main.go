package main

import (
	"fmt"
)

//Quando usamso o Switch no Go, nao precisamo colocar o break no fim dos cases
func main() {
	var numero int
	fmt.Print("Digite um numero: ")
	fmt.Scanf("%d", &numero)

	switch numero {
	case 1:
		fmt.Println("numero igual a 1")
	case 2:
		fmt.Println("numero igual a 2")
	case 3:
		fmt.Println("numero igual a 3")
	case 4:
		fmt.Println("numero igual a 4")
	default:
		fmt.Print("Numero nao incluso no Switch")
	}
	switch numero {
	case 2, 4:
		fmt.Print("Numero par")
	case 1, 3:
		fmt.Print("Numero impar")
	}
}
