package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/external/libs"
	"github.com/rzjprogramador/base_golang/AppMotor/testers"
)

func main() {
	fmt.Println(" ========= Main do AppMotor ========= .")

	fmt.Println(" ------- VEICULO ------- ")
	testers.TesterVeiculo()

	fmt.Println(" ------- PESSOA ------- ")
	testers.TesterPessoa()

	fmt.Println(" ------- LIBS_TESTER ------- ")
	fmt.Println(libs.UseSoma(2, 2))
}
