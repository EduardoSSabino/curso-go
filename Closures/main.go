package main

import (
	"fmt"
)

//CLOSURES: é uma variavel que gaurda uma função
func main() {
	var numero1, numero2 int
	var operacao string
	fmt.Print("Diite o primeiro numero: ")
	fmt.Scanf("%d\n", &numero1)
	fmt.Print("Digite a operação(+ - * / $): ")
	fmt.Scanf("%s\n", &operacao)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scanf("%d\n", &numero2)
	var metodoOperacao func(n1 int, n2 int) (int, int) //aqui vemos que eu criei uma variavel que guarda um função
	if operacao == "+" {
		metodoOperacao = func(n1, n2 int) (int, int) { //função que retornara dois valores inteiros
			return n1 + n2, 0 //retorna o resultado de n1+n2 e retorna zero como segundo valor
		}
	} else if operacao == "-" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 - n2, 0
		}
	} else if operacao == "*" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 * n2, 0
		}
	} else if operacao == "/" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 / n2, n1 % n2
		}
	} else if operacao == "$" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 + n2, n1 - n2
		}
	}
	resultado1, resultado2 := metodoOperacao(numero1, numero2) //resultado1 sera igual ao meu primero retorno, e resultado2 sera igual ao meu segundo retorno
	fmt.Printf("%d %s %d = %d, %d", numero1, operacao, numero2, resultado1, resultado2)
}
