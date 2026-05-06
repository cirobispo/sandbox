package jogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/placares"
	pointscore "github.com/cirobispo/sandbox/internal/common/placares/ponto"
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
)

func PontoParaPlacar(p *point.Point) (placares.EstadoEParametroPlacar, error) {
	if !p.Done() {
		return nil, errors.New("point is still in play.")
	}

	result := pointscore.New()
	result.PontuarA()
	if p.Side() == pointing.PSOpposite {
		result.PontuarB()
	}

	return result, nil
}
