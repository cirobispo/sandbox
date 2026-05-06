package match

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type testItem struct {
	score placares.EstadoEResultadoPlacar
	sets  []placares.EstadoResultadoEParametroPlacar
}

func (t testItem) Score() placares.EstadoEResultadoPlacar {
	return t.score
}

type score struct {
	sideToBegin    turning.TurningSide
	bestOf         int
	scoreA, scoreB int
}

func (m score) Terminado() bool {
	sA, sB := m.Resultado()

	amountToWin := (m.bestOf / 2) + (m.bestOf % 2)

	sideAWins := (sA >= amountToWin && sA-sB >= 1)
	sideBWins := (sB >= amountToWin && sB-sA >= 1)
	result := sideAWins || sideBWins

	return result
}

func (m score) Resultado() (int, int) {
	return m.scoreA, m.scoreB
}

func (m score) Lado() placares.LadoDoPlacar {
	if m.scoreB > m.scoreA {
		return placares.LPOposto
	}

	return placares.LPServico
}

func (m score) Tipo() placares.TipoDoPlacar {
	return placares.TPJogo
}

func newItem(scoreA, scoreB int) testItem {
	result := testItem{score: newScore(turning.TSA, scoreA, scoreB), sets: make([]placares.EstadoResultadoEParametroPlacar, 0)}
	return result
}

func newScore(servingSide turning.TurningSide, scoreA, scoreB int) score {
	return score{
		sideToBegin: servingSide,
		scoreA:      scoreA,
		scoreB:      scoreB,
	}
}

func runTest(blocks []testItem, SideA, SideB int, t *testing.T) {
	myTurn := turn.New(turn.WithTurningSide(turning.TSA))
	myMatch := New(DefaultMatch())

	sideToServe := myTurn.CurrentSide()
	myMatch.AddOnAddingSetEvent(func(scoreA, scoreB int, done bool) {
		if done {
			t.Log("FINAL ")
		}

		t.Logf("%s -> Score (%d x %d)\n", sideToServe, scoreA, scoreB)
	})

	// for j := range blocks {
	// 	currentSet := myMatch.NewSet()
	// 	item := blocks[j]
	// 	item.setServingSide(currentSet.ServingSide())
	// 	sideToServe = item.servingSide
	// 	if err := myMatch.AddGame(item); err != nil {
	// 		t.Errorf("Error ao adicionar %v. Mensagem: %s", item, err)
	// 	}
	// }
	// sA, sB := myMatch.Score().Result()
	// if sA != SideA || sB != SideB {
	// 	for i := range blocks {
	// 		item := blocks[i]
	// 		a, b := item.Score().Result()
	// 		t.Logf("Score game #%d (%d x %d)\n", i+1, a, b)
	// 	}
	// 	t.Errorf("\n\nSet should be (%d x %d) not (%d x %d)\n", SideA, SideB, sA, sB)
	// }
}

func Test_2x1(t *testing.T) {
	data := []testItem{
		newItem(6, 4), newItem(2, 6),
		newItem(6, 3), newItem(1, 6),
		newItem(6, 4), newItem(4, 6),
		newItem(6, 4), newItem(4, 6),
		newItem(6, 4), newItem(6, 4),
		// newItem(6, 4), newItem(6, 4),
	}
	runTest(data, 6, 4, t)
}

func Test_1x2(t *testing.T) {
	data := []testItem{
		newItem(6, 4), newItem(2, 6),
		newItem(6, 3), newItem(1, 6),
		newItem(6, 4), newItem(4, 6),
		newItem(6, 4), newItem(4, 6),
		newItem(5, 7), newItem(6, 6),
		// newItem(6, 4), newItem(6, 4),
	}
	runTest(data, 4, 6, t)
}
