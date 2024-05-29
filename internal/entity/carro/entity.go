package carro

type Carro struct {
	ID         string
	Fabricante string
	Modelo     string
	Ano        int
}

func CreateCarro(c Carro) Carro {
	return c
}
