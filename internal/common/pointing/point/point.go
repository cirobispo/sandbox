package point

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
	"github.com/cirobispo/sandbox/internal/common/turnos/turno"
)

type Point struct {
	ballSide          *turno.Turno
	hits              *[]hitting.Hitting
	done              bool
	onAfterScoreEvent *[]pointing.OnScoringPoint
}

func New(sideControl *turno.Turno) Point {
	hit := make([]hitting.Hitting, 0, 3)
	events := make([]pointing.OnScoringPoint, 0)
	return Point{
		done:              false,
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

func (p *Point) AddOnAfterScore(callback pointing.OnScoringPoint) {
	*p.onAfterScoreEvent = append(*p.onAfterScoreEvent, callback)
}

func (p *Point) AddHit(h hitting.Hitting) {
	if p.done {
		return
	}

	*p.hits = append(*p.hits, h)

	ballInPlay := (h.Side() == hitting.HTDNone || h.Side() == hitting.HTDChangeSide)
	if ballInPlay {
		p.ballSide.Execute()
		return
	}

	if h.Side() == hitting.HTDConditional && !hasDoubleFault(p.hits) {
		return
	}

	p.done = true
	p.executeOnScore(h.Type(), h.Side(), p.done)
}

func (p Point) Hits() []hit.Hit {
	result := make([]hit.Hit, 0, len(*p.hits))
	for j := range *p.hits {
		item := (*p.hits)[j]
		result = append(result, hit.New(item.Type(), item.Side()))
	}

	return result
}

func (p Point) Length() int {
	result := len(*p.hits)
	return result
}

func (p Point) LastHit() (hitting.HitType, error) {
	hitCount := p.Length()
	if hitCount == 0 {
		return hitting.HTDoubleFault, errors.New("no hit found.")
	}

	return (*p.hits)[hitCount-1].Type(), nil
}

func (p Point) Ball() turno.Turno {
	return *p.ballSide
}

func (p Point) Side() pointing.PointSide {
	if !p.Done() {
		return pointing.PSNone
	}

	lastHit := (*p.hits)[len(*p.hits)-1]
	isDoubleFault := (lastHit.Side() == hitting.HTDConditional && hasDoubleFault(p.hits))
	isOrdinaryPoint := (lastHit.Side() == hitting.HTDSameSide || lastHit.Side() == hitting.HTDOppositeSide)
	if !isOrdinaryPoint && !isDoubleFault {
		return pointing.PSNone
	}

	if isDoubleFault {
		lastHit = hit.NewDoubleFault()
	}

	if p.ballSide.LadoCorrente() == p.ballSide.LadoInicial() {
		return HitSide2PointSide(lastHit.Side())
	} else {
		return HitSide2PointSide(lastHit.Side()).Inverse()
	}
}

func (p Point) Done() bool {
	return p.done
}

func (p Point) Clone() Point {
	result := New(p.ballSide.Clonar(p.ballSide.LadoInicial()))
	copy(*result.hits, *p.hits)
	copy(*result.onAfterScoreEvent, *p.onAfterScoreEvent)
	result.done = p.done

	return result
}

func (p Point) executeOnScore(hitType hitting.HitType, side hitting.HitSide, done bool) {
	for i := range *p.onAfterScoreEvent {
		event := (*p.onAfterScoreEvent)[i]
		event(hitType, side, done)
	}
}

func HitSide2PointSide(s hitting.HitSide) pointing.PointSide {
	switch s {
	case hitting.HTDSameSide:
		return pointing.PSServing
	case hitting.HTDOppositeSide:
		return pointing.PSOpposite
	default:
		return pointing.PSNone
	}
}
