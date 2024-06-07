package veiculo

import "fmt"

type Carro struct {
	ID         string
	Fabricante string
	Modelo     string
	Ano        int
}

func (c *Carro) Buzina() string {
	return fmt.Sprint("buzinou com o CARRO ... >> ", c.ID)
}
