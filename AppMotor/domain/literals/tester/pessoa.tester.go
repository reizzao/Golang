package tester

import (
	"fmt"

	seed "github.com/rzjprogramador/base_golang/AppMotor/domain/literals/seed"
	"github.com/rzjprogramador/base_golang/AppMotor/external/use_rzlibs"
)

func TesterPessoa() {
	fmt.Println(
		use_rzlibs.UseConvertObjectInJson(seed.Pessoa1Seed),
		// prototypes - computados
		seed.Pessoa1Seed.NomeCompleto(),
		seed.Pessoa2Seed.NomeCompleto(),
	)

}
