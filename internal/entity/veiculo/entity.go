package veiculo

import "fmt"

type Veiculo interface {
	buzina() string
}

// --
type Moto struct {
	ID         string
	Fabricante string
	Modelo     string
	Ano        int
}

func (m *Moto) buzina() string {
	return fmt.Sprint("%s buzinou com a MOTO ...", m.ID)
}

// --
type Carro struct {
	ID         string
	Fabricante string
	Modelo     string
	Ano        int
}

func (c *Carro) buzina() string {
	return fmt.Sprint("%s buzinou com o CARRO ...", c.ID)
}
