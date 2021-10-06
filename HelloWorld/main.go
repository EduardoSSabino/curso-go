package main //é obrigatório trazer o package main! e assim como no Java, é a primeira coisa que eu devo declarar

import ( //aqui dentro dessas chaves, vou colocar todas minhas importações necessarias. OBS: as importações ficarão todas dentro de aspas
	"fmt" //bom, essa importação fmt é ncessaria em todos os codigos, é com ela que conseguiremos usar o println por exemplo
)

//Meus metodos eu posso declarar com uma função
func main() { //func main corresponde ao public static void main, ou seja, minha principal função

	fmt.Println("Ola mundo!")            //imprime na tela Olá mundo
	fmt.Println("TreinaWeb", " Cursos")  // imprimira TreinaWeb Cursos
	fmt.Println("1 + 1 = ", 1+1)         //imprimira 1 + 1 = 2
	fmt.Println("1.1 + 2.2 = ", 1.1+2.2) //imprimira 1.1 + 2.2 = 3.3
	fmt.Println(true)                    //imprimira true

	//definindo variaveis em go

	//NUMEROS SEM SINAL

	var s1 uint8 = 255    //No caso de uint8, poderia usar byte. Variavel chamada primeiraVar do tipo uint8 (ocupa ate 8 bits) tem o valor de 255. OBS: uma variavel do tipo uint8 tem o valor maximo de 255
	var s2 uint16 = 65535 //variavel chamada segundaVar do tipo uint16 (ocupa ate 16 bits) tem o valor de 65535. OBS: uma variavel do tipo uint16 tem o valor maximo de 65535
	var s3 uint32 = 44    //variavel chamada terceiroVar do tipo uint32 (ocupa ate 32 bits) tem o valor de 44. OBS: uma variavel do tipo byte tem o valor maximo de >4bi
	var s4 uint64 = 10    //variavel chamada quartoVar do tipo uint64 (ocupa ate 64 bits) tem o valor de 10. OBS: uma variavel do tipo uint64 tem o valor maximo de >18vintelhoes
	fmt.Println("var sem sinal\n", s1, s2, s3, s4)

	//INTEIROS COM SINAL
	//OBS: quando estamos tratando de uma var do tipo int, ou seja, uma variavel que tem sinal, a capacidade de armazanamento cai pela metade. EX: unit8 = 255, int8 = 127

	var c1 int8 = 127    //variavel chamada primeiraVar do tipo int8 (ocupa ate 8 bits) tem o valor de 2127. OBS: uma variavel do tipo int8 tem o valor maximo de 127
	var c2 int16 = 32335 //variavel chamada segundaVar do tipo int16 (ocupa ate 16 bits) tem o valor de 32335. OBS: uma variavel do tipo int16 tem o valor maximo de 32767
	var c3 int32 = 44    //No caso de um int32, poderia usar rune. Variavel chamada terceiroVar do tipo int32 (ocupa ate 32 bits) tem o valor de 44. OBS: uma variavel do tipo int32 tem o valor maximo de >2bi
	var c4 int64 = 10    //variavel chamada quartoVar do tipo int64 (ocupa ate 64 bits) tem o valor de 10. OBS: uma variavel do tipo int64 tem o valor maximo de >9vintel
	fmt.Println("var com sinal\n", c1, c2, c3, c4)

	//ATALHOS "INT/UINT SEM O BIT DECLARADO"

	var sBit1 uint = 30 //o valor do bit vai de acordo com minha arquitetura
	var sBit2 int = 60  //o valor do bit vai de acordo com minha arquietetura
	fmt.Println("Declarando variavel sem declarar o bit\n", sBit1, sBit2)

	//PONTOS FLUTUANTES, FLOAT

	var f1 float32 = 1.1 //segue a mesma logica das variaveis acima
	fmt.Println("Variavel do tipo float\n", f1)

	//NUMEROS COMPLEXOS

	var cx1 complex64 = 3 //nuemeros complexos, suporta numeros imaginarios, segue a mesma logica das variaveis acima
	fmt.Println("Declarando variaveis com numeros complexos\n", cx1)

	//STRINGS

	var string1 string = "TreinaWeb"       //variavel do tipo string usando aspas duplas
	var string2 string = `Curso TreinaWeb` //variavel do tipo string usando crase
	fmt.Println("Declarando variaveis do tipo string\n", string1, string2)

	//BOOLEANOS

	var bool1 bool = true  //variavel do tipo booleana, true
	var bool2 bool = false //variavel do tipo booleana, false
	fmt.Println("Declarando variavel booleana\n", bool1, bool2)

	//MULTIPLAS DECLARAÇÕES

	var nome, sobrenome string = "Eduardo", "Sabino" //criando duas variaveis ao mesmo tempo, obviamente eas sao do mesmo tipo
	var (
		apelido string  = "Dudao"
		idade   int     = 20
		altura  float32 = 1.80
	)
	fmt.Println("Multiplas declaracoes\n", apelido, idade, altura, nome, sobrenome)

	//DECLARAÇÃO DE VARIAVEL SEM DECLARAR O TIPO "INFERENCIA DE TIPO"

	var abreviacaoNome = "Edu"
	var sizePe = 43
	fmt.Println("Declarando variavel sem declarar o tipo\n", abreviacaoNome, sizePe)

	//ATALHOS A DECLARAÇÃO DE VARIAVEIS

	cursoFacul := "Engenharia" //usando o atalho := eu evito de digitar "var", deixando a construção do codigo menos verbosa
	fmt.Println("Atalho para evitar o uso do 'var'\n", cursoFacul)

	//CONSTANTES

	const constante1 string = "TreinaWeb" //declarando uma constante do tipo string
	const (
		primeiroNome = "Eduardo"
		segundoNome  = "Sabino"
	)
	fmt.Println("Declarando constantes\n" + primeiroNome + " " + segundoNome + " esta usando a plataforma de estudos " + constante1)

} //no meu terminal digitarei go run main.go, esses comandos executarão meu programa
