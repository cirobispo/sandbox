package ponto

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes/golpe"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turnocontador"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turnotemporizador"
)

func newPoint(side turnos.LadoDoTurno) Ponto {
	ctt := turno.New(turnocontador.ComOutroTurno(turno.New((turnotemporizador.ComOutroTurno(turno.New(turno.DefinindoLado(side)))))))
	return New(ctt)
}

func TestEverySinglePoint(tt *testing.T) {
	everyHit := []golpe.Golpe{golpe.NewAce(), golpe.NewFootFault(), golpe.NewServeOut(),
		golpe.NewHitBackIn(), golpe.NewMiss(), golpe.NewHitNet(), golpe.NewHitOut(), golpe.NewReturnIn(),
		golpe.NewReturnNet(), golpe.NewReturnOut(), golpe.NewServeNet(), golpe.NewServeIn(),
		golpe.NewServeLet(), golpe.NewWinner(), golpe.NewToast(), golpe.NewNetTouch(),
	}

	p := newPoint(turnos.LTB)
	points := make([]Ponto, 0, len(everyHit))
	points = append(points, p)

	for i := range everyHit {
		if p.LadoDoPonto() != pontos.LPNulo {
			p = newPoint(turnos.LTB)
			points = append(points, p)
		}

		hit := everyHit[i]
		p.AdicionarGolpe(hit)
		// pointSide := p.LadoDoPonto()
		// hitSide := hit.Lado()
		// if hitSide == golpes.HTDSameSide || hitSide == golpes.HTDOppositeSide {
		// 	isSameSide := (pointSide == pontos.LPCorrente && hitSide == golpes.HTDSameSide)
		// 	isOppositeSide := (pointSide == pontos.LPOposto && hitSide == golpes.HTDOppositeSide)
		// 	if isSameSide || isOppositeSide {
		// 		tt.Logf("On point last ( %d ) side was: %s, point type is %s (%s), point side is %s\n", p.Tamanho(), p.LadoDaBola(), hit.Tipo(), hit.Lado(), p.LadoDoPonto())
		// 	}
		// } else {
		// 	if hitSide == golpes.HTDConditional {
		// 		tt.Logf("point type is %s (%s), point side is %s\n", hit.Tipo(), hit.Lado(), p.LadoDoPonto())
		// 	}
		// }
	}

	showPoints(tt, points)
}

func showPoints(tt *testing.T, points []Ponto) {
	ss, os := 0, 0

	for j := range points {
		p := (points)[j]
		if side := p.LadoDoPonto(); side != pontos.LPNulo {
			ss++
			if side == pontos.LPOposto {
				ss--
				os++
			}

			tt.Logf("Point (%d) last side was: %s, point side is %s => (%d x %d)\n", p.Tamanho(), p.LadoDaBola(), p.LadoDoPonto(), ss, os)
		}
	}
}
