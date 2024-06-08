package master_pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/external/use_rzlibs"
)

func MainTesterPessoa() {
	fmt.Println(
		use_rzlibs.UseConvertObjectInJson(Pessoa1Seed), "\n", "\n",
		Pessoa1Seed.Andou(), "\n", "\n",
		Pessoa1Seed.Veiculo.Buzina(), "\n", "\n",
		Pessoa2Seed.Veiculo.Buzina(), "\n", "\n",
	)
}
