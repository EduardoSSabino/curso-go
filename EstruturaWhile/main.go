package main

import (
	"fmt"
)

//Na linguagem go, nao temos a função while, para fazer a mesma função que o while, basta subistituir o while por for, unica alteração necessaria
func main() {
	var numero int
	fmt.Print("Digite um numero: ")
	fmt.Scanf("%d", &numero)
	for numero < 10 {
		fmt.Print("\n", numero)
		numero++
	}

}
