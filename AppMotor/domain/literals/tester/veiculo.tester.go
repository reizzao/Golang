package tester

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/domain/pessoa"
)

func TesterVeiculo() {
	fmt.Println(
		pessoa.Pessoa1Seed.Andou(),
		pessoa.Pessoa2Seed.Andou(),
		pessoa.Pessoa1Seed.Veiculo.Buzina(),
		pessoa.Pessoa2Seed.Veiculo.Buzina(),
	)
}
