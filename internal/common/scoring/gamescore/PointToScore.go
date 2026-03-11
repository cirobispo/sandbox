package gamescore

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/scoring/pointscore"
)

func PointToScore(p *point.Point, scoreA, scoreB int, decidingPoint bool) (scoring.Scoring, error) {
	if !p.Done() {
		return nil, errors.New("point is still in play.")
	}

	result := pointscore.New()
	result.SetA()
	if p.Side() == pointing.PSOpposite {
		result.SetB()
	}

	return result, nil
}
