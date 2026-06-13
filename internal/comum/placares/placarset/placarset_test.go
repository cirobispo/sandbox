package placarset

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type placar struct {
	placarA, placarB int
}

func (p placar) Terminado() bool {
	return true
}

func (p placar) Resultado() (int, int) {
	return p.placarA, p.placarB
}

func (p placar) Lado() placares.LadoDoPlacar {
	if p.placarB > p.placarA {
		return placares.LPOposto
	}

	return placares.LPServico
}

func (p placar) Tipo() placares.TipoDoPlacar {
	return placares.TPJogo
}

func runTest(custom bool, results []placares.EstadoResultadoEParametroPlacar, PlacarA, PlacarB int, t *testing.T) {
	score := New(SetPadrao(turnos.LTA))
	if custom {
		if PlacarB > PlacarA {
			score = New(TamanhoVantagemETieBreak(turnos.LTA, PlacarB, true, true))
		} else {
			score = New(TamanhoVantagemETieBreak(turnos.LTA, PlacarA, true, true))
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

	t.Logf("Testando um resultado para Set com %d jogos. Esperado (%d x %d)\n", len(results), PlacarA, PlacarB)
	for j := range results {
		item := results[j]
		score.AdicionarPlacar(item)
	}

	sA, sB := score.Resultado()
	if sA != PlacarA || sB != PlacarB {
		var descricao = []string{"love", "15", "30", "40", "ad", "game"}
		for j := range results {
			sA, sB = results[j].Resultado()
			scoreA, scoreB := placares.TraduzirPlacar(descricao, sA, sB)
			t.Logf("Score (%v x %v)\n", scoreA, scoreB)
		}
		t.Errorf("\n\nSet deveria ser (%d x %d) e não (%d x %d)\n", PlacarA, PlacarB, sA, sB)
	}
}

func Test_6x4(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 2}, placar{placarA: 3, placarB: 5},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 0, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 4, placarB: 0},
	}

	runTest(false, scores, 6, 4, t)
}

func Test_4x6(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 2}, placar{placarA: 3, placarB: 5},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 0, placarB: 4},

		placar{placarA: 0, placarB: 4}, placar{placarA: 1, placarB: 4},
	}

	runTest(false, scores, 4, 6, t)
}

func Test_6x0(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 4, placarB: 1},

		placar{placarA: 4, placarB: 2}, placar{placarA: 4, placarB: 0},

		placar{placarA: 5, placarB: 3}, placar{placarA: 4, placarB: 0},
	}

	runTest(false, scores, 6, 0, t)
}

func Test_0x6(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarB: 4, placarA: 0}, placar{placarB: 4, placarA: 1},

		placar{placarB: 4, placarA: 2}, placar{placarB: 4, placarA: 0},

		placar{placarB: 5, placarA: 3}, placar{placarB: 4, placarA: 0},
	}

	runTest(false, scores, 0, 6, t)
}

func Test_7x6(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 2}, placar{placarA: 3, placarB: 5},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 0, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, // score{scoreA: 1, scoreB: 4},
	}

	runTest(false, scores, 7, 6, t)
}

func Test_6x7(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 2}, placar{placarA: 3, placarB: 5},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 0, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		// score{scoreA: 4, scoreB: 0},
		placar{placarA: 1, placarB: 4},
	}

	runTest(false, scores, 6, 7, t)
}

func Test_7x5(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 2}, placar{placarA: 3, placarB: 5},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 0, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, // score{scoreA: 1, scoreB: 4},

		placar{placarA: 4, placarB: 0}, // score{scoreA: 1, scoreB: 4},
	}

	runTest(false, scores, 7, 5, t)
}

func Test_4x3(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 2}, placar{placarA: 3, placarB: 5},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, // score{scoreA: 0, scoreB: 4},
	}

	runTest(true, scores, 4, 3, t)
}

func Test_3x4(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 2}, placar{placarA: 3, placarB: 5},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		// score{scoreA: 4, scoreB: 0},
		placar{placarA: 0, placarB: 4},
	}

	runTest(true, scores, 3, 4, t)
}

func Test_TooManyGames(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 2}, placar{placarA: 3, placarB: 5},

		placar{placarA: 4, placarB: 0}, placar{placarA: 1, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 0, placarB: 4},

		placar{placarA: 4, placarB: 0}, placar{placarA: 4, placarB: 0},

		placar{placarA: 4, placarB: 0}, placar{placarA: 4, placarB: 0},

		placar{placarA: 4, placarB: 0}, placar{placarA: 4, placarB: 0},
	}

	runTest(false, scores, 6, 4, t)
}
