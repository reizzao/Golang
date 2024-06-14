package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/domain/literals/test"
	"github.com/rzjprogramador/base_golang/AppMotor/external/use_rzlibs"
)

func main() {
	fmt.Println("Estou no Main Root inicial")

	// testers s
	test.MainTesterPessoa()

	// testers libs
	fmt.Println(use_rzlibs.UseSoma(2, 2))
}
