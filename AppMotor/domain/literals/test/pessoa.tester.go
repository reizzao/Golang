package test

import (
	"fmt"

	seed "github.com/rzjprogramador/base_golang/AppMotor/domain/literals/seed"
	"github.com/rzjprogramador/base_golang/AppMotor/external/use_rzlibs"
)

func MainTesterPessoa() {
	fmt.Println(
		use_rzlibs.UseConvertObjectInJson(seed.Pessoa1Seed), "\n", "\n",
		seed.Pessoa1Seed.Andou(), "\n", "\n",
		seed.Pessoa1Seed.Veiculo.Buzina(), "\n", "\n",
		seed.Pessoa2Seed.Veiculo.Buzina(), "\n", "\n",
	)
}
