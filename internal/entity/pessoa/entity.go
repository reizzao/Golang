package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/veiculo"
)

type Pessoa struct {
	ID      string
	Nome    string
	Veiculo veiculo.Veiculo
}

// metodos
func (p *Pessoa) Andou() string {
	return fmt.Sprintf("%s andou ", p.Nome)
}
