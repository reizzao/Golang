package pessoa

import "github.com/rzjprogramador/AppMotor_GO/App/entitys/veiculo"

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
