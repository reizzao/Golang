package veiculo

// main
type IVeiculo interface {
	Buzina() string
}

// components
type Carro struct {
	ID         string
	Fabricante string
	Modelo     string
	Ano        int
}

type Moto struct {
	ID         string
	Fabricante string
	Modelo     string
	Ano        int
}
