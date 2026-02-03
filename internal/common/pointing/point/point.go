package point

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type Point struct {
	ballSide          *turn.Turn
	hits              *[]hitting.Hitting
	onAfterScoreEvent *[]pointing.OnPointScore
}

func New(sideControl *turn.Turn) Point {
	hit := make([]hitting.Hitting, 0, 3)
	events := make([]pointing.OnPointScore, 0)
	return Point{
		ballSide:          sideControl,
		hits:              &hit,
		onAfterScoreEvent: &events,
	}
}

func hasDoubleFault(hits *[]hitting.Hitting) bool {
	lastHit := (*hits)[len(*hits)-1]
	fault := lastHit.Type() == hitting.HTFootFault || lastHit.Type() == hitting.HTServeNet || lastHit.Type() == hitting.HTServeOut
	if lastHit.Side() != hitting.HTDConditional && !fault {
		return false
	}

	count := 0
	result := false
	for i := range *hits {
		hit := (*hits)[i]
		if tp := hit.Type(); tp == hitting.HTServeOut || tp == hitting.HTServeNet || tp == hitting.HTFootFault {
			count++
			if count > 1 {
				result = true
				break
			}
		}
	}
	return result
}

func (p *Point) AddOnBeforeScore(callback pointing.OnPointScore) {
	*p.onAfterScoreEvent = append(*p.onAfterScoreEvent, callback)
}

func (p Point) executeOnScore(hitType hitting.HitType, side hitting.HitSide) {
	for i := range *p.onAfterScoreEvent {
		event := (*p.onAfterScoreEvent)[i]
		event(hitType, side)
	}
}

func (p *Point) AddHit(h hitting.Hitting) hit.Hit {
	//prevent adding new hit after pointing has been met
	if p.PointSide() != pointing.PSNone {
		lastHit := (*p.hits)[len(*p.hits)-1]
		return hit.New(lastHit.Type(), lastHit.Side())
	}

	*p.hits = append(*p.hits, h)

	result := hit.New(h.Type(), h.Side())
	ballInPlay := (result.Side() == hitting.HTDNone || result.Side() == hitting.HTDChangeSide)
	if !ballInPlay {
		if result.Side() == hitting.HTDConditional {
			if !hasDoubleFault(p.hits) {
				return hit.New(h.Type(), h.Side())
			}
			result = hit.NewDoubleFault()
		}

		p.executeOnScore(result.Type(), result.Side())
	} else {
		p.ballSide.Execute()
	}

	return result
}

func (p Point) HitCount() int {
	result := len(*p.hits)
	return result
}

func (p Point) LastHit() (hitting.HitType, error) {
	hitCount := p.HitCount()
	if hitCount == 0 {
		return hitting.HTDoubleFault, fmt.Errorf("no hit found.")
	}

	return (*p.hits)[hitCount-1].Type(), nil
}

func HitSide2PointSide(s hitting.HitSide) pointing.PointSide {
	switch s {
	case hitting.HTDSameSide:
		return pointing.PSStartingSide
	case hitting.HTDOppositeSide:
		return pointing.PSOppositeSide
	default:
		return pointing.PSNone
	}

}

func (p Point) PointSide() pointing.PointSide {
	hitCount := len(*p.hits)
	if hitCount < 1 {
		return pointing.PSNone
	}

	lastHit := (*p.hits)[hitCount-1]
	isDoubleFault := (lastHit.Side() == hitting.HTDConditional && hasDoubleFault(p.hits))
	isOrdinaryPoint := (lastHit.Side() == hitting.HTDSameSide || lastHit.Side() == hitting.HTDOppositeSide)
	if !isOrdinaryPoint && !isDoubleFault {
		return pointing.PSNone
	}

	if isDoubleFault {
		lastHit = hit.NewDoubleFault()
	}

	if p.ballSide.LastSide() == p.ballSide.StartSide() {
		return HitSide2PointSide(lastHit.Side())
	} else {
		return HitSide2PointSide(lastHit.Side()).Inverse()
	}
}

func (p Point) BallStartingSide() turning.SideTurn {
	return p.ballSide.StartSide()
}

func (p Point) BallLastSide() turning.SideTurn {
	return p.ballSide.LastSide()
}
