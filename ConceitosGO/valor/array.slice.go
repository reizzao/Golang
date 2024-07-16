package valor

import "fmt"

// Arrays_Slices

var array1 = []int{10, 20}

func Tester_Array() {
	// add valor no array de forma hard : referencia o arrayAlvo e dentro de colchetes a posicao onde deseja inserir o novo item.
	array1[1] = 1000

	// add valor no array de forma dinamica
	// uso: dinamicamente: append( arrayAlvo, dpois quantos items desejar separados por virgula) -- vai devolve um novo array com os items que ja tinha e os add depois da variavel do array alvo
	novoArray := append(array1, 30, 40, 50)

	// pegar pedacos de arrayAlvo : dentro de [] indique o intervalo de posicoes que deseja pegar
	pedaco2items := novoArray[1:3]

	ultimaPosicaoDoNovoArray := novoArray[len(novoArray)-1]

	fmt.Println(novoArray)
	fmt.Println(pedaco2items)
	fmt.Println(ultimaPosicaoDoNovoArray) // resposta = 50
}
