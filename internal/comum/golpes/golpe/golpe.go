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
	result := NewGolpe(golpes.APeNaQuadra, golpes.RPCondicionado, false)
	return result
}

func NewAce() Golpe {
	return NewGolpe(golpes.AAce, golpes.REPLadoCorrente, false)
}

func NewServicoFora() Golpe {
	result := NewGolpe(golpes.AServicoFora, golpes.RPCondicionado, false)
	return result
}

func NewServicoDentro() Golpe {
	return NewGolpe(golpes.AServicoDentro, golpes.RProsseguir, true)
}

func NewLET() Golpe {
	return NewGolpe(golpes.ALet, golpes.RNulo, false)
}

func NewServicoNaRede() Golpe {
	result := NewGolpe(golpes.AServicoNaRede, golpes.RPCondicionado, false)
	return result
}

func NewRetornoNaRede() Golpe {
	return NewGolpe(golpes.ARetornoNaRede, golpes.REPLadoOposto, false)
}

func NewRetornoDentro() Golpe {
	return NewGolpe(golpes.ARetornoDentro, golpes.RProsseguir, true)
}

func NewRetornoFora() Golpe {
	return NewGolpe(golpes.ARetornoFora, golpes.REPLadoOposto, false)
}

func NewDuplaFalta() Golpe {
	return NewGolpe(golpes.ADuplaFalta, golpes.REPLadoOposto, false)
}

func NewDevolveuNaRede() Golpe {
	return NewGolpe(golpes.ADevolucaoNaRede, golpes.REPLadoOposto, false)
}

func NewDevolveuDentro() Golpe {
	return NewGolpe(golpes.ADevolucaoDentro, golpes.RProsseguir, false)
}

func NewWinner() Golpe {
	// Aqui será preciso avaliar como atuar aqui.
	// Se o winner é confirmado depois ou no momento do golpe.
	// Provavelmente deverá ser depois. E daí, ficaria, devolveu na rede, fora, winner adversario.
	return NewGolpe(golpes.AWinner, golpes.REPLadoCorrente, false)
}

func NewDevolveuFora() Golpe {
	return NewGolpe(golpes.ADevolucaoFora, golpes.REPLadoOposto, false)
}

func NewNaoTocou() Golpe {
	return NewGolpe(golpes.ANaoTocouNaBola, golpes.REPLadoOposto, false)
}

func NewQueimou() Golpe {
	return NewGolpe(golpes.AToqueNaBola, golpes.REPLadoOposto, false)
}

func NewToqueNaRede() Golpe {
	return NewGolpe(golpes.AToqueNaRede, golpes.REPLadoOposto, false)
}
