package golpe

import "github.com/cirobispo/sandbox/internal/comum/golpes"

type VerificarAcaoCondicional func(listaDeGolpes []golpes.Golpear) golpes.Reacao

type Golpe struct {
	acao         golpes.Acao
	reacao       golpes.Reacao
	executaTurno bool
}

func (g Golpe) Acao() golpes.Acao {
	return g.acao
}

func (g Golpe) Reacao() golpes.Reacao {
	return g.reacao
}

func (g Golpe) ExecutaTurno() bool {
	return g.executaTurno
}

func NewGolpe(t golpes.Acao, acao golpes.Reacao, executaTurno bool) Golpe {
	return Golpe{
		acao:         t,
		reacao:       acao,
		executaTurno: executaTurno,
	}
}

func NewFootFault() Golpe {
	result := NewGolpe(golpes.HTFootFault, golpes.TACondicional, false)
	return result
}

func NewAce() Golpe {
	return NewGolpe(golpes.HTAce, golpes.TAEncerrarPLC, false)
}

func NewServicoFora() Golpe {
	result := NewGolpe(golpes.HTServeOut, golpes.TACondicional, false)
	return result
}

func NewServicoDentro() Golpe {
	return NewGolpe(golpes.HTServeIn, golpes.TAProsseguir, true)
}

func NewLET() Golpe {
	return NewGolpe(golpes.HTServeLet, golpes.TANulo, false)
}

func NewServicoNaRede() Golpe {
	result := NewGolpe(golpes.HTServeNet, golpes.TACondicional, false)
	return result
}

func NewRetornoNaRede() Golpe {
	return NewGolpe(golpes.HTReturnNet, golpes.TAEncerrarPLO, false)
}

func NewRetornoDentro() Golpe {
	return NewGolpe(golpes.HTReturnIn, golpes.TAProsseguir, true)
}

func NewRetornoFora() Golpe {
	return NewGolpe(golpes.HTReturnOut, golpes.TAEncerrarPLO, false)
}

func NewDuplaFalta() Golpe {
	return NewGolpe(golpes.HTDoubleFault, golpes.TAEncerrarPLO, false)
}

func NewDevolveuNaRede() Golpe {
	return NewGolpe(golpes.HTNet, golpes.TAEncerrarPLO, false)
}

func NewDevolveuDentro() Golpe {
	return NewGolpe(golpes.HTIn, golpes.TAProsseguir, false)
}

func NewWinner() Golpe {
	// Aqui será preciso avaliar como atuar aqui.
	// Se o winner é confirmado depois ou no momento do golpe.
	// Provavelmente deverá ser depois. E daí, ficaria, devolveu na rede, fora, winner adversario.
	return NewGolpe(golpes.HTWinner, golpes.TAEncerrarPLC, false)
}

func NewDevolveuFora() Golpe {
	return NewGolpe(golpes.HTOut, golpes.TAEncerrarPLO, false)
}

func NewNaoTocou() Golpe {
	return NewGolpe(golpes.HTMiss, golpes.TAEncerrarPLO, false)
}

func NewQueimou() Golpe {
	return NewGolpe(golpes.HTToast, golpes.TAEncerrarPLO, false)
}

func NewToqueNaRede() Golpe {
	return NewGolpe(golpes.HTNetTouch, golpes.TAEncerrarPLO, false)
}
