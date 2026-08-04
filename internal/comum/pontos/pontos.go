package pontos

import (
	"github.com/cirobispo/sandbox/internal/comum/golpes"
)

type AoAdicionarGolpe func(tipoDoGolpe golpes.TipoDoGolpe, terminado bool)

type Lado int

const (
	LPCorrente Lado = 1
	LPOposto   Lado = 2
	LPNulo     Lado = 0
)

type Pontuando interface {
	Terminado() bool
	LadoDoPonto() Lado
}

func (s Lado) String() string {
	switch s {
	case LPCorrente:
		return "Lado corrente"
	case LPOposto:
		return "Lado oposto"
	default:
		return "Nulo"
	}
}

func (s Lado) Inverso() Lado {
	switch s {
	case LPCorrente:
		return LPOposto
	case LPOposto:
		return LPCorrente
	default:
		return s
	}
}
