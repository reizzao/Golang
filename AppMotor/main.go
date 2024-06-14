package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/domain/literals/tester"
	"github.com/rzjprogramador/base_golang/AppMotor/external/use_rzlibs"
)

func main() {
	fmt.Println(" ========= Main do AppMotor ========= .")

	fmt.Println(" ------- VEICULO ------- ")
	tester.TesterVeiculo()

	fmt.Println(" ------- PESSOA ------- ")
	tester.TesterPessoa()

	fmt.Println(" ------- LIBS_TESTER ------- ")
	fmt.Println(use_rzlibs.UseSoma(2, 2))
}
