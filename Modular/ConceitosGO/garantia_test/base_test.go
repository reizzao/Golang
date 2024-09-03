package garantia_test

import (
	"testing"
)

// Imports
// libs
func MessageErrorTest() string {
	return "\n ***********RESPONSE TEST *************\n Ops... Esperado: %t --> Tentativa: %t \n*********** FINAL RESPONSE TEST ************* "
}

// funcao Sut
func Import_Sut_FuncaoUseCase_Target(i ResInput1) ResSut {
	return i
}

//

type ResSut = int
type ResInput1 = int

var auxImport_InputTester ResInput1 = 1000

var sut ResSut = Import_Sut_FuncaoUseCase_Target(auxImport_InputTester)

func Test_Entity(t *testing.T) {
	expect_request := sut
	var compare_request ResSut = 1000

	/* -- Suites -- */

	// SUITE :: TARGET: Test #TODO - TITULO: deve retornar #TODO
	if sut != compare_request {
		t.Error(MessageErrorTest(), expect_request, compare_request)
	}
}

/*
Test fluxo
	   - crie o tipo do response do sut : ex: ResSut = tipo
	   - crie a var auxiliar global no test , chamada sut , para ser a funcionalidade principal a ser testada no test.
	   - crie suite de test conforme pede a linguagem : em golang tem que comecar com Test_NOmeEntidade( e receber parametro da lib de test ex: t *testing.T )
	   - dentro crie a varAuxiliar expect que sera o resultado esperado no test,
	   - e tambem a varAuxiliar compare que sera o resultado que voce queira comparar com o resultado do sut,
	   - faça a condicional usando operadores logicos para definir o que fazer comparando sut e o que voce quer comparar
	   - por consequencia lance um erro da linguagem usando o parametro de test , para caso nao seja o esperado exploda um erro parando o test.
*/
