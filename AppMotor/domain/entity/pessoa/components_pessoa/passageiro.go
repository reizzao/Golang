package components_pessoa

import (
	"github.com/rzjprogramador/base_golang/AppMotor/domain/entity/pessoa/master_pessoa"
	"github.com/rzjprogramador/base_golang/AppMotor/domain/entity/pessoa/use_pessoa"
)

type Passageiro struct{
	Pessoa master_pessoa.Pessoa
	Cargo use_pessoa.Cargo
}