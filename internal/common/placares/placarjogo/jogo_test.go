package placarjogo

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/turnos"
	"github.com/cirobispo/sandbox/internal/common/turnos/turno"
)

func runTest(blocks []point.TestBlock, SideA, SideB int, t *testing.T) {
	score := New(turnos.LTA, false)
	breakPoint := 0

	score.AdicionaAoMudarPlacar(func(scoreA, scoreB int, done bool) {
		var description = []string{"love", "15", "30", "40", "ad", "game"}
		a, b := placares.TraduzirPlacar(description, scoreA, scoreB)
		if done {
			t.Logf("Game FINAL status: ( %s x %s )\n", a, b)
			return
		}

		t.Logf("Game status: ( %s x %s )\n", a, b)
	})

	score.AdicionaAoMudarPlacar(func(scoreA, scoreB int, done bool) {
		if !done && (scoreB >= 3 && scoreA < scoreB) {
			breakPoint++
			t.Logf("Break point: ( #%v )\n", breakPoint)
		}
	})

	t.Log("Game status: ( 0 x 0 )")
	for i := range blocks {
		block := blocks[i]
		point := point.New(turno.New(turno.MudandoLado(turnos.LTA)))
		for j := range block.Items {
			hit := block.Items[j].Value
			point.AddHit(hit)
		}

		scoreToAdd, error := PontoParaPlacar(&point)
		if error != nil {
			t.Errorf("\n\n%s", error.Error())
		}
		score.AdicionaPlacar(scoreToAdd)
	}

	a, b := score.Resultado()
	if a != SideA || b != SideB {
		t.Errorf("\n\nScore should be (%d x %d) not (%d x %d)\n", SideA, SideB, a, b)
	}
}

func TestToSideA(t *testing.T) {
	blocks := []point.TestBlock{point.AcePoint(), point.AcePoint(), point.WinnerSSPoint(), point.WinnerOSPoint(),
		point.WinnerOSPoint(), point.WinnerOSPoint(), point.DoubleFault(), point.AcePoint(), point.AcePoint(), point.AcePoint(),
		point.LongRallieOSPoint(3, point.NetOppositeSide(true)),
	}

	runTest(blocks, 5, 3, t)
}

func TestToSideB(t *testing.T) {
	blocks := []point.TestBlock{point.AcePoint(), point.AcePoint(), point.WinnerSSPoint(), point.WinnerOSPoint(),
		point.WinnerOSPoint(), point.WinnerOSPoint(), point.DoubleFault(), point.AcePoint(), point.DoubleFault(), point.DoubleFault(),
		point.LongRallieOSPoint(3, point.NetOppositeSide(true)),
	}

	runTest(blocks, 3, 5, t)
}
