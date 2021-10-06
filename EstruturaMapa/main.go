package main

import (
	"bufio" // mesma importação do Scanner no Java, substituira o Scanf ja que o scanf nao permite fazer varias leituras
	"fmt"
	"os"
)

//MAPAS

//mapas são uteis quando tenho duas informações co relacionadas
func main() {
	mapaCurso := make(map[string]int)     //dentro da variavel mapaCurso, criei um mapa que a chave sera o nome do curso (string) e  rerspectivo valor (int, correponsde a carga horaria)
	scanner := bufio.NewScanner(os.Stdin) //variavel scanner que ira ler o valor de entrada (stdin) do meu sistema operacional (os). ATENÇÃO às importações necessarias!
	curso := ""
	for curso != "q" {
		var cargaHoraria int
		fmt.Print("Digite um curso ou digite q para sair: ")
		scanner.Scan()         //pegara os vaores digitados no meu console e anexara dentro do meu objeto scanner
		curso = scanner.Text() //o comando Text trara todo valor lido pelo metodo chamado scan
		if curso != "q" {
			fmt.Print("Digite a carga horaria do curso: ")
			fmt.Scanf("%d\n", &cargaHoraria)
			//ja que temos o nosso curso e a carga horaria, agora é hora de colocar as informações dentro do nosso mapa
			mapaCurso[curso] = cargaHoraria //aonde eu estiver a chave correspondente ao nome do curso, eu terei a carga hora
			/*EXEMPLO: se tivermos um curso de Go de 50h de carga horaria e você digiar 'Go',
			o valor corresponde a carga horaria desse curso sera 50. A chave é o nome do curso, ou seja, 'Go'*/

		}
	}
	fmt.Println("Cursos registrados: ")
	for nomeCurso, cargaHorariaDoCurso := range mapaCurso { //ler os valores do meu mapa
		fmt.Printf("-%s: %dh\n", nomeCurso, cargaHorariaDoCurso) //imprime o curso e a carga horaria correspondente
	}
	curso = ""
	for curso != "q" {
		fmt.Print("Digite o nome do curso a ser excluido ou digite q para cancelar: ")
		scanner.Scan()         //le o que fooi digitado no console e armazena dentro da minha variavel scanner
		curso = scanner.Text() //minha variavel curso sera igual ao valor que foi anexado dentro da minha variavel scanner
		if curso != "q" {      //se meu curso for diferente de q
			cargaHorariaExcluido, cursoExiste := mapaCurso[curso] //crio duas variaveis
			if cursoExiste {                                      //se meu curso existir
				delete(mapaCurso, curso) //delete o curso que foi digitadno no if anterior
				fmt.Printf("O curso %s com %dh foi removido \n", curso, cargaHorariaExcluido)
			} else {
				fmt.Println("O curso digitado não existe")
			}
		}
	}
	fmt.Println("Cursos registrados: ")
	for nomeCurso, cargaHorariaDoCurso := range mapaCurso { //ler os valores do meu mapa
		fmt.Printf("-%s: %dh\n", nomeCurso, cargaHorariaDoCurso) //imprime o curso e a carga horaria correspondente
	}
}
