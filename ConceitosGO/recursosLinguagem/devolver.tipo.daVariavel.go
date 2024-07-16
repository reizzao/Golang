package recursosLinguagem

import (
	"fmt"
	"reflect"
)

func Tester_RecursosLinguagem() {
	texto := "foo"
	resultTipo := reflect.TypeOf(texto)
	fmt.Println(resultTipo)
}