package garantia_test

import (
	"testing"
)

var resultado_certo = 1000

func funcaoalvo() int {
	return resultado_certo
}

func TestPessoa(t *testing.T) {
	sut := funcaoalvo()
	esperado := resultado_certo
	comparando_com_resultado_certo := 1000

	if sut != comparando_com_resultado_certo {
		t.Error("Ops... Esperado: ", esperado, "Tentativa: ", comparando_com_resultado_certo)
	}

	// Mais tests
	// if testado != 9 {
	// 	t.Error("Ops...")
	// }
}
