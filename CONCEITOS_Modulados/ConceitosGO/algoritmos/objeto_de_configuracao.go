package algoritmos

type ObjectoDeConfiguracao struct {
	C1 string
	C2 string
}

var ObjetoRequest = ObjectoDeConfiguracao{
	C1: "a",
	C2: "",
}

func configurandoObj(o ObjectoDeConfiguracao) ObjectoDeConfiguracao {

	if len(o.C1) <= 0 {
		o.C1 = "valorFIXO_1"
	}
	if len(o.C2) <= 0 {
		o.C2 = "valorFIXO_2"
	}
	// Tatica : Se o que receber do campo for { Menor ou Igual a zero } nada foi enviado entao este campo tera o valor que vou confgiurar na consequencia.

	res := ObjectoDeConfiguracao{
		C1: o.C1,
		C2: o.C2,
	}
	return res
}

func EditandoObjetoDeConfiguracao() ObjectoDeConfiguracao {
	res := configurandoObj(ObjetoRequest)
	return res
}
