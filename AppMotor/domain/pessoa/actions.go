package pessoa

import "fmt"

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
