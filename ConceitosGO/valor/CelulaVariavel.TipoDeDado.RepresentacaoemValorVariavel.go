package valor

import (
	"fmt"
	"strings"
)

func Variavel_Tipo_ValorRepresentacao() {
	numero_Inteiro := 10
	texto_Conjunto_de_Caracteres := "Sou um texto"
	logico := true

	tamanhoTexto := len(texto_Conjunto_de_Caracteres)
	textoModificado := strings.ReplaceAll(texto_Conjunto_de_Caracteres, texto_Conjunto_de_Caracteres, "texto Mudado")

	fmt.Println(tamanhoTexto, textoModificado, numero_Inteiro, logico)
}

/**/
