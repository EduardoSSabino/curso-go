package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estou abrindo a conexão ao banco de dados!")
	defer fmt.Println("Estou fechando a conexão ao bando de dados!") //o defer me garante que independente de qualquer coisa
	//essa linha sera executada quando a função main chegar ao fim
	executarCosnulta()
}

func executarCosnulta() {
	fmt.Println("Estou executando a consulta ao banco de dados!")
}

//se olharmos no terminal, vimos que mesmo que a linha do defer esteja no meio da funação, o programa segura e execução da linha pra deixar pra ser executada por ultimo
