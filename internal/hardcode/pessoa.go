package hardcode

import (
	"github.com/rzjprogramador/base_golang/internal/entity"
)

func PessoaReinaldo() entity.Pessoa {
	reinaldo := entity.Pessoa{
		ID:    "1",
		Nome:  "Reinaldo",
		Carro: CarroBmw(),
	}
	return reinaldo
}

func AndouReinaldo() string {
	return PessoaReinaldo().Andou()
}