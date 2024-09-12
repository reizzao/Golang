package main

import "fmt"

type Igeneric interface {
	int | float64
}

type IGeneric_em_Estruturas[T Igeneric | string | any] struct {
	Data T
}

func SomaGenerica[T Igeneric](i T) T {
	var res T = i + 1
	return res
}

func GenericEmEstruturas() {
	numeros := IGeneric_em_Estruturas[int]{Data: 10}
	textos := IGeneric_em_Estruturas[string]{Data: "Foo"}

	fmt.Println(numeros.Data) // [1 2 3]
	fmt.Println(textos.Data)  // [1 2 3]

	// passing string as type parameter
	// modelStr := Model[string]{Data: []string{"a", "b", "c"}}
	// fmt.Println(modelStr.Data)
}

func main() {
	// fmt.Println(SomaGenerica(22.22))
	GenericEmEstruturas()
}


