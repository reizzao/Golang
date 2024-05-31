package main

import (
	"fmt"

	"github.com/rzjprogramador/base_golang/internal/helper"
)

func main() {
	fmt.Println("Estou no Main Root inicial")
	// internal.MainPackageInternal()

	// Testers Entitys
	// pessoa.HardMethodsPessoa()

	// Tester helpers
	// helper.HardTester_Soma()
	fmt.Println(helper.UseSoma(3, 3))
}
