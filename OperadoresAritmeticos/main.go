package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numero1 int
	var numero2 int
	fmt.Print("Digite o primeiro numero: ")
	fmt.Scanln(&numero1) //com o scanln ou scanf, eu posso falar a variavel que armazenara o valor digitado no console
	fmt.Print("Digite o segundo numero: ")
	fmt.Scanln(&numero2)
	fmt.Printf("%d + %d = %d \n", numero1, numero2, numero1+numero2)  //substituira os "%d" respectivamente pel valores colocados depois das aspas de fechamento
	fmt.Printf("%d - %d = %d \n", numero1, numero2, numero1-numero2)  //substituira os "%d" respectivamente pel valores colocados depois das aspas de fechamento
	fmt.Printf("%d * %d = %d \n", numero1, numero2, numero1*numero2)  //substituira os "%d" respectivamente pel valores colocados depois das aspas de fechamento
	fmt.Printf("%d / %d = %d \n", numero1, numero2, numero1/numero2)  //substituira os "%d" respectivamente pel valores colocados depois das aspas de fechamento
	fmt.Printf("%d %% %d = %d \n", numero1, numero2, numero1%numero2) //substituira os "%d" respectivamente pel valores colocados depois das aspas de fechamento

	incremento := numero1
	incremento += numero2 //é o mesmo que fazer incremento = incremento - numero2
	fmt.Printf("O incremento de %d com %d é %d\n", numero1, numero2, incremento)
	decremento := numero1
	decremento -= numero2 //é o mesmo que fazer decremento = decremento - numero2
	fmt.Printf("O decremento de %d com %d é %d\n", numero1, numero2, decremento)

	//CONVERSÃO DE TIPO DE DADOS
	//poderia fazer a mesma coisa usando o metodo parse, por exemplo:
	var texto string
	var texto1 string
	var number1 int

	fmt.Printf("Digite um numero: ")
	fmt.Scanf("%s", &texto1)
	fmt.Print("Digite outro numero: ")
	fmt.Scanf("%s", &texto)

	number, _ := strconv.ParseInt(texto, 10, 32) //falo minha variavel a ter seu tipo transformado, depois falo se é um numero decimal/hexadecimal/oitavo, e por fim declaro o tamanho, ous eja, se vai ocupar 32b/64b e por ai vai/
	var conv = float64(number)                   //aqui eu converti de int32 para float64
	fmt.Println(number)
	fmt.Println(conv)

	//agora a meta é ler a informação digitada como se fosse um numero, pra isso, usarei outra importação,

	number1, _ = strconv.Atoi(texto1) //transformara minha variavel texto em um valor inteiro e armazena dentro dentro da minha variavel numero
	fmt.Println(number1)

}
