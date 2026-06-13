package placarpartida

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/placares"
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
	return placares.TPSet
}

func runTest(score *Partida, results []placares.EstadoResultadoEParametroPlacar, SideA, SideB int, t *testing.T) {
	score.AdicionarEventoAoPontuarNaPartida(func(scoreA, scoreB int, done bool) {
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

func Test_Set_2x1(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 6, scoreB: 4}, score{scoreA: 4, scoreB: 6},
		score{scoreA: 7, scoreB: 5}, score{scoreA: 6, scoreB: 4},
	}

	match := New(Padrao())
	runTest(&match, scores, 2, 1, t)
}

func Test_Set_1x2(t *testing.T) {
	scores := []placares.EstadoResultadoEParametroPlacar{
		score{scoreA: 6, scoreB: 4}, score{scoreA: 3, scoreB: 6},
		score{scoreA: 5, scoreB: 7}, score{scoreA: 6, scoreB: 4},
	}

	match := New(Padrao())
	runTest(&match, scores, 1, 2, t)
}
