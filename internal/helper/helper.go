package helper

import (
	"github.com/rzjprogramador/Libs_Rzj/Go_Libs/soma"
)

// POLO_UNICO : REPASSA AS LIBS IMPORTADAS

func UseSoma(x int, y int) (int, error) {
	return soma.Soma(x, y)
}
