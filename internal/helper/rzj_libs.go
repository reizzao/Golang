package helper

import (
	"github.com/rzjprogramador/RzLibs_GO/soma"
)

// POLO_UNICO : REPASSA AS LIBS IMPORTADAS : Rzj_Libs

func UseSoma(x int, y int) (int, error) {
	return soma.Soma(x, y)
}
