package point

import (
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type Point struct {
	side         turn.Turn
	hits         []hitting.Hitting
	OnAfterScore []pointing.OnPointScore
}

func checkDoubleFault(hits *[]hitting.Hitting) hitting.HitSide {
	count := 0
	result := hitting.HTDNone
	for i := range *hits {
		hit := (*hits)[i]
		if tp := hit.Type(); tp == hitting.HTServeOut || tp == hitting.HTServeNet || tp == hitting.HTFootFault {
			count++
			if count > 1 {
				result = hitting.HTDOpositeSide
				break
			}
		}
	}
	return result
}

func checkPointSide(startSide, currentSide turning.SideTurn) hitting.HitSide {
	result := hitting.HTDSameSide
	if startSide != currentSide {
		result = hitting.HTDOpositeSide
	}
	return result
}

func New(t turn.Turn) Point {
	return Point{
		side:         t,
		hits:         make([]hitting.Hitting, 0, 3),
		OnAfterScore: make([]pointing.OnPointScore, 0),
	}
}

func oppositeSide(side pointing.PointSide) pointing.PointSide {
	if side == pointing.PSA {
		return pointing.PSB
	}

	return pointing.PSA
}

func (p *Point) getCorrectSide() pointing.PointSide {
	if len(p.hits) < 1 {
		return pointing.PSNone
	}

	hitCount := len(p.hits)
	isSameSide := (hitCount - ((hitCount / 2) * 2)) == 0 // can I trust on int division on this case?
	hitSide := p.hits[hitCount-1].Side()

	var result pointing.PointSide
	result = pointing.PointSide(p.side.StartSide())
	if hitSide == hitting.HTDOpositeSide {
		result = oppositeSide(result)
	}

	if !isSameSide {
		result = oppositeSide(result)
	}

	return result
}

func (p *Point) AddBeforeScoreEvent(pse pointing.OnPointScore) {
	p.OnAfterScore = append(p.OnAfterScore, pse)
}

func (p Point) executeOnScore(hitType hitting.HitType, side hitting.HitSide) {
	for i := range p.OnAfterScore {
		event := p.OnAfterScore[i]
		event(hitType, side)
	}
}

func (p *Point) AddHit(h hitting.Hitting) hitting.HitSide {
	p.hits = append(p.hits, h)

	result := h.Side()
	if result == hitting.HTDConditional {
		if t := h.Type(); t == hitting.HTFootFault || t == hitting.HTServeNet || t == hitting.HTServeOut {
			result = checkDoubleFault(&p.hits)
		}
	}

	if result == hitting.HTDOpositeSide || result == hitting.HTDSameSide {
		result = checkPointSide(p.side.StartSide(), p.side.CurrentSide())
	}

	if result != hitting.HTDNone && result != hitting.HTDChangeSide {
		p.side.Execute()
		p.executeOnScore(h.Type(), h.Side())
	}

	return result
}

func (p Point) Side() pointing.PointSide {
	result := p.getCorrectSide()

	return result
}
