package helper

import (
	"github.com/rzjprogramador/base_golang/internal/helper/soma"
)

// POLO_UNICO : HELPER : repassar aqui em funcoes publicas todos os Helpers.

func UseSoma(x int, y int) (int, error) {
	return soma.Soma(x, y)
}
