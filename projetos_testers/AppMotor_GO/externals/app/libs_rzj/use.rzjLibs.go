package libs_rzj

import (
	"github.com/rzjprogramador/RzLibs_GO/App/convert"
	"github.com/rzjprogramador/RzLibs_GO/App/soma"
)

// POLO_UNICO : REPASSA AS LIBS IMPORTADAS : Rzj_Libs

func UseSoma(x int, y int) (int, error) {
	return soma.Soma(x, y)
}

func UseConvertObjectInJson(obj any) string {
	res := convert.ConvertObjectInJson(obj)
	return res
}
