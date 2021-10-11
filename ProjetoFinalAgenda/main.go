package main

import (
	"bufio" //importação necessaria pra poder usar o scanner
	"fmt"
	"os"
	"strings"
)

// ARQUIVO : indica o locar onde os contatos serão salvos
const ARQUIVO string = "agenda.txt" //é nesse arquivo que colocaremos nossos contatos

// ConversiveParaString : interface que especifica como um item podem ser convertudo para uma string
type ConversivelParaString interface { //um "contrato" que precisara ser seguido
	toString() string
}

// Contato : estrutura que representa um contato
type Contato struct { //crio uma classe contato e defino seus atributos
	nome           string
	formaConato    string
	valorDoContato string
}

func (contato *Contato) toString() string {
	return fmt.Sprintf("%s|%s|%s \n", contato.nome, contato.formaConato, contato.valorDoContato)
}

// GerenciadorContatos : componente responsavel por gerenciar os contatos
type GerenciadorContatos struct{}

func (gerenciador *GerenciadorContatos) carregarContatos() ([]Contato, error) {
	contatos := make([]Contato, 0)
	if _, e := os.Stat(ARQUIVO); !os.IsNotExist(e) { /*ele rtorna duas coisas, um ponteiro com uma estrutura que representa as varias informções
		do arquivo que estamos dando stat e retorna tbm uma estrutura de erro indicando caso haja algum erro
		 na hora de se obter as infromações sobre o arquivo, inclusive, ele retornara um erro caso o arquivo nao exista*/

		arquivo, err := os.Open(ARQUIVO) //abre o arquivo para leitura
		if err != nil {
			return contatos, err
		}
		defer arquivo.Close()
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			listaContato := scanner.Text()
			//agora vou fazer um slip pra separar cada informção pelo po line, pra isso terei que fazer mais uma importação
			infoContato := strings.Split(listaContato, "|")                                                                         //vai separar minhas informações pelo '|'
			contatos = append(contatos, Contato{nome: infoContato[0], formaConato: infoContato[1], valorDoContato: infoContato[2]}) //adicionei um novo contato
		}
	}
	return contatos, nil
}

func (gerenciador *GerenciadorContatos) salvarConato(contato ConversivelParaString) error {
	arquivo, err := os.OpenFile(ARQUIVO, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) //abro meu arquivo(arquivo a ser aberto, adiciono linhas ao arquivo, e posso criar algo, e por fim, e posso escrever nesse arquivo)
	// esse 0666, uma função que é a escrita toda ali no arquivo
	//retornara uma string para o arquivo e um possivel erro que pode acontecer
	if err != nil { //ou seja, houve um erro
		return err
	}
	defer arquivo.Close()                                          //garanto que meu arquivo sera encerrado
	if _, e := arquivo.WriteString(contato.toString()); e != nil { //qual a string que eu espero que seja escrita dentro do meu arquivo
		return e //retorna meu erro
	}
	return nil
}

func main() {
	gerenciador := GerenciadorContatos{}
	opcao := 0
	for true {
		fmt.Println("O que você deseja fazer? ")
		fmt.Println(" 1 - Listar os contatos")
		fmt.Println(" 2 - Criar novo contatos")
		fmt.Println(" 3 - Sair")
		fmt.Scanf("%d", opcao)
		switch opcao {
		case 1:
			listarContatos(&gerenciador)
		case 2:
			criarNovoContato(&gerenciador)
		}
		if opcao == 3 {
			break
		}
	}
	fmt.Print("Tchau!")
}

func listarContatos(gerenciador *GerenciadorContatos) { //listarContatos deve receber um parametro GerenciadorContatos
	contatos, err := gerenciador.carregarContatos() //carregarContatos é um metodo
	if err != nil {
		fmt.Printf("Não foi possivel carregar os contatos: %s \n", err)
	} else { //se eu tiver um erro, é pq eu consegui carregar minha lista de contatos e eu vou escrever ess linha
		fmt.Println("Lista de contatos: ")
		for _, contato := range contatos { //vou varrer meu slice
			fmt.Printf("   - %s, %s: %s \n", contato.nome, contato.formaConato, contato.valorDoContato) //dessa maneira eu consigo fazer a lista de contatos
		}
	}
}
func criarNovoContato(gerenciador *GerenciadorContatos) {
	novoContato := Contato{}
	fmt.Print("Nome do contato: ")
	fmt.Scanf("%s", &novoContato.nome)
	fmt.Print("Tipo do contato: ")
	fmt.Scanf("%s", &novoContato.formaConato)
	fmt.Print("Contato: ")
	fmt.Scanf("%s", &novoContato.valorDoContato)
	err := gerenciador.salvarConato(&novoContato)
	if err != nil {
		fmt.Printf("Houve um erro ao salvar o contato: %s \n", err)
	}

}
