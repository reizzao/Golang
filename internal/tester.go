package internal

import (
	"github.com/rzjprogramador/base_golang/internal/entity/pessoa"
	"github.com/rzjprogramador/base_golang/internal/helper"
)

func MainNomeCompletoPessoa1Seed() string {
	return pessoa.Pessoa1Seed.NomeCompleto()
}

func MainInternalSoma() (int, error) {
	return helper.UseSoma(3, 3)
}
