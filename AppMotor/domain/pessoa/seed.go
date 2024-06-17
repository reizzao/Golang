package pessoa

import "github.com/rzjprogramador/base_golang/AppMotor/domain/veiculo"

var Pessoa1Seed = Pessoa{
	ID:        "1",
	Nome:      "Reinaldo",
	Sobrenome: "Z. Junior",
	Veiculo:   &veiculo.Moto1Seed,
}

var Pessoa2Seed = Pessoa{
	ID:        "2",
	Nome:      "Gabriel",
	Sobrenome: "Arthur",
	Veiculo:   &veiculo.Carro1Seed,
}
