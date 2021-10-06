package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bem vindo!") //da as boas vindas ao usuario
	defer func() {            //crio uma função anonima usando o defer, ou seja, SERA EXECUTADO POR ULTIMO DENTRO DO MEU ESCOPO MAIN
		ex := recover() //crio uma variavel ex que armazenara meu ultimo erro. recover = recupera meu ultimo erro.
		if ex != nil {  //depois eu faço a comparação se meu ex é difernte de nulo, ou seja, se for diferente de nulo foi pq realemente teve algum erro
			fmt.Printf("Desculpe. Ocorreu um erro: %s", ex)
		}
		fmt.Println("Obrigado por usar nosso sofware!")
	}() //esses parenteses extras são necessarios para garantir que o defer ira ser executado no fim da função main
	var numero int
	fmt.Print("Digite um numero maior que 10: ")
	fmt.Scanf("%d", &numero)
	if numero <= 10 {
		panic("Numero inválido!\n")
	} else {
		fmt.Println("Numero válido. Parabéns!")
	}
}
