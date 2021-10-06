package main

import "fmt"

//SLICE
func main() {
	amigos := make([]string, 3) /*uma variavel amigos que armazenara um slice com valores do tipo
								string e tem um tamanho inicial de 3 posiçoes que poderão ser alterado
	 							sem nenhum problmea. É atraves do make que pode criar o slice*/
	nome := ""
	i := 0
	for nome != "q" { //se o nome digitado for diferente de 'q'
		fmt.Print("Digite o nome dos seus amigos ou 'q' para parar: ") //imprima no console
		fmt.Scanf("%s\n", &nome)                                       //o nome digitado sera anexado a variavel nome
		if nome != "q" {                                               //se nome for diferente de q
			if i < 3 { //se i for menor que 3, que é meu numero maximo de posição dentro do meu slice
				amigos[i] = nome //o indice i do meu slice amigos, tera o valor de nome
			} else { // caso meu i for >= 3
				amigos = append(amigos, nome) /*faço um append, uma extensão da quantidade de elementos do meu
				  slice amigos e adiciono a variavel que foi digitade e armazenada no 'nome'*/

			}
			i++
		}
	}
	fmt.Println(amigos)                             //imprime minha lista de amigos no console
	fmt.Printf("Voce tem %d amigos\n", len(amigos)) //len(amigos) trás o tamanho do meu slice

	for _, nomeAmigo := range amigos { //vou imprimir todos os nomes presentes no meu slice
		fmt.Printf(" - %s\n", nomeAmigo)
	}

	//CRIANDO UM SUBSLICE
	outroSlice := amigos[1:3] //estou criando um subslice, como os valores o indice 1 ate o indice 2, o 3 nao sera incluso nesse meu subslice
	fmt.Println(outroSlice)
}
