package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/app1/domain/entity/pessoa"
	"github.com/rzjprogramador/base_golang/external/use_rzlibs"
)

func main() {
	fmt.Println("Estou no Main Root inicial")

	// testers
	fmt.Println(pessoa.TesterSeedPessoa1())
	fmt.Println(use_rzlibs.UseSoma(2, 2))
	// fmt.Println(app1/domain.MainNomeCompletoPessoa1Seed())
}
