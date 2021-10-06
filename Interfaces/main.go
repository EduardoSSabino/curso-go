package main

import (
	"errors"
	"fmt"
)

//INTERFACES: contratos]

type Acelerador interface {
	//agora vou descrever os comportamentos exigijidos pela minha interface acelerador
	acelerar() error //quem se submeter a usar minha interface acelerador, se compromete a usar meu metodo acelerar retornadno um erro
}

type Freador interface {
	frear() error
}

type Carro struct { //crio uma classe carro e adiciono somente os atributos
	modelo     string
	marca      string
	velocidade float64
}

func (carro *Carro) acelerar() error { //meu carro tem o comportamento de acelerar.crio um metodo acelerar. tera um retorno do tipo erro
	if carro.velocidade < 15 {
		carro.velocidade += 5
		return nil //retornara que nao houve erros
	} else {
		return errors.New("o carro ja esta na sua velocidade maxima")
	}
}

func (carro *Carro) frear() error { //Meu carro tem o comportamento de frear. crio um metodo frear. tera um retorno do tipo erro
	if carro.velocidade > 0 {
		carro.velocidade -= 5
		return nil //retornara que nao houve erros
	} else {
		return errors.New("o carro ja está parado")
	}
}
func main() {
	carro := Carro{modelo: "Palio", marca: "Fiat", velocidade: 0} //se fosse no java, eu daria une new pra poder criar o objeto. estou instanciando
	opcao := 0
	for opcao != 3 { //ja que eu nao tenho while, usei o laço for
		fmt.Println("O que voce deseja fazer?")
		fmt.Println("1- Acelerar")
		fmt.Println("2- Frear")
		fmt.Println("3- Sair")
		fmt.Scanf("%d\n", &opcao) //leio o console e anexo o valor na variavel opcao
		var err error = nil       //se variavel err =  nil que é o mesmo que nul que nessa situação significa a ausência de erros
		switch opcao {
		case 1:
			err = fazerAcelerar(&carro) //acione o metodo acelerar
		case 2:
			err = carro.frear() //acione o metodo frear
		}
		if err != nil { //se eu tiver um erro
			fmt.Printf("** Não foi possivel completar a ação: %s ** \n", err) //imprima isso no console
		} else if opcao != 3 {
			fmt.Printf("O carro da %s da marca %s está a %.2f km/h", carro.modelo, carro.marca, carro.velocidade)
		}
	}
	fmt.Println("Fim da execução!")
}

func fazerAcelerar(veiculo Acelerador) error {
	return veiculo.acelerar()

}

func fazerFrear(veiculo Freador) error {
	return veiculo.frear()
}
