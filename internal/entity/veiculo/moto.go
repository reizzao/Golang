package veiculo

import "fmt"

type Moto struct {
	ID         string
	Fabricante string
	Modelo     string
	Ano        int
}

func (m *Moto) Buzina() string {
	return fmt.Sprint("buzinou com a MOTO ... >> ", m.ID)
}
