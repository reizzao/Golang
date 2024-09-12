package main

import "fmt"

/*
Recurso : Condicionais : DAR AO USER A OPÇÃO DE DAR UM INPUT - QUE ESTEJA ENTRE AS OPCOES DISPONIVEIS EM UM OBJETO > SEM IFS
*/

type ReturnOptions = string

var objLiteralMethods = []ReturnOptions{"UM", "DOIS"}

type ReturnFNGetOption = string

func getOption(objOpcoes []ReturnOptions, posicao string) ReturnFNGetOption {
	var res = ""

	for _, v := range objOpcoes {
		if v == posicao {
			res = v
		}
	}
	return res
}

// AO USAR A FUNCAO E ENVIAR O INPUT - E ESSE INPUT FOR A CHAVE DE UM DOS METODOS DISPONIVEIS O USER TERA O RESULTADO ENTRE OS PRE ESTABELECIDOS.

func main() {
	fmt.Println(getOption(objLiteralMethods, "DOIS"))
}
