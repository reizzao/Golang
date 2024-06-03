package pessoa

import (
	"github.com/rzjprogramador/base_golang/internal/entity/veiculo"
)

type Pessoa struct {
	ID      string
	Nome    string
	Sobrenome string
	Veiculo veiculo.IVeiculo
}

type IPessoa interface {
	NomeCompleto() string
}
