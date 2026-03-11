package gamescore

import (
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
)

type score struct {
	scoreA, scoreB int
}

func (s score) Done() bool {
	return s.scoreA != s.scoreB
}

func (s score) Result() (int, int) {
	return s.scoreA, s.scoreB
}

func (s score) Side() scoring.ScoringSide {
	if s.scoreB > s.scoreA {
		return scoring.SSB
	}

	return scoring.SSA
}

func (s score) Type() scoring.ScoringType {
	return scoring.STPoint
}

func (s score) inverseScore(side *int) *int {
	sA, sB := &s.scoreA, &s.scoreB
	if side == sA {
		return sB
	}
	return sA
}

func PointToScore(p *point.Point, decidingPoint bool) scoring.Scoring {
	result := score{scoreA: 0, scoreB: 0}

	if who := p.Side(); who != pointing.PSNone {
		incr := 1
		sideToAdd := &result.scoreA
		if who == pointing.PSOppositeSide {
			sideToAdd = &result.scoreB
		}

		if (!decidingPoint) && (result.scoreA > 3 || result.scoreB > 3) {
			if *sideToAdd == 3 {
				incr = -1
				sideToAdd = result.inverseScore(sideToAdd)
			}
		}
		*sideToAdd += incr
	}

	return result
}
