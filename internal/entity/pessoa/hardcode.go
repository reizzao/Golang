package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/moto"
)

var reinaldo = Pessoa{
	ID:   "1",
	Nome: "Reinaldo",
	// Veiculo: moto.MotoBmw1,
	Veiculo: moto.MotoBmw1,
}

func HardMethodsPessoa() {
	// fmt.Println(reinaldo.Andou())
	fmt.Println(reinaldo.Veiculo.buzina()) // TODO
}
