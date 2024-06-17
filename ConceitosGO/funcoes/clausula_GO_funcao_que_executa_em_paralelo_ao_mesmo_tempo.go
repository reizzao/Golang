package funcao

import "fmt"

func FuncaoOutraTred() {
	fmt.Println("Sou um Gorotine estou em outra Tred/Saida e vou rodar ao emsmo tempo que qualquer uma.")
}

func Gorotinas() {
	fmt.Println("Sou a Primeira Funcao")

	go func() {
		FuncaoOutraTred()
	}()
}

/*
Conceito: Gorotines são funcoes de rotinas GO que rodam são executadas em paralelo ao mesmo tempo com quaisquer outras que estao sendo executadas no momento, tudo isto em outra Tred basta ser executada que ela vai rodar ao mesmo tempo que outras.

sintaxe: go funcao() {}() //ela só pode ser criada dentro do escopo de uma outra funcao / jamais solta - e seu uso comum é auto executar na sua configuracao colocando () parenteses no seu final

*/
