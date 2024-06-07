package pessoa

import (
	"github.com/rzjprogramador/base_golang/internal/entity/veiculo"
	"github.com/rzjprogramador/base_golang/internal/helper"
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

func TesterSeedPessoa1() any {
	obj := helper.UseConvertObjectInJson(Pessoa1Seed)
	return obj
}

// fmt.Println(Pessoa1Seed.Andou())
// fmt.Println(Pessoa1Seed.Veiculo.Buzina())
// fmt.Println(Pessoa2Seed.Veiculo.Buzina())
