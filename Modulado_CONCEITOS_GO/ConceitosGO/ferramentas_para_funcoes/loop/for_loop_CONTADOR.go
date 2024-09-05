package loop

import (
	"fmt"
	"time"
)

func Loop_For_Contador_Na_Mesma_Linha() {
	// na mesma linha : valor; condicao; mudarValor-Pra_mudar_a_condicao

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println(j)
	}

}
