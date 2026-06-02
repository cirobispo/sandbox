package placarjogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/placares/placarponto"
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
)

func PontoParaPlacar(p *ponto.Ponto) (placares.EstadoEParametroPlacar, error) {
	if !p.Terminado() {
		return nil, errors.New("point is still in play.")
	}

	result := placarponto.New()
	result.PontuarA()
	if p.LadoDoPonto() == pontos.LPOposto {
		result.PontuarB()
	}

	return result, nil
}
