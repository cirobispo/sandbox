package set

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type testItem struct {
	servingSide turning.TurningSide
	score       scoring.ScoringResulting
	points      []point.Point
}

func (t testItem) ServingSide() turning.TurningSide {
	return t.servingSide
}

func (t testItem) Score() scoring.ScoringResulting {
	return t.score
}

func (t testItem) Points() []point.Point {
	return t.points
}

type score struct {
	servingSide    turning.TurningSide
	decidingPoint  bool
	scoreA, scoreB int
}

func (s score) Done() bool {
	sA, sB := s.Result()
	diff := 2
	if s.decidingPoint {
		diff--
	}

	AWins := sA >= 4 && (sA-sB) >= diff
	BWins := sB >= 4 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}

func (s score) Result() (int, int) {
	return s.scoreA, s.scoreB
}

func (s score) Side() scoring.ScoringSide {
	if s.scoreB > s.scoreA {
		return scoring.SSOpposite
	}

	return scoring.SSServing
}

func (s score) Type() scoring.ScoringType {
	return scoring.STGame
}

func (i *testItem) setServingSide(side turning.TurningSide) {
	a, b := i.score.Result()
	i.servingSide = side
	i.score = newScore(side, a, b)
}

func newItem(scoreA, scoreB int) testItem {
	result := testItem{score: newScore(turning.TSA, scoreA, scoreB), points: make([]point.Point, 0)}
	return result
}

func newScore(servingSide turning.TurningSide, scoreA, scoreB int) score {
	return score{
		servingSide: servingSide,
		scoreA:      scoreA,
		scoreB:      scoreB,
	}
}

func runTest(blocks []testItem, SideA, SideB int, t *testing.T) {
	myTurn := turn.New(turn.WithTurningSide(turning.TSA))
	mySet := New(WithDefaultSet(myTurn))

	sideToServe := myTurn.CurrentSide()
	mySet.AddOnAddingGameEvent(func(scoreA, scoreB int, done bool) {
		if done {
			t.Log("FINAL ")
		}

		t.Logf("%s -> Score (%d x %d)\n", sideToServe, scoreA, scoreB)
	})

	mySet.AddOnPlayerChangeEvent(func() {
		t.Logf("Players change side\n")
	})

	for j := range blocks {
		currentGame := mySet.NewGame()
		item := blocks[j]
		item.setServingSide(currentGame.ServingSide())
		sideToServe = item.servingSide
		if err := mySet.AddGame(item); err != nil {
			t.Errorf("Error ao adicionar %v. Mensagem: %s", item, err)
		}
	}
	sA, sB := mySet.Score().Result()
	if sA != SideA || sB != SideB {
		for i := range blocks {
			item := blocks[i]
			a, b := item.Score().Result()
			t.Logf("Score game #%d (%d x %d)\n", i+1, a, b)
		}
		t.Errorf("\n\nSet should be (%d x %d) not (%d x %d)\n", SideA, SideB, sA, sB)
	}
}

func Test_6x4(t *testing.T) {
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

func Test_4x6(t *testing.T) {
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
