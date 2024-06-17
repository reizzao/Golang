package tester

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/domain/pessoa"
	"github.com/rzjprogramador/base_golang/AppMotor/external/libs"
)

func TesterPessoa() {
	fmt.Println(
		libs.UseConvertObjectInJson(pessoa.Pessoa1Seed),
		// prototypes - computados
		pessoa.Pessoa1Seed.NomeCompleto(),
		pessoa.Pessoa2Seed.NomeCompleto(),
	)

}
