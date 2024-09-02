package main

import "fmt"

var myMapa = map[string]string{
	"UM":   "valor1",
	"DOIS": "valor2",
}

func main() {
	//
	res := myMapa
	fmt.Println(res)
}

/*
recurso: OBJETO MAPA DEVOLVE UM OBJETO COM CHAVE E VALOR DO EMSMO TIPO PROMETIDO NA DECLARACAO
- declarar mapa : variavel = map[tipoValorChave]tipoValorDoValordaChave
exemplo : var myMapa = map[string]string{ // aqui as chaves e valores de acordo com o tipo prometido apra cada uma }

*/
