package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/contract"
)

type Pessoa struct {
	ID      string
	Nome    string
	Veiculo contract.Veiculo
	// Carro carro.Carro
}

// metodos
func (p *Pessoa) Andou() string {
	return fmt.Sprintf("%s andou ", p.Nome)
}

// func (p *Pessoa) Andou() string {
// 	return fmt.Sprintf("%s andou com o %s de ID: %s do ano %d", p.Nome, p.Veiculo.Fabricante, p.Veiculo.ID, p.Veiculo.Ano)
// }
