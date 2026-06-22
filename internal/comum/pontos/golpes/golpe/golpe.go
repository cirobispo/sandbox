package golpe

import (
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes"
	"github.com/cirobispo/sandbox/internal/comum/pontos/utilitario"
)

type VerificarGolpes func(golpes []golpes.Golpe) golpes.TipoAcao

type Golpe struct {
	tipoGolpe golpes.TipoDoGolpe
	tipoAcao  golpes.TipoAcao
	verificar VerificarGolpes
}

func (g Golpe) Tipo() golpes.TipoDoGolpe {
	return g.tipoGolpe
}

func (g Golpe) Acao(gs []golpes.Golpe) golpes.TipoAcao {
	if g.tipoAcao == golpes.TACondicional {
		if g.verificar != nil {
			return g.verificar(gs)
		}
	}
	return g.tipoAcao
}

func New(t golpes.TipoDoGolpe, acao golpes.TipoAcao, verificador VerificarGolpes) Golpe {
	return Golpe{
		tipoGolpe: t,
		tipoAcao:  acao,
		verificar: verificador,
	}
}

func verificaDuplaFalta(gps []golpes.Golpe) golpes.TipoAcao {
	if utilitario.ExisteDuplaFalta(&gps) {
		return golpes.TAEncerrarPLO
	}
	return golpes.TACondicional
}

func NewFootFault() Golpe {
	return New(golpes.HTFootFault, golpes.TACondicional, verificaDuplaFalta)
}

func NewAce() Golpe {
	return New(golpes.HTAce, golpes.TAEncerrarPLC, nil)
}

func NewServeOut() Golpe {
	return New(golpes.HTServeOut, golpes.TACondicional, verificaDuplaFalta)
}

func NewServeIn() Golpe {
	return New(golpes.HTServeIn, golpes.TAProsseguir, nil)
}

func NewServeLet() Golpe {
	return New(golpes.HTServeLet, golpes.TACondicional, nil)
}

func NewServeNet() Golpe {
	return New(golpes.HTServeNet, golpes.TACondicional, verificaDuplaFalta)
}

func NewReturnNet() Golpe {
	return New(golpes.HTReturnNet, golpes.TAEncerrarPLO, nil)
}

func NewReturnIn() Golpe {
	return New(golpes.HTReturnIn, golpes.TAProsseguir, nil)
}

func NewReturnOut() Golpe {
	return New(golpes.HTReturnOut, golpes.TAEncerrarPLO, nil)
}

func NewDoubleFault() Golpe {
	return New(golpes.HTDoubleFault, golpes.TAEncerrarPLO, nil)
}

func NewHitNet() Golpe {
	return New(golpes.HTNet, golpes.TAEncerrarPLO, nil)
}

func NewHitBackIn() Golpe {
	return New(golpes.HTIn, golpes.TAProsseguir, nil)
}

func NewWinner() Golpe {
	// Aqui será preciso avaliar como atuar aqui.
	// Se o winner é confirmado depois ou no momento do golpe.
	// Provavelmente deverá ser depois. E daí, ficaria, devolveu na rede, fora, winner adversario.
	return New(golpes.HTWinner, golpes.TAEncerrarPLC, nil)
}

func NewHitOut() Golpe {
	return New(golpes.HTOut, golpes.TAEncerrarPLO, nil)
}

func NewMiss() Golpe {
	return New(golpes.HTOut, golpes.TAEncerrarPLO, nil)
}

func NewToast() Golpe {
	return New(golpes.HTToast, golpes.TAEncerrarPLO, nil)
}

func NewNetTouch() Golpe {
	return New(golpes.HTNetTouch, golpes.TAEncerrarPLO, nil)
}
