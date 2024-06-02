package pessoa

import "fmt"

// metodos

func (p *Pessoa) Andou() string {
	return fmt.Sprintf("%s andou ", p.Nome)
}
