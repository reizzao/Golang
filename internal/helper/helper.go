package helper

import (
	"github.com/rzjprogramador/Libs_Rzj/blob/main/Go_Libs/soma"
	// "/home/rzj/x/_github_rz_/_NO_GITHUB_/SuperLibs_TS/Go_Libs/soma/soma"
	// "github.com/rzjprogramador/base_golang/internal/helper/soma"
)

// POLO_UNICO : HELPER : repassar aqui em funcoes publicas todos os Helpers.

func UseSoma(x int, y int) (int, error) {
	return soma.Soma(x, y)
}
