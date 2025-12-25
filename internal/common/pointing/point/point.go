package point

import (
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type Point struct {
	side turn.Turn
	hits []hitting.Hitting
}

func New(t turn.Turn) Point {
	return Point{side: t,
		hits: make([]hitting.Hitting, 0, 3),
	}
}

func (p *Point) AddHit(h hitting.Hitting) hitting.HitPointDestination {
	p.hits = append(p.hits, h)

	result := h.PointDestination(&p.hits)
	if result != hitting.HTDNone {
		p.side.Execute()
	}

	return result
}
