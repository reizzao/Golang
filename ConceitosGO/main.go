package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/ConceitosGO/condicionais"
	funcao "github.com/rzjprogramador/base_golang/ConceitosGO/funcoes"
)

func main() {
	fmt.Println("===== Estou na main do ConceitosGO =====")

	fmt.Println("===== FUNCAO GOROTINES =====")
	funcao.Gorotinas()

	fmt.Println("===== FUNCAO DEFER =====")
	funcao.FuncaoDefer_EsperaTodas()

	fmt.Println("===== SELECT =====")
	condicionais.UseSelect()
}
