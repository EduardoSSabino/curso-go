package main

import (
	"fmt"
)

//PANIC: ela registra uma excessão (throw no Java), uma situação anormal, um erro de digitação, exception
func main() {
	//por exemplo, nesse caso teremos a regra que o usuario tera que digitar um numero maior que 10, se ele digitar um numero menor que 10, teremos uma situação erro, exception
	fmt.Println("Bem vindo!")
	defer fmt.Println("Obrigado por utilizar nosso software!") //defer sempre sera ultima linha a ser executada, indenpendente da sua posição no codigo!
	fmt.Print("Digite um numero maior que 10: ")
	var numero int
	fmt.Scanf("%d", &numero)
	if numero <= 10 {
		panic("Numero inválido!") //o usuario digitou um numero não condizente com o esperado, ou seja, um erro de execução
	} else {
		fmt.Print("Você fez tudo correto. Parabéns!\n")
	}
}
