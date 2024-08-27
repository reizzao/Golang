package main

import "fmt"

type Foo2 struct {
	C1 int
	C2 int
}

var Bar = Foo2{C1: 10, C2: 20}

func main() {

	// iniciar uma variavel e preencher valor depois : usar clausula var variavelalvo e o tipo sem atribuir
	var v Foo2

	// preenchendo depois com sinal atribuicao normal
	v = Bar

	fmt.Println(v)

}
