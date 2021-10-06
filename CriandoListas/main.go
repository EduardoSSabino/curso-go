package main

import (
	"container/list" //importação necessaria para fazer uma lista um pouco mais tradicional, um pouco mais proxima do que temos em outras linguagens
	"fmt"
)

func main() {
	numeros := list.New()     // list.New(), euestarei inicializando uma lista
	el := numeros.PushBack(1) //insere  um elemento na parte de tras da lista, nesse caso o numero 1 e esarei armazenando dentro da minha variavel 'el'
	numeros.PushFront(0)      //insere um elemento na parte da frente da lista, nesse caso o nnumero 0
	numeros.PushBack(2)
	/*a sequencia ficou assim, PushBack de 1, depois PushFronte de 0,
	o 0 ja ficou na frente do 1 por ser Front, depois dei um PushBack de 2,
	 como é PushBack, foi pro final, ficando assim: 0 1 2*/

	//agora eu vou varrer essa lista
	for e := numeros.Front(); e != nil; e = e.Next() { //crio a variavel "e" e ela sera igualao numero da frente da minha lista, ou seja o primeiro numero.
		//se minha variavel "e" for nula (em Go null é escrito como nil)
		//minha variavel "e" sera igual ao proximo valor de "e", lembrando, essa lista sabe qual é o proximo valor de "e"
		fmt.Println(e.Value) //imprime o valor da variavel "e"
	}
	fmt.Println("------------")                        //colocando uns traços para facilitar a compreensão
	numeros.Remove(el)                                 //remove a variavel "el" da minhas lista numeros
	for e := numeros.Front(); e != nil; e = e.Next() { //esse 'for' ja foi comentado acima
		fmt.Println(e.Value)

	}
}
