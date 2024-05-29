package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/veiculo"
)

var reinaldo = Pessoa{
	ID:      "1",
	Nome:    "Reinaldo",
	Veiculo: &veiculo.MotoBmw1,
}

func HardMethodsPessoa() {
	fmt.Println(reinaldo.Andou())
	fmt.Println(reinaldo.Veiculo.MotoBmw1.buzina()) // TODO
	// fmt.Println(reinaldo.Veiculo.buzina()) // TODO
}
