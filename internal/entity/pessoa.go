package entity

import "fmt"

type Pessoa struct {
	ID    string
	Nome  string
	Carro Carro
}

// metodos
func (p Pessoa) Andou() string {
	return fmt.Sprintf("%s andou com o carro %s", p.Nome, p.Carro.Fabricante)
}
