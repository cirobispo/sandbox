package pointscore

import "github.com/cirobispo/sandbox/internal/common/scoring"

type PointScore struct {
	scoreA, scoreB int
}

func (p PointScore) Done() bool {
	return p.scoreA != p.scoreB
}

func (p PointScore) Result() (int, int) {
	return p.scoreA, p.scoreB
}

func (p PointScore) Side() scoring.ScoringSide {
	if p.scoreB > p.scoreA {
		return scoring.SSB
	}

	return scoring.SSA
}

func (p *PointScore) SetA() {
	p.scoreA, p.scoreB = 1, 0
}

func (p *PointScore) SetB() {
	p.scoreB, p.scoreA = 1, 0
}

func (p *PointScore) UnsetA() {
	p.scoreA = 0
}

func (p *PointScore) UnsetB() {
	p.scoreB = 0
}

func (p PointScore) Type() scoring.ScoringType {
	return scoring.STPoint
}

func (p PointScore) InverseScore(callback func()) func() {
	// if callback == p.UnsetA {
	// 	return p.UnsetB
	// }
	return p.SetA
}

func New() PointScore {
	return PointScore{scoreA: 0, scoreB: 0}
}
