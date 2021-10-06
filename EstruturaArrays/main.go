package main

import (
	"fmt"
)

func main() {
	var amigos [5]string // em go, eu coloca o nome do meu array, depois o tamnho e por fim o tipo da variavel
	fmt.Println(amigos)
	for i := 0; i < len(amigos); i++ { //len(nomeVar[]), pega o tamanho do meu array
		fmt.Print("Digite o nome de seus amigos: ")
		fmt.Scanf("%s\n", &amigos[i])
	}
	fmt.Println(amigos)
	//agora e vou mostrar todos os nomes presentes no meu vetor
	fmt.Println("Seus amigos são: ")
	for _, amigo := range amigos { //range: varrera meu array, trazendo todos valores presentes. Mesma coisa que um foreatch

		/*esse underline depois do meu for, faz com que eu ignore o indice,
		esse ciclo for, representado dessa maneira, me tras duas coisas,
		o indice que estou inteirando e o valor do indice, como disse, esse underline faz com que eu
		so receba o meu valor presente naquele respectivo indice*/
		fmt.Printf(" - %s\n", amigo)
	}

	//declaradno a variavel ja com o valores de seus indices
	arrayInicializado := [3]int{2, 4, 6}
	fmt.Println(arrayInicializado)

	//declaradno uma matriz
	var matriz [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Digite um numero: ")
			fmt.Scanf("%d\n", &matriz[i][j])
		}
	}
	fmt.Print(matriz)
}
