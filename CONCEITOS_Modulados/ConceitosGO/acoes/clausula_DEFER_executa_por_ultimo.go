package funcao

import "fmt"

func adiada() {
	fmt.Println("EU SOU A ADIADA - MARCADA COM DEFER- VOU SER EXECUTADA POR ULTIMO - NÃO IMPORTA A ORDEM QUE ELA SEJA USADA")
}
func primeira() {
	fmt.Println("SOU A PRIMEIRA")
}
func outra() {
	fmt.Println("SOU OUTRA")
}

func FuncaoDefer_EsperaTodas() {
	defer adiada()
	primeira()
	outra()
}

/*
Significado : clausula: defer = VOU SER EXECUTADA POR ULTIMO - NÃO IMPORTA A ORDEM QUE ELA SEJA USADA

Conceito: é como um else espera todas serem executadas e por ultimo executa a que esta demarcada como defer que quer dizer adiar


*/
