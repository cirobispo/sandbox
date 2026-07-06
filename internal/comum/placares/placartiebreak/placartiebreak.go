package placartiebreak

import (
	"github.com/cirobispo/sandbox/internal/comum/placares/placarjogo"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type AoMudarPlacar func(placarA, placarB int, tieBreak, terminado bool)
type AoMudarSacador func(lado turnos.LadoDoTurno)

type ParamOption func() TieBreak

type TieBreak struct {
	*placarjogo.Jogo
}

func terminado7(valores ...int) bool {
	sA, sB, pontoDecisivo := valores[0], valores[1], valores[2] == 1
	diff := 2
	if pontoDecisivo {
		diff--
	}

	AWins := sA >= 7 && (sA-sB) >= diff
	BWins := sB >= 7 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}

func terminado10(valores ...int) bool {
	sA, sB, pontoDecisivo := valores[0], valores[1], valores[2] == 1
	diff := 2
	if pontoDecisivo {
		diff--
	}

	AWins := sA >= 10 && (sA-sB) >= diff
	BWins := sB >= 10 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}

func ChegarEm7(ladoInicial turnos.LadoDoTurno, pontoDecisivo bool) ParamOption {
	return func() TieBreak {
		result := TieBreak{Jogo: placarjogo.New(ladoInicial, pontoDecisivo)}
		result.Jogo.DefinirTestaEncerramento(terminado7)
		return result
	}
}

func ChegarEm10(ladoInicial turnos.LadoDoTurno, pontoDecisivo bool) ParamOption {
	return func() TieBreak {
		result := TieBreak{Jogo: placarjogo.New(ladoInicial, pontoDecisivo)}
		result.Jogo.DefinirTestaEncerramento(terminado10)
		return result
	}
}

func New(param ParamOption) TieBreak {
	return param()
}
