package pointscore

import "github.com/cirobispo/sandbox/internal/common/scoring"

type PointScore struct {
	scoreA, scoreB int
}

func (p PointScore) Done() bool {
	return p.scoreA != p.scoreB
}

func (p PointScore) Side() scoring.ScoringSide {
	if p.scoreB > p.scoreA {
		return scoring.SSOpposite
	}

	return scoring.SSServing
}

func (p *PointScore) SetA() {
	p.scoreA, p.scoreB = 1, 0
}

func (p *PointScore) SetB() {
	p.scoreB, p.scoreA = 1, 0
}

func (p *PointScore) Unset() {
	p.scoreA, p.scoreB = 0, 0
}

func (p PointScore) Type() scoring.ScoringType {
	return scoring.STPoint
}

func New() PointScore {
	return PointScore{scoreA: 0, scoreB: 0}
}
