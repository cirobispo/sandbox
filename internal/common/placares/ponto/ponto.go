package ponto

import "github.com/cirobispo/sandbox/internal/common/placares"

type Ponto struct {
	placarA, placarB int
}

func (p Ponto) Terminado() bool {
	return p.placarA != p.placarB
}

func (p Ponto) Lado() placares.LadoDoPlacar {
	if p.placarB > p.placarA {
		return placares.LPOposto
	}

	return placares.LPServico
}

func (p *Ponto) PontuarA() {
	p.placarA, p.placarB = 1, 0
}

func (p *Ponto) PontuarB() {
	p.placarB, p.placarA = 1, 0
}

func (p *Ponto) Zerar() {
	p.placarA, p.placarB = 0, 0
}

func (p Ponto) Tipo() placares.TipoDoPlacar {
	return placares.TPPonto
}

func New() Ponto {
	return Ponto{placarA: 0, placarB: 0}
}
