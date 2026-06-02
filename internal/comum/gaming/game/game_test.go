package game

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

func runTest(personToServe turnos.LadoDoTurno, blocks []ponto.TestBlock, SideA, SideB int, t *testing.T) {
	g := New(turno.New(turno.MudandoLado(personToServe)), false)
	g.AddOnAddingPointEvent(func(scoreA, scoreB int, done bool) {
		var description = []string{"love", "15", "30", "40", "ad", "game"}
		tA, tB := placares.TraduzirPlacar(description, scoreA, scoreB)
		if done {
			t.Logf("Game FINAL status: ( %v x %v )\n", tA, tB)
			t.Log()
			return
		}

		t.Logf("Game status: ( %v x %v )\n", tA, tB)
		t.Log()
	})

	for i := range blocks {
		block := blocks[i]
		tn := turno.New(turno.MudandoLado(turnos.LTA))
		p := ponto.New(tn)

		for j := range block.Items {
			item := block.Items[j]
			t.Logf("%s hits %s, ", tn.LadoCorrente().String(), item.Value.Tipo())
			p.AdicionaGolpe(item.Value)
		}

		g.AddPoint(p)
	}

	a, b := g.Score().Resultado()

	if a != SideA || b != SideB {
		t.Errorf("\n\nGame should be (%d x %d) not (%d x %d)\n", SideA, SideB, a, b)
	}
}

func TestTurnA_Game40(t *testing.T) {
	blocks := []ponto.TestBlock{ponto.AcePoint(), ponto.AcePoint(), ponto.WinnerSSPoint(),
		ponto.WinnerOSPoint(), ponto.WinnerOSPoint(), ponto.WinnerOSPoint(), ponto.WinnerOSPoint(),
		// point.DoubleFault(),
		ponto.LongRallieOSPoint(2, ponto.NetOppositeSide(true)),
		ponto.LongRallieOSPoint(2, ponto.NetOppositeSide(true)),
		// point.LongRallieOSPoint(2, point.NetSameSide(true)),
		ponto.AcePoint(),
	}

	runTest(turnos.LTA, blocks, 5, 3, t)
}

func TestTurnB_40Game(t *testing.T) {
	blocks := []ponto.TestBlock{ponto.AcePoint(), ponto.AcePoint(), ponto.WinnerSSPoint(),
		ponto.WinnerOSPoint(), ponto.WinnerOSPoint(), ponto.WinnerOSPoint(), ponto.WinnerOSPoint(),
		// point.DoubleFault(),
		ponto.LongRallieOSPoint(2, ponto.NetOppositeSide(true)),
		ponto.LongRallieOSPoint(2, ponto.NetOppositeSide(true)),
		// point.LongRallieOSPoint(2, point.NetSameSide(true)),
		ponto.AcePoint(),
	}

	runTest(turnos.LTB, blocks, 3, 5, t)
}
