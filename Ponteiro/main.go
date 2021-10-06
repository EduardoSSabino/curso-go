package main

import (
	"fmt"
)

func main() {
	var saldo float64 //crio uma variavel sald do tipo float64
	fmt.Print("Digite o seu saldo: ")
	fmt.Scanf("%f\n", &saldo)
	calcularRendimento(&saldo)                                   //aqui dentro eu nao passo o valor da variavel saldo, e sim a posição de memoria
	fmt.Printf("Seu saldo com rendimento é de R$%.2f\n ", saldo) //meu saldo é uma variavel do tipo float com a precisão de duas casas decimais (%.2f)
}

func calcularRendimento(saldo *float64) {
	*saldo += *saldo * 0.03 //nesse caso, os dois primeiros *, são ponteiros de memoria
}
