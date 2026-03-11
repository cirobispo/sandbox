package gamescore

import (
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/scoring/pointscore"
)

func PointToScore(p *point.Point, scoreA, scoreB int, decidingPoint bool) scoring.Scoring {
	result := pointscore.New()

	if who := p.Side(); who != pointing.PSNone {
		// incr := 1
		// sideToAdd := scoreA
		// if who == pointing.PSOppositeSide {
		// 	sideToAdd = scoreB
		// }

		// sA, sB := result.Result()
		// if (!decidingPoint) && (sA > 3 || sB > 3) {
		// 	if sideToAdd == 3 {
		// 		incr = -1
		// 		sideToAdd = result.InverseScore(sideToAdd)
		// 	}
		// }
		// *sideToAdd += incr
	}

	return result
}
