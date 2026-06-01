package placarjogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/placares/placarponto"
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
)

func PontoParaPlacar(p *point.Ponto) (placares.EstadoEParametroPlacar, error) {
	if !p.Terminado() {
		return nil, errors.New("point is still in play.")
	}

	result := placarponto.New()
	result.PontuarA()
	if p.LadoDoPonto() == pointing.LPOposto {
		result.PontuarB()
	}

	return result, nil
}
