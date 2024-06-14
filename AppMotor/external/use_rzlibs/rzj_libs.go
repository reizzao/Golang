package use_rzlibs

import (
	"github.com/rzjprogramador/RzLibs_GO/convert"
	"github.com/rzjprogramador/RzLibs_GO/soma"
)

// POLO_UNICO : REPASSA AS LIBS IMPORTADAS : Rzj_Libs

func UseSoma(x int, y int) (int, error) {
	return soma.Soma(x, y)
}

func UseConvertObjectInJson(obj any) string {
	res := convert.ConvertObjectInJson(obj)
	return res
}
