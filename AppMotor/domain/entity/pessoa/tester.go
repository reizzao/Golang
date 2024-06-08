package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/external/use_rzlibs"
)

func MainTesterPessoa() {
	fmt.Println(
		use_rzlibs.UseConvertObjectInJson(Pessoa1Seed),
		Pessoa1Seed.Andou(),
		Pessoa1Seed.Veiculo.Buzina(),
		Pessoa2Seed.Veiculo.Buzina(),
	)
}
