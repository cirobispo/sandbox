package point

import (
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type Point struct {
	side turn.Turn
	hits []hitting.Hitting
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
	return Point{side: t,
		hits: make([]hitting.Hitting, 0, 3),
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

	if result != hitting.HTDChangeSide {
		p.side.Execute()
	}

	return result
}
