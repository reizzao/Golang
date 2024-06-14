package test

import (
	"testing"

	seed "github.com/rzjprogramador/base_golang/AppMotor/domain/literals/seed"
	"github.com/rzjprogramador/base_golang/AppMotor/domain/pessoa"
)

func TestPessoa(t *testing.T) {
	instance := pessoa.CreatePessoa(seed.Pessoa1Seed)

	if instance.Nome != "Reinaldo" {
		t.Error("Ops nao criou corretamente o campo Nome de Pessoa")
	}

	if instance.NomeCompleto() != "Reinaldo Z. Junior" {
		t.Error("Ops falha no nome completo de Pessoa1Seed")
	}
}
