package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/veiculo"
)

var reinaldo = Pessoa{
	ID:      "1",
	Nome:    "Reinaldo",
	Veiculo: &veiculo.Moto1,
}

var gabriel = Pessoa{
	ID:      "2",
	Nome:    "Gabriel",
	Veiculo: &veiculo.Carro1,
}

func HardMethodsPessoa() {
	fmt.Println(reinaldo.Andou())
	fmt.Println(reinaldo.Veiculo.Buzina())
	fmt.Println(gabriel.Veiculo.Buzina())
}
