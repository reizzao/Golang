package hardcode

import (
	"github.com/rzjprogramador/base_golang/internal/entity"
)

func CarroBmw() entity.Carro {
	bmw := entity.Carro {
		Fabricante: "Bmw",
		Modelo: "Modelo BMW",
		Ano: 2024,
	}
	return bmw
}