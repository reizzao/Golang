package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/AppMotor/domain/veiculo"
)

type Pessoa struct {
	ID        string           `json: id`
	Nome      string           `json: nome`
	Sobrenome string           `json: sobrenome`
	Veiculo   veiculo.IVeiculo `json: veiculo`
}

// prototypes
type IPessoa interface {
	NomeCompleto() string
}

// metodos
func (p *Pessoa) NomeCompleto() string {
	return fmt.Sprintf("%s %s", p.Nome, p.Sobrenome)
}

func (p *Pessoa) Andou() string {
	return fmt.Sprintf("%s andou ", p.Nome)
}

// handlers
func CreatePessoa(p Pessoa) *Pessoa {
	return &p
}
