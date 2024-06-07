package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/pessoa"
)

func main() {
	fmt.Println("Estou no Main Root inicial")

	// testers internal
	fmt.Println(pessoa.TesterSeedPessoa1())
	// fmt.Println(internal.MainNomeCompletoPessoa1Seed())
}
