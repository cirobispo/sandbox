package placarjogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/placares/placarponto"
	"github.com/cirobispo/sandbox/internal/common/pontos"
	"github.com/cirobispo/sandbox/internal/common/pontos/ponto"
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
