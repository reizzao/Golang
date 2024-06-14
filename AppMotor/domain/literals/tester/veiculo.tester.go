package tester

import (
	"fmt"

	seed "github.com/rzjprogramador/base_golang/AppMotor/domain/literals/seed"
)

func TesterVeiculo() {
	fmt.Println(
		seed.Pessoa1Seed.Andou(),
		seed.Pessoa2Seed.Andou(),
		seed.Pessoa1Seed.Veiculo.Buzina(),
		seed.Pessoa2Seed.Veiculo.Buzina(),
	)
}
