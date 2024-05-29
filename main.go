package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/pessoa"
	// "github.com/rzjprogramador/base_golang/internal/entity/carro"
	// "github.com/rzjprogramador/base_golang/internal/entity/moto"
)

func main() {
	fmt.Println("Estou no Main Root inicial")
	// internal.MainPackageInternal()

	// Testers Entitys
	pessoa.HardMethodsPessoa()
	// carro.HardMethodsCarro()
	// moto.HardMethodsMoto()
}
