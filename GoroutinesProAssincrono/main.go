package main

import (
	"fmt"
	"time" //importação que me permite dar um delay no programa
)

//GOROUTINES

//Dei uma pesquisada, e cheguei a conclusão que a definição pra uma goroutine é a mesma de uma thread.
/*Um exemplo, criei um jogo, preciso fornecer imagem e audio, obviamente ao mesmo tempo
(na visão humana, mas sera executado uma linha por vez obviamente), isso seria um exemplo de goroutine*/

//OUTRA DEFINIÇÃO DE GOROUTINE: goroutine constituem métodos que são executados em threads diferentes da thread principal e, por isso, podem ser executadas de maneira assíncrona

func main() {
	var limite int           //crio uma variavel limite do tipo int, endereçada dentro do meu escopo main
	canal1 := make(chan int) //crio meus channels. trocade informções entre minhas goroutines
	canal2 := make(chan int)
	fmt.Print("Informe um limite: ")
	fmt.Scanf("%d", &limite)
	go conteAte(limite, canal1) //o go fara com que minha função main seja executada junto com a função conteAte, de maneira assicrona
	go func(n int) {            //crio outr goroutine , ou seja, ira funcionar de forma assicrona
		resultado1 := 0
		for j := 0; j <= n*10; j++ { //enquanto meu j for menor ou igual a n, imprima o valor de j
			fmt.Printf("- [anonimo] O numero é %d \n", j)
			resultado1 = j * 10
		}
		canal2 <- resultado1 // e aqui esta acontecendo o channel, estou anexando o valor da resultado1 dentro do canal2
	}(limite)
	for p := 0; p <= limite*10; p++ {
		fmt.Printf(" - [main] O numero é %d \n ", p)
	}
	time.Sleep(10 * time.Second) //delay de 10 segundos
	resultadocanal1 := <-canal1  //crio uma variavel resultadocanal1 que armazenara o valor da minha variavel canal1
	resultadocanal2 := <-canal2
	fmt.Printf("Canal 1 = %d, canal 2 = %d \n", resultadocanal1, resultadocanal2)
}

func conteAte(limite int, canal chan int) { //como segundo parametro, crio uma variavel do tipo int que sera um canal pra receber alguma informação
	resultado := 0
	for i := 0; i <= limite*20; i++ {
		fmt.Printf("- [conteAte] O numero é %d \n", i)
		resultado = i * 20
	}
	canal <- resultado //minha variavel canal armazenara o valor de resultado, resultado= ao meu ultimo resultado
}

//CHANNEL são estruturas que permitem que informção sejam trocadas entre goroutines.
//NESTE EXEMPLO, eu tenho duas variaveis em dois escopos diferentes (resultao1 e a resultado), como eu posso trazer-las sendo que estao em funções diferentes?
//Pra isso usarei o channel
