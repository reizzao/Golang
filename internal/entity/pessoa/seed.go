package pessoa

import (
	"github.com/rzjprogramador/base_golang/internal/entity/veiculo"
)

var Pessoa1Seed = Pessoa{
	ID:        "1",
	Nome:      "Reinaldo",
	Sobrenome: "Z. Junior",
	Veiculo:   &veiculo.Moto1Seed,
}

var Pessoa2Seed = Pessoa{
	ID:        "2",
	Nome:      "Gabriel",
	Sobrenome: "Z. Junior",
	Veiculo:   &veiculo.Carro1Seed,
}

// fmt.Println(Pessoa1Seed.Andou())
// fmt.Println(Pessoa1Seed.Veiculo.Buzina())
// fmt.Println(Pessoa2Seed.Veiculo.Buzina())
