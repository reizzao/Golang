package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/carro"
)

type Pessoa struct {
	ID    string
	Nome  string
	Carro carro.Carro
}

// metodos
func (p Pessoa) Andou() string {
	return fmt.Sprintf("%s andou com o %s de ID: %s do ano %d", p.Nome, p.Carro.Fabricante, p.Carro.ID, p.Carro.Ano)
}
