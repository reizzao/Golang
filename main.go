package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/domain/entity/pessoa/master_pessoa"
	"github.com/rzjprogramador/base_golang/external/use_rzlibs"
)

func main() {
	fmt.Println("Estou no Main Root inicial")

	// testers entitys
	master_pessoa.MainTesterPessoa()

	// testers libs
	fmt.Println(use_rzlibs.UseSoma(2, 2))
}
