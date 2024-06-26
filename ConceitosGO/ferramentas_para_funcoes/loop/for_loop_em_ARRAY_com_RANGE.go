package loop

import "fmt"

func Loop_Em_Array_Com_Range() string {
	arrayNomes := []string{"Rei", "Guga", "Leo"}
	vazio := ""

	for _, valores := range arrayNomes {
		fmt.Println(valores)
		vazio += fmt.Sprint(" -- ", valores)
	}

	return vazio

}
