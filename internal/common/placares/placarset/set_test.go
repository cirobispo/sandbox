package placarset

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/turnos"
)

type score struct {
	scoreA, scoreB int
}

func (s score) Terminado() bool {
	return true
}

func (s score) Resultado() (int, int) {
	return s.scoreA, s.scoreB
}

func (s score) Lado() placares.LadoDoPlacar {
	if s.scoreB > s.scoreA {
		return placares.LPOposto
	}

	return placares.LPServico
}

func (s score) Tipo() placares.TipoDoPlacar {
	return placares.TPJogo
}

func runTest(custom bool, results []placares.EstadoResultadoEParametroPlacar, SideA, SideB int, t *testing.T) {
	score := New(SetPadrao(turnos.LTA))
	if custom {
		if SideB > SideA {
			score = New(TamanhoETieBreak(turnos.LTA, SideB, true, true))
		} else {
			score = New(TamanhoETieBreak(turnos.LTA, SideA, true, true))
		}
	}

	score.AdicionarAoMudarPlacar(func(scoreA, scoreB int, isTieBreak, done bool) {
		if isTieBreak && !done {
			t.Logf("TieBreak (%d x %d)\n", scoreA, scoreB)
		}
	})

	score.AdicionarAoMudarPlacar(func(scoreA, scoreB int, isTieBreak, done bool) {
		if done {
			t.Logf("Score (%d x %d)\n", scoreA, scoreB)
		}
	})

	t.Logf("Testing a result for a Set with %d games. Expected result (%d x %d)\n", len(results), SideA, SideB)
	for j := range results {
		item := results[j]
		score.AdicionarPlacar(item)
	}

	sA, sB := score.Resultado()
	if sA != SideA || sB != SideB {
		var description = []string{"love", "15", "30", "40", "ad", "game"}
		for j := range results {
			sA, sB = results[j].Resultado()
			scoreA, scoreB := placares.TraduzirPlacar(description, sA, sB)
			t.Logf("Score (%v x %v)\n", scoreA, scoreB)
		}
		t.Errorf("\n\nSet should be (%d x %d) not (%d x %d)\n", SideA, SideB, sA, sB)
	}
}

func Test_6x4(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 0, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 4, scoreB: 0},
	}

	runTest(false, scores, 6, 4, t)
}

func Test_4x6(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 0, scoreB: 4},

		score{scoreA: 0, scoreB: 4}, score{scoreA: 1, scoreB: 4},
	}

	runTest(false, scores, 4, 6, t)
}

func Test_6x0(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 4, scoreB: 1},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 4, scoreB: 0},

		score{scoreA: 5, scoreB: 3}, score{scoreA: 4, scoreB: 0},
	}

	runTest(false, scores, 6, 0, t)
}

func Test_0x6(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreB: 4, scoreA: 0}, score{scoreB: 4, scoreA: 1},

		score{scoreB: 4, scoreA: 2}, score{scoreB: 4, scoreA: 0},

		score{scoreB: 5, scoreA: 3}, score{scoreB: 4, scoreA: 0},
	}

	runTest(false, scores, 0, 6, t)
}

func Test_7x6(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 0, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, // score{scoreA: 1, scoreB: 4},
	}

	runTest(false, scores, 7, 6, t)
}

func Test_6x7(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 0, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		// score{scoreA: 4, scoreB: 0},
		score{scoreA: 1, scoreB: 4},
	}

	runTest(false, scores, 6, 7, t)
}

func Test_7x5(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 0, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, // score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, // score{scoreA: 1, scoreB: 4},
	}

	runTest(false, scores, 7, 5, t)
}

func Test_4x3(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, // score{scoreA: 0, scoreB: 4},
	}

	runTest(true, scores, 4, 3, t)
}

func Test_3x4(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		// score{scoreA: 4, scoreB: 0},
		score{scoreA: 0, scoreB: 4},
	}

	runTest(true, scores, 3, 4, t)
}

func Test_TooManyGames(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 0, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 4, scoreB: 0},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 4, scoreB: 0},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 4, scoreB: 0},
	}

	runTest(false, scores, 6, 4, t)
}
