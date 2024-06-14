package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/ConceitosGO/condicionais"
	funcao "github.com/rzjprogramador/base_golang/ConceitosGO/funcoes"
	"github.com/rzjprogramador/base_golang/ConceitosGO/loop"
)

func main() {
	fmt.Println("===== Estou na main do ConceitosGO =====")

	fmt.Println("===== LOOP FOR =====")
	// fmt.Println(loop.For_Loop_Base())
	// loop.Loop_Infinito() // somente testar - pra nao ficar em loop o console
	// loop.Loop_For_Contador_Na_Mesma_Linha()
	fmt.Println(loop.Loop_Em_Array_Com_Range())

	fmt.Println("===== FUNCAO GOROTINES =====")
	funcao.Gorotinas()

	fmt.Println("===== FUNCAO DEFER =====")
	funcao.FuncaoDefer_EsperaTodas()

	fmt.Println("===== SELECT =====")
	condicionais.UseSelect()
}
