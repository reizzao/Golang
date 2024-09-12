package garantia_test

import (
	"testing"

)

// NO PROJETO REAL SERÃO IMPORTADOS

type ResFuncTarget = int

func funcaoTarget(i int) ResFuncTarget { return i }

//

// USE OFICIAL NO TEST

type ResSut = ResFuncTarget

var (
	sut        ResSut = funcaoTarget(10)
	inputTest1        = 10
)

func Test_Entity(t *testing.T) {
	expected_request_ := sut
	compare_request_ := inputTest1

	/* -- Suites -- */

	// Test: [ TODO_NameTest] : deve retornar TODO { TODO }
	if sut != compare_request_ {
		t.Error(MessageErrorTest(expected_request_, compare_request_))
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
