package seed

import "github.com/rzjprogramador/base_golang/AppMotor/domain/veiculo"

var Carro1Seed = veiculo.Carro{
	ID:         "id_carro_1",
	Fabricante: "Fab Carro",
	Modelo:     "Modelo BMW Carro",
	Ano:        2024,
}

var Moto1Seed = veiculo.Moto{
	ID:         "id_moto_1",
	Fabricante: "Moto1",
	Modelo:     "Modelo BMW Moto",
	Ano:        2024,
}
