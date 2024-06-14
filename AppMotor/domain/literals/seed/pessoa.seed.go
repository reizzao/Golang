package seed

import (
	"github.com/rzjprogramador/base_golang/AppMotor/domain/pessoa"
)

var Pessoa1Seed = pessoa.Pessoa{
	ID:        "1",
	Nome:      "Reinaldo",
	Sobrenome: "Z. Junior",
	Veiculo:   &Moto1Seed,
}

var Pessoa2Seed = pessoa.Pessoa{
	ID:        "2",
	Nome:      "Gabriel",
	Sobrenome: "Arthur",
	Veiculo:   &Carro1Seed,
}
