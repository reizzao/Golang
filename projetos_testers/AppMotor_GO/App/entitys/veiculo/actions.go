package veiculo

import "fmt"

// carro
func (c *Carro) Buzina() string {
	return fmt.Sprint("buzinou com o CARRO ... >>", c.ID)
}

// moto
func (m *Moto) Buzina() string {
	return fmt.Sprint("buzinou com a MOTO ... >> ", m.ID)
}
