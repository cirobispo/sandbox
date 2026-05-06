package ponto

import "github.com/cirobispo/sandbox/internal/common/placares"

type Ponto struct {
	scoreA, scoreB int
}

func (p Ponto) Terminado() bool {
	return p.scoreA != p.scoreB
}

func (p Ponto) Lado() placares.LadoDoPlacar {
	if p.scoreB > p.scoreA {
		return placares.LPOposto
	}

	return placares.LPServico
}

func (p *Ponto) SetA() {
	p.scoreA, p.scoreB = 1, 0
}

func (p *Ponto) SetB() {
	p.scoreB, p.scoreA = 1, 0
}

func (p *Ponto) Unset() {
	p.scoreA, p.scoreB = 0, 0
}

func (p Ponto) Tipo() placares.TipoDoPlacar {
	return placares.TPPonto
}

func New() Ponto {
	return Ponto{scoreA: 0, scoreB: 0}
}
