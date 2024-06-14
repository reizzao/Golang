package product

import "github.com/rzjprogramador/base_golang/AppMotor/domain/pessoa"

type Viajem struct {
	Embarque          EmbarqueDesembarque
	Dembarque         EmbarqueDesembarque
	Passageiro        pessoa.Passageiro
	MotoristaParceiro pessoa.MotoristaParceiro
	ValorViajem       ValueViajem
}

type EmbarqueDesembarque struct {
	EnderecoInicio Endereco
	Km_inicio      Km
	Km_Final       Km
	EnderecoFinal  Endereco
}

type ValueViajem struct {
	Value float64
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
