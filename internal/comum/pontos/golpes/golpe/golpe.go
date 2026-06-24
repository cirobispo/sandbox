package golpe

import (
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes"
	"github.com/cirobispo/sandbox/internal/comum/pontos/utilitario"
)

type VerificarAcaoCondicional func(golpes []golpes.Golpe) golpes.TipoAcao

type Golpe struct {
	tipoGolpe   golpes.TipoDoGolpe
	tipoAcao    golpes.TipoAcao
	verificador VerificarAcaoCondicional
}

func (g *Golpe) defineVerificadorCondicional(verificador VerificarAcaoCondicional) {
	g.verificador = verificador
}
func (g Golpe) Tipo() golpes.TipoDoGolpe {
	return g.tipoGolpe
}

func (g Golpe) Acao(gs []golpes.Golpe) golpes.TipoAcao {
	if g.tipoAcao == golpes.TACondicional {
		if g.verificador != nil {
			return g.verificador(gs)
		}
	}
	return g.tipoAcao
}

func verificaDuplaFalta(gps []golpes.Golpe) golpes.TipoAcao {
	if utilitario.ExisteDuplaFalta(&gps) {
		return golpes.TAEncerrarPLO
	}
	return golpes.TACondicional
}

func NewGolpe(t golpes.TipoDoGolpe, acao golpes.TipoAcao, verificador VerificarAcaoCondicional) Golpe {
	return Golpe{
		tipoGolpe: t,
		tipoAcao:  acao,
	}
}

func NewFootFault() Golpe {
	result := NewGolpe(golpes.HTFootFault, golpes.TACondicional, verificaDuplaFalta)
	result.defineVerificadorCondicional(verificaDuplaFalta)
	return result
}

func NewAce() Golpe {
	return NewGolpe(golpes.HTAce, golpes.TAEncerrarPLC, nil)
}

func NewServicoFora() Golpe {
	result := NewGolpe(golpes.HTServeOut, golpes.TACondicional, verificaDuplaFalta)
	result.defineVerificadorCondicional(verificaDuplaFalta)
	return result
}

func NewServicoDentro() Golpe {
	return NewGolpe(golpes.HTServeIn, golpes.TAProsseguir, nil)
}

func NewLET() Golpe {
	return NewGolpe(golpes.HTServeLet, golpes.TANulo, nil)
}

func NewServicoNaRede() Golpe {
	result := NewGolpe(golpes.HTServeNet, golpes.TACondicional, verificaDuplaFalta)
	result.defineVerificadorCondicional(verificaDuplaFalta)
	return result
}

func NewRetornoNaRede() Golpe {
	return NewGolpe(golpes.HTReturnNet, golpes.TAEncerrarPLO, nil)
}

func NewRetornoDentro() Golpe {
	return NewGolpe(golpes.HTReturnIn, golpes.TAProsseguir, nil)
}

func NewRetornoFora() Golpe {
	return NewGolpe(golpes.HTReturnOut, golpes.TAEncerrarPLO, nil)
}

func NewDuplaFalta() Golpe {
	return NewGolpe(golpes.HTDoubleFault, golpes.TAEncerrarPLO, nil)
}

func NewGolpeuNaRede() Golpe {
	return NewGolpe(golpes.HTNet, golpes.TAEncerrarPLO, nil)
}

func NewDevolveuDentro() Golpe {
	return NewGolpe(golpes.HTIn, golpes.TAProsseguir, nil)
}

func NewWinner() Golpe {
	// Aqui será preciso avaliar como atuar aqui.
	// Se o winner é confirmado depois ou no momento do golpe.
	// Provavelmente deverá ser depois. E daí, ficaria, devolveu na rede, fora, winner adversario.
	return NewGolpe(golpes.HTWinner, golpes.TAEncerrarPLC, nil)
}

func NewGolpeuFora() Golpe {
	return NewGolpe(golpes.HTOut, golpes.TAEncerrarPLO, nil)
}

func NewNaoTocou() Golpe {
	return NewGolpe(golpes.HTMiss, golpes.TAEncerrarPLO, nil)
}

func NewQueimou() Golpe {
	return NewGolpe(golpes.HTToast, golpes.TAEncerrarPLO, nil)
}

func NewToqueNaRede() Golpe {
	return NewGolpe(golpes.HTNetTouch, golpes.TAEncerrarPLO, nil)
}
