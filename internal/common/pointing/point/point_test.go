package point_test

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/countingturn"
	"github.com/cirobispo/sandbox/internal/common/turning/timingturn"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

func newPoint(side turning.SideTurn) point.Point {
	ctt := timingturn.NewFromTurn(countingturn.NewFromTurn(turn.New(side)))

	return point.New(ctt)
}

func TestEverySinglePoint(tt *testing.T) {
	everyHit := []hit.Hit{hit.NewAce(), hit.NewFootFault(), hit.NewServeOut(),
		hit.NewIn(), hit.NewMiss(), hit.NewNet(), hit.NewOut(), hit.NewReturnIn(),
		hit.NewReturnNet(), hit.NewReturnOut(), hit.NewServeNet(), hit.NewServeIn(),
		hit.NewServeLet(), hit.NewWinner(), hit.NewToast(), hit.NewNetTouch(),
	}

	p := newPoint(turning.STB)
	points := make([]point.Point, 0, len(everyHit))
	points = append(points, p)

	for i := range everyHit {
		if p.PointSide() != pointing.PSNone {
			p = newPoint(turning.STB)
			points = append(points, p)
		}

		hit := everyHit[i]
		p.AddHit(hit)
		pointSide := p.PointSide()
		hitSide := hit.Side()
		if hitSide == hitting.HTDSameSide || hitSide == hitting.HTDOppositeSide {
			isSameSide := (pointSide == pointing.PSStartingSide && hitSide == hitting.HTDSameSide)
			isOppositeSide := (pointSide == pointing.PSOppositeSide && hitSide == hitting.HTDOppositeSide)
			if isSameSide || isOppositeSide {
				tt.Logf("On point last ( %d ) side was: %s, point type is %s (%s), point side is %s\n", p.HitCount(), p.BallLastSide(), hit.Type(), hit.Side(), p.PointSide())
			}
		} else {
			if hitSide == hitting.HTDConditional {
				tt.Logf("point type is %s (%s), point side is %s\n", hit.Type(), hit.Side(), p.PointSide())
			}
		}
	}

	showPoints(tt, points)
}

func showPoints(tt *testing.T, points []point.Point) {
	ss, os := 0, 0

	for j := range points {
		p := (points)[j]
		if side := p.PointSide(); side != pointing.PSNone {
			ss++
			if side == pointing.PSOppositeSide {
				ss--
				os++
			}

			tt.Logf("Point (%d) last side was: %s, point side is %s => (%d x %d)\n", p.HitCount(), p.BallLastSide(), p.PointSide(), ss, os)
		}
	}
}
