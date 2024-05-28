package test

import (
	"fmt"
	"testing"

	"github.com/rzjprogramador/base_golang/internal/entity_func"
)

func TestSoma(t *testing.T) {

	action, _ := entity_func.Soma(2, 1)
	hardcodeOK := 3
	// hardcodeError := errors.New("Ops.. resultado maior que a regra que é 10")
	// TODO: Testar a mensagem de erro em arquivo diferente

	if action != hardcodeOK {
		t.Error("ops..Quebrou...")
	}
	// t.Log(action)
	fmt.Println("o Resultado OK é: ", action)

}
