package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/entity/carro"
	"github.com/rzjprogramador/base_golang/internal/entity/pessoa"
)

func main() {
	fmt.Println("Estou no Main Root inicial")
	// internal.MainPackageInternal()

	// Testers Entitys
	carro.HardMethodsCarro()
	pessoa.HardMethodsPessoa()
}
