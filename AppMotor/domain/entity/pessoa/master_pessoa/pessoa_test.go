package master_pessoa

import (
	"testing"
)

func TestPessoa(t *testing.T) {
	instance := CreatePessoa(Pessoa1Seed)

	if instance.Nome != "Reinaldo" {
		t.Error("Ops nao criou corretamente o campo Nome de Pessoa")
	}

	if instance.NomeCompleto() != "Reinaldo Z. Junior" {
		t.Error("Ops falha no nome completo de Pessoa1Seed")
	}
}
