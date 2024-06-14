package seed

import (
	"github.com/rzjprogramador/base_golang/AppMotor/domain/pessoa"
	"github.com/rzjprogramador/base_golang/AppMotor/domain/veiculo"

)

var Pessoa1Seed = pessoa.Pessoa{
	ID:        "1",
	Nome:      "Reinaldo",
	Sobrenome: "Z. Junior",
	Veiculo:   &veiculo.Moto1Seed,
}

var Pessoa2Seed = pessoa.Pessoa{
	ID:        "2",
	Nome:      "Gabriel",
	Sobrenome: "Z. Junior",
	Veiculo:   &veiculo.Carro1Seed,
}
