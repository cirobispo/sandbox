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

func newPoint(side turnos.LadoDoTurno) Ponto {
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
	points := make([]Ponto, 0, len(everyHit))
	points = append(points, p)

	for i := range everyHit {
		if p.LadoDoPonto() != pointing.LPNulo {
			p = newPoint(turnos.LTB)
			points = append(points, p)
		}

		hit := everyHit[i]
		p.AdicionaGolpe(hit)
		pointSide := p.LadoDoPonto()
		hitSide := hit.Side()
		if hitSide == hitting.HTDSameSide || hitSide == hitting.HTDOppositeSide {
			isSameSide := (pointSide == pointing.LPServico && hitSide == hitting.HTDSameSide)
			isOppositeSide := (pointSide == pointing.LPOposto && hitSide == hitting.HTDOppositeSide)
			if isSameSide || isOppositeSide {
				tt.Logf("On point last ( %d ) side was: %s, point type is %s (%s), point side is %s\n", p.Tamanho(), p.LadoDaBola().LadoCorrente(), hit.Type(), hit.Side(), p.LadoDoPonto())
			}
		} else {
			if hitSide == hitting.HTDConditional {
				tt.Logf("point type is %s (%s), point side is %s\n", hit.Type(), hit.Side(), p.LadoDoPonto())
			}
		}
	}

	showPoints(tt, points)
}

func showPoints(tt *testing.T, points []Ponto) {
	ss, os := 0, 0

	for j := range points {
		p := (points)[j]
		if side := p.LadoDoPonto(); side != pointing.LPNulo {
			ss++
			if side == pointing.LPOposto {
				ss--
				os++
			}

			tt.Logf("Point (%d) last side was: %s, point side is %s => (%d x %d)\n", p.Tamanho(), p.LadoDaBola().LadoCorrente(), p.LadoDoPonto(), ss, os)
		}
	}
}
