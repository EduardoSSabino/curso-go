package main

import (
	"fmt"
)

//ESCOPO: é o bloco ao qual minha variavel pertence, no codigo abaixo, eu usei a variavel resultado algumas vezes e mesmo assim nao de erro, pois ela estva em outro escopo (bloco)
func main() {
	var numero1, numero2 int
	var operacao string
	fmt.Print("Digite o primeiro numero: ")
	fmt.Scanf("%d\n", &numero1)
	fmt.Print("Digite a operação (+ - * / $): ")
	fmt.Scanf("%s\n", &operacao)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scanf("%d\n", &numero2)
	if operacao == "+" { //se for uma soma
		somar(numero1, numero2)
	} else if operacao == "-" { //s for uma substração
		resultado := substrair(numero1, numero2)
		fmt.Printf("%d - %d = %d", numero1, numero2, resultado)
	} else if operacao == "*" { //se for uma multiplicação
		resultado := multiplicar(numero1, numero2)
		fmt.Printf("%d x %d = %d", numero1, numero2, resultado)
	} else if operacao == "/" { //se for uma divisão
		resultado, resto := dividir(numero1, numero2)
		fmt.Printf("QUOCIENTE = %d; RESTO = %d", resultado, resto)
	} else if operacao == "$" { //se for um incremento ou decremento
		incremento, decremento := incrementoDecremento(numero1, numero2)
		fmt.Printf("INCREMENTO = %d; DECREMENTO = %d", incremento, decremento)
	} else { //se nao for nenhuma das operações acima
		fmt.Print("Operação inválida!")
	}
}

func somar(n1 int, n2 int) { //função void
	fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
}

func substrair(n1 int, n2 int) int { //ja que essa função nao se trata de uma função void, depois dos parametros, tenho que especificar o tipo do meu retorno, nesse caso um int
	return n1 - n2 //ira retornar a subtração que que sera do tipo int
}

func multiplicar(n1 int, n2 int) (resultado int) { //como essa função tera um retorno, estou armazenando o retorno dentro da minha variavel resultado que é do tipo int
	resultado = n1 * n2
	return //ira retornar resultado
}

func dividir(n1 int, n2 int) (int, int) { //essa função tera dois retornos do tipo int, a primeira sera meu quociente e o segundo meu resto
	quociente := n1 / n2
	resto := n1 % n2
	return quociente, resto
}

func incrementoDecremento(n1 int, n2 int) (inc int, dec int) { //essa função terá dois retornos armazenados dentros de duas variaveis (inc e dec, ambas do tipo int)
	inc = n1 + n2
	dec = n1 - n2
	return //ira retornar inc e dec
}
