package placarponto

import (
	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
)

type Ponto struct {
	ponto ponto.Ponto
}

func New(ponto *ponto.Ponto) Ponto {
	return Ponto{ponto: ponto.Clonar()}
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
	if p.ponto.Terminado() {
		if p.ponto.LadoDoPonto() == pontos.LPCorrente {
			return 1, 0
		}
		return 0, 1
	}

	return 0, 0
}
