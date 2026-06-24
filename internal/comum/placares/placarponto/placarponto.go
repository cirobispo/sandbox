package placarponto

import (
	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
)

type Ponto struct {
	ponto            ponto.Ponto
	placarA, placarB int
}

func New(ponto *ponto.Ponto) Ponto {
	result := Ponto{ponto: ponto.Clonar(), placarA: 0, placarB: 0}

	result.placarA, result.placarB = 1, 0
	if ponto.LadoDoPonto() == pontos.LPOposto {
		result.placarA, result.placarB = 0, 1
	}

	return result
}

func (p Ponto) Lado() placares.LadoDoPlacar {
	return placares.LPNulo
}

func (p Ponto) Tipo() placares.TipoDoPlacar {
	return placares.TPPonto
}

func (p Ponto) Terminado() bool {
	return p.ponto.Terminado()
}

func (p Ponto) Resultado() (int, int) {
	return p.placarA, p.placarB
}
