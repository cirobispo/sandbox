package point

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
	"github.com/cirobispo/sandbox/internal/common/turnos"
	"github.com/cirobispo/sandbox/internal/common/turnos/turno"
	"github.com/cirobispo/sandbox/internal/common/turnos/turnocontador"
	"github.com/cirobispo/sandbox/internal/common/turnos/turnotemporizador"
)

func newPoint(side turnos.LadoDoTurno) Point {
	ctt := turno.New(turnocontador.ComOutroTurno(turno.New((turnotemporizador.ComOutroTurno(turno.New(turno.MudandoLado(side)))))))
	return New(ctt)
}

func TestEverySinglePoint(tt *testing.T) {
	everyHit := []hit.Hit{hit.NewAce(), hit.NewFootFault(), hit.NewServeOut(),
		hit.NewHitBackIn(), hit.NewMiss(), hit.NewHitNet(), hit.NewHitOut(), hit.NewReturnIn(),
		hit.NewReturnNet(), hit.NewReturnOut(), hit.NewServeNet(), hit.NewServeIn(),
		hit.NewServeLet(), hit.NewWinner(), hit.NewToast(), hit.NewNetTouch(),
	}

	p := newPoint(turnos.LTB)
	points := make([]Point, 0, len(everyHit))
	points = append(points, p)

	for i := range everyHit {
		if p.Side() != pointing.PSNone {
			p = newPoint(turnos.LTB)
			points = append(points, p)
		}

		hit := everyHit[i]
		p.AddHit(hit)
		pointSide := p.Side()
		hitSide := hit.Side()
		if hitSide == hitting.HTDSameSide || hitSide == hitting.HTDOppositeSide {
			isSameSide := (pointSide == pointing.PSServing && hitSide == hitting.HTDSameSide)
			isOppositeSide := (pointSide == pointing.PSOpposite && hitSide == hitting.HTDOppositeSide)
			if isSameSide || isOppositeSide {
				tt.Logf("On point last ( %d ) side was: %s, point type is %s (%s), point side is %s\n", p.Length(), p.Ball().LadoCorrente(), hit.Type(), hit.Side(), p.Side())
			}
		} else {
			if hitSide == hitting.HTDConditional {
				tt.Logf("point type is %s (%s), point side is %s\n", hit.Type(), hit.Side(), p.Side())
			}
		}
	}

	showPoints(tt, points)
}

func showPoints(tt *testing.T, points []Point) {
	ss, os := 0, 0

	for j := range points {
		p := (points)[j]
		if side := p.Side(); side != pointing.PSNone {
			ss++
			if side == pointing.PSOpposite {
				ss--
				os++
			}

			tt.Logf("Point (%d) last side was: %s, point side is %s => (%d x %d)\n", p.Length(), p.Ball().LadoCorrente(), p.Side(), ss, os)
		}
	}
}
