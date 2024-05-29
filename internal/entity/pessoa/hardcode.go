package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/carro"
)

var reinaldo = Pessoa{
	ID:    "1",
	Nome:  "Reinaldo",
	Carro: carro.Bmw1,
}

func HardMethodsPessoa() {
	fmt.Println(reinaldo.Andou())
}
