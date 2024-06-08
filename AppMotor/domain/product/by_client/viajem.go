package by_client

import (
	"github.com/rzjprogramador/base_golang/AppMotor/domain/pessoa/components_pessoa"
	"github.com/rzjprogramador/base_golang/AppMotor/domain/product/use_product"
)

type Viajem struct {
	Embarque          EmbarqueDesembarque
	Dembarque         EmbarqueDesembarque
	Passageiro        components_pessoa.Passageiro
	MotoristaParceiro components_pessoa.MotoristaParceiro
	ValorViajem       use_product.ValueViajem
}

type EmbarqueDesembarque struct {
	EnderecoInicio Endereco
	Km_inicio      Km
	Km_Final       Km
	EnderecoFinal  Endereco
}

type Endereco struct {
	Cep        string
	Longadouro string
	Numero     string
	Regiao     string
	Cidade     string
	Estado     string
	UF         string
	Pais       string
}

type Km struct {
	Value int
}
