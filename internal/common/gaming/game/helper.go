package game

import (
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
)

type score struct {
	done           bool
	scoreA, scoreB int
}

func (s score) Done() bool {
	return true
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

func PointToScore(p *point.Point) scoring.Scoring {
	result := score{done: p.Done(), scoreA: 0, scoreB: 0}
	if p.Done() {
		toAdd := &result.scoreA
		if p.Side() == pointing.PSOppositeSide {
			toAdd = &result.scoreB
		}
		(*toAdd)++
	}

	return result
}
