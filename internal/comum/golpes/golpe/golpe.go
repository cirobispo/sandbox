package golpe

import "github.com/cirobispo/sandbox/internal/comum/golpes"

type VerificarAcaoCondicional func(listaDeGolpes []golpes.TipoAcaoGolpe) golpes.TipoAcao

type Golpe struct {
	tipoGolpe golpes.TipoDoGolpe
	tipoAcao  golpes.TipoAcao
}

func (g Golpe) Tipo() golpes.TipoDoGolpe {
	return g.tipoGolpe
}

func (g Golpe) Acao() golpes.TipoAcao {
	return g.tipoAcao
}

func NewGolpe(t golpes.TipoDoGolpe, acao golpes.TipoAcao) Golpe {
	return Golpe{
		tipoGolpe: t,
		tipoAcao:  acao,
	}
}

func NewFootFault() Golpe {
	result := NewGolpe(golpes.HTFootFault, golpes.TACondicional)
	return result
}

func NewAce() Golpe {
	return NewGolpe(golpes.HTAce, golpes.TAEncerrarPLC)
}

func NewServicoFora() Golpe {
	result := NewGolpe(golpes.HTServeOut, golpes.TACondicional)
	return result
}

func NewServicoDentro() Golpe {
	return NewGolpe(golpes.HTServeIn, golpes.TAProsseguir)
}

func NewLET() Golpe {
	return NewGolpe(golpes.HTServeLet, golpes.TANulo)
}

func NewServicoNaRede() Golpe {
	result := NewGolpe(golpes.HTServeNet, golpes.TACondicional)
	return result
}

func NewRetornoNaRede() Golpe {
	return NewGolpe(golpes.HTReturnNet, golpes.TAEncerrarPLO)
}

func NewRetornoDentro() Golpe {
	return NewGolpe(golpes.HTReturnIn, golpes.TAProsseguir)
}

func NewRetornoFora() Golpe {
	return NewGolpe(golpes.HTReturnOut, golpes.TAEncerrarPLO)
}

func NewDuplaFalta() Golpe {
	return NewGolpe(golpes.HTDoubleFault, golpes.TAEncerrarPLO)
}

func NewGolpeuNaRede() Golpe {
	return NewGolpe(golpes.HTNet, golpes.TAEncerrarPLO)
}

func NewDevolveuDentro() Golpe {
	return NewGolpe(golpes.HTIn, golpes.TAProsseguir)
}

func NewWinner() Golpe {
	// Aqui será preciso avaliar como atuar aqui.
	// Se o winner é confirmado depois ou no momento do golpe.
	// Provavelmente deverá ser depois. E daí, ficaria, devolveu na rede, fora, winner adversario.
	return NewGolpe(golpes.HTWinner, golpes.TAEncerrarPLC)
}

func NewGolpeuFora() Golpe {
	return NewGolpe(golpes.HTOut, golpes.TAEncerrarPLO)
}

func NewNaoTocou() Golpe {
	return NewGolpe(golpes.HTMiss, golpes.TAEncerrarPLO)
}

func NewQueimou() Golpe {
	return NewGolpe(golpes.HTToast, golpes.TAEncerrarPLO)
}

func NewToqueNaRede() Golpe {
	return NewGolpe(golpes.HTNetTouch, golpes.TAEncerrarPLO)
}
