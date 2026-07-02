package pontos

import (
	"github.com/cirobispo/sandbox/internal/comum/golpes"
)

type AoAdicionarGolpe func(tipoDoGolpe golpes.TipoDoGolpe, terminado bool)

type LadoDoPonto int

const (
	LPCorrente LadoDoPonto = 1
	LPOposto   LadoDoPonto = 2
	LPNulo     LadoDoPonto = 0
)

type Pontuando interface {
	Terminado() bool
	LadoDoPonto() LadoDoPonto
}

func (s LadoDoPonto) String() string {
	switch s {
	case LPCorrente:
		return "Lado corrente"
	case LPOposto:
		return "Lado oposto"
	default:
		return "Nulo"
	}
}

func (s LadoDoPonto) Inverso() LadoDoPonto {
	switch s {
	case LPCorrente:
		return LPOposto
	case LPOposto:
		return LPCorrente
	default:
		return s
	}
}
