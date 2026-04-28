package pointscore

import "github.com/cirobispo/sandbox/internal/common/scoring"

type PointScore struct {
	scoreA, scoreB int
}

func (p PointScore) Terminado() bool {
	return p.scoreA != p.scoreB
}

func (p PointScore) Lado() scoring.LadoDoPlacar {
	if p.scoreB > p.scoreA {
		return scoring.LPOposto
	}

	return scoring.LPServico
}

func (p *PointScore) SetA() {
	p.scoreA, p.scoreB = 1, 0
}

func (p *PointScore) SetB() {
	p.scoreB, p.scoreA = 1, 0
}

func (p *PointScore) Unset() {
	p.scoreA, p.scoreB = 0, 0
}

func (p PointScore) Tipo() scoring.TipoDoPlacar {
	return scoring.TPPonto
}

func New() PointScore {
	return PointScore{scoreA: 0, scoreB: 0}
}
