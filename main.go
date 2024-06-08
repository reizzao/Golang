package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/domain/entity/pessoa"
	"github.com/rzjprogramador/base_golang/external/use_rzlibs"
)

func main() {
	fmt.Println("Estou no Main Root inicial")

	// testers
	pessoa.MainTesterPessoa()
	fmt.Println(use_rzlibs.UseSoma(2, 2))
	// fmt.Println(app1/domain.MainNomeCompletoPessoa1Seed())
}
