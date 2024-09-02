package main

import "fmt"

/*
Recurso : Condicionais : DAR AO USER A OPÇÃO DE DAR UM INPUT - QUE ESTEJA ENTRE AS OPCOES DISPONIVEIS EM UM OBJETO > SEM IFS
*/

type ObjOptions struct {
	UM   ReturnOptions
	DOIS ReturnOptions
}
type ReturnOptions = string

// type ReturnOptions = func() string

var objOptions = ObjOptions{
	UM:   "UM",
	DOIS: "DOIS",
}

type ReturnFNGetOption = string

// func objOptions() ObjOptions {
// 	um := func() ReturnOptions {
// 		return "UM"
// 	}()
// 	dois := func() ReturnOptions {
// 		return "DOIS"
// 	}()
// 	res := ObjOptions{
// 		UM:   um,
// 		DOIS: dois,
// 	}
// 	return res
// }

func getOption(objOpcoes ObjOptions, posicao string) ReturnFNGetOption {
	res := objOpcoes[posicao]
	return res
}

// AO USAR A FUNCAO E ENVIAR O INPUT - E ESSE INPUT FOR A CHAVE DE UM DOS METODOS DISPONIVEIS O USER TERA O RESULTADO ENTRE OS PRE ESTABELECIDOS.

func main() {
	fmt.Println(getOption(objOptions, "UM"))
}
