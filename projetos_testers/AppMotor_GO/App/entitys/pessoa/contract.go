package pessoa

import "github.com/rzjprogramador/AppMotor_GO/App/entitys/veiculo"

// main
type Pessoa struct {
	ID        string           `json: id`
	Nome      string           `json: nome`
	Sobrenome string           `json: sobrenome`
	Veiculo   veiculo.IVeiculo `json: veiculo`
}

// helpers
type Cargo struct {
	Value string
}

// components
type Passageiro struct {
	Pessoa Pessoa
	Cargo  Cargo
}

type MotoristaParceiro struct {
	Pessoa Pessoa
	Cargo  Cargo
}
