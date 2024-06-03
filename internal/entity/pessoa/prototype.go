package pessoa

import "fmt"

// metodos

func (p *Pessoa) NomeCompleto() string {
	return fmt.Sprintf("%s %s", p.Nome, p.Sobrenome)
}

func (p *Pessoa) Andou() string {
	return fmt.Sprintf("%s andou ", p.Nome)
}
