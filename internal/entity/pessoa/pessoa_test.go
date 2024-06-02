package pessoa

import (
	"testing"
)

func TestPessoa(t *testing.T) {
	instance := CreatePessoa(Pessoa1)

	if instance.Nome != "Reinaldo" {
		t.Error("Ops nao criou corretamente o campo Nome de Pessoa")
	}
}
