package point

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
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

func (p *Point) AddHit(h hitting.Hitting) hitting.HitSide {
	//prevent adding new hit after pointing has been met
	if p.PointSide() != pointing.PSNone {
		return (*p.hits)[len(*p.hits)-1].Side()
	}

	*p.hits = append(*p.hits, h)

	result := h.Side()
	ballInPlay := (result == hitting.HTDNone || result == hitting.HTDChangeSide)
	if !ballInPlay {
		if result == hitting.HTDConditional {
			if !hasDoubleFault(p.hits) {
				return hitting.HTDNone
			}
		}

		p.ballSide.Execute()
		p.executeOnScore(h.Type(), h.Side())
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
		return hitting.HTMiss, fmt.Errorf("no hit found.")
	}

	return (*p.hits)[hitCount-1].Type(), nil
}

func (p Point) PointSide() pointing.PointSide {
	hitCount := len(*p.hits)
	if (hitCount < 1) || (hitCount > 0) && ((*p.hits)[hitCount-1].Side() == hitting.HTDNone || (*p.hits)[hitCount-1].Side() == hitting.HTDChangeSide) {
		return pointing.PSNone
	}

	lastHit := (*p.hits)[hitCount-1]
	if lastHit.Side() == hitting.HTDConditional {
		if hasDoubleFault(p.hits) {
			return pointing.PSOppositeSide
		}
		return pointing.PSNone
	}

	if lastHit.Side() == hitting.HTDOppositeSide {
		return pointing.PSOppositeSide
	}

	return pointing.PSStartingSide
}

func (p Point) BallStartingSide() turning.SideTurn {
	return p.ballSide.StartSide()
}

func (p Point) BallLastSide() turning.SideTurn {
	return p.ballSide.LastSide()
}
