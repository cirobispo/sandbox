package pointing

import "github.com/cirobispo/sandbox/internal/common/pointing/hitting"

type AoPontuarNoPlacar func(golpe hitting.HitType, lado hitting.HitSide, done bool)

type LadoDoPonto int

const (
	LPServico LadoDoPonto = 1
	LPOposto  LadoDoPonto = 2
	LPNulo    LadoDoPonto = 0
)

func (s LadoDoPonto) String() string {
	switch s {
	case LPServico:
		return "Serving side"
	case LPOposto:
		return "Opposite side"
	default:
		return "None"
	}
}

func (s LadoDoPonto) Inverso() LadoDoPonto {
	switch s {
	case LPServico:
		return LPOposto
	case LPOposto:
		return LPServico
	default:
		return s
	}
}
