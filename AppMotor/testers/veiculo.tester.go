package testers

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/entitys/pessoa"
)

func TesterVeiculo() {
	fmt.Println(
		pessoa.Pessoa1Seed.Andou(),
		pessoa.Pessoa2Seed.Andou(),
		pessoa.Pessoa1Seed.Veiculo.Buzina(),
		pessoa.Pessoa2Seed.Veiculo.Buzina(),
	)
}
