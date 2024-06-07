package pessoa

import (
	"github.com/rzjprogramador/base_golang/app1/domain/entity/veiculo"
)

type Pessoa struct {
	ID        string           `json: id`
	Nome      string           `json: nome`
	Sobrenome string           `json: sobrenome`
	Veiculo   veiculo.IVeiculo `json: veiculo`
}

type IPessoa interface {
	NomeCompleto() string
}
