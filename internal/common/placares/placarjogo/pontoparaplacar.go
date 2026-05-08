package placarjogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/placares/placarponto"
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
)

func PontoParaPlacar(p *point.Point) (placares.EstadoEParametroPlacar, error) {
	if !p.Done() {
		return nil, errors.New("point is still in play.")
	}

	result := placarponto.New()
	result.PontuarA()
	if p.Side() == pointing.PSOpposite {
		result.PontuarB()
	}

	return result, nil
}
