package placarponto

import "github.com/cirobispo/sandbox/internal/comum/placares"

type PlacarPonto struct {
	placarA, placarB int
}

func (p PlacarPonto) Lado() placares.LadoDoPlacar {
	if p.placarB > p.placarA {
		return placares.LPOposto
	}

	return placares.LPServico
}

func (p PlacarPonto) Tipo() placares.TipoDoPlacar {
	return placares.TPPonto
}

func (p PlacarPonto) Terminado() bool {
	return p.placarA != p.placarB
}

func (p PlacarPonto) Resultado() (int, int) {
	return 0, 0
}

func (p *PlacarPonto) PontuarA() {
	p.placarA, p.placarB = 1, 0
}

func (p *PlacarPonto) PontuarB() {
	p.placarB, p.placarA = 1, 0
}

func (p *PlacarPonto) Zerar() {
	p.placarA, p.placarB = 0, 0
}

func New() PlacarPonto {
	return PlacarPonto{placarA: 0, placarB: 0}
}
