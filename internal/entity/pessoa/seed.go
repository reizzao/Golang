package pessoa

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/veiculo"
)

var Pessoa1 = Pessoa{
	ID:      "1",
	Nome:    "Reinaldo",
	Veiculo: &veiculo.Moto1,
}

var Pessoa2 = Pessoa{
	ID:      "2",
	Nome:    "Gabriel",
	Veiculo: &veiculo.Carro1,
}

func HardMethodsPessoa() {
	fmt.Println(Pessoa1.Andou())
	fmt.Println(Pessoa1.Veiculo.Buzina())
	fmt.Println(Pessoa2.Veiculo.Buzina())
}
