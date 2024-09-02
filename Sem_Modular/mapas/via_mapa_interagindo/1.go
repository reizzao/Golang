package main

import "fmt"

var mapa = map[string]string{
	"key1": "valor1",
	"key2": "valor2",
}

func main() {
	// ITERAR SOBRE MAPA
	for _, v := range mapa {
		fmt.Printf("Valores: %s\n", v)
		// fmt.Printf("key[%s] value[%s]\n", k, v)
	}
}
