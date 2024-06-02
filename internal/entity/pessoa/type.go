package pessoa

import (
	"github.com/rzjprogramador/base_golang/internal/entity/veiculo"
)

type Pessoa struct {
	ID      string
	Nome    string
	Veiculo veiculo.Veiculo
}
