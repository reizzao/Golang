package valor

import "fmt"

type FooDefaults struct {
	C1 string
	C2 string
}

var Input_C1_Default = "valueC1"
var Input_C2_Default = "valueC2"

func Tester_valorDefault_Hardcode() {
	input := FooDefaults{}
	input.C1 = Input_C1_Default
	input.C2 = Input_C2_Default

	res := input

	fmt.Println(res)
}

/*
Conceito: Valor Default em campo de novos objetos
- crio o valorDefault_Hardcode e uso na funcao que vai criar novos objetos.
- pego o campo alvo que quero o default valor e atribuo o valorDefault_Hardcode
- no final devolvo o objeto modificado.
*/
