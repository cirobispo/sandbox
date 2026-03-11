package setscore

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type score struct {
	scoreA, scoreB int
}

func (s score) Done() bool {
	return true
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
	return scoring.STSet
}

func runTest(custom bool, results []scoring.ScoringResulting, SideA, SideB int, t *testing.T) {
	score := New(WithDefaultSet(turning.STA))
	if custom {
		if SideB > SideA {
			score = New(WithSideSizeAndTieBreak(turning.STA, SideB, true, true))
		} else {
			score = New(WithSideSizeAndTieBreak(turning.STA, SideA, true, true))
		}
	}

	score.AddOnAfterScoreEvent(func(scoreA, scoreB int, isTieBreak, done bool) {
		if isTieBreak && !done {
			t.Logf("TieBreak (%d x %d)\n", scoreA, scoreB)
		}
	})

	score.AddOnAfterScoreEvent(func(scoreA, scoreB int, isTieBreak, done bool) {
		if done {
			t.Logf("Score (%d x %d)\n", scoreA, scoreB)
		}
	})

	t.Logf("Testing a result for a Set with %d games. Expected result (%d x %d)\n", len(results), SideA, SideB)
	for j := range results {
		item := results[j]
		score.AddScore(item)
	}

	sA, sB := score.Result()
	if sA != SideA || sB != SideB {
		for j := range results {
			scoreA, scoreB := scoring.Score2GameText(results[j].Result())
			t.Logf("Score (%v x %v)\n", scoreA, scoreB)
		}
		t.Errorf("\n\nSet should be (%d x %d) not (%d x %d)\n", SideA, SideB, sA, sB)
	}
}

func Test_Set_6x4(t *testing.T) {
	scores := []scoring.ScoringResulting{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 0, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 4, scoreB: 0},
	}

	runTest(false, scores, 6, 4, t)
}

func Test_Set_4x6(t *testing.T) {
	scores := []scoring.ScoringResulting{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 0, scoreB: 4},

		score{scoreA: 0, scoreB: 4}, score{scoreA: 1, scoreB: 4},
	}

	runTest(false, scores, 4, 6, t)
}

func Test_Set_6x0(t *testing.T) {
	scores := []scoring.ScoringResulting{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 4, scoreB: 1},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 4, scoreB: 0},

		score{scoreA: 5, scoreB: 3}, score{scoreA: 4, scoreB: 0},
	}

	runTest(false, scores, 6, 0, t)
}

func Test_Set_0x6(t *testing.T) {
	scores := []scoring.ScoringResulting{
		score{scoreB: 4, scoreA: 0}, score{scoreB: 4, scoreA: 1},

		score{scoreB: 4, scoreA: 2}, score{scoreB: 4, scoreA: 0},

		score{scoreB: 5, scoreA: 3}, score{scoreB: 4, scoreA: 0},
	}

	runTest(false, scores, 0, 6, t)
}

func Test_Set_7x6(t *testing.T) {
	scores := []scoring.ScoringResulting{
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

func Test_Set_6x7(t *testing.T) {
	scores := []scoring.ScoringResulting{
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

func Test_Set_7x5(t *testing.T) {
	scores := []scoring.ScoringResulting{
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

func Test_Set_4x3(t *testing.T) {
	scores := []scoring.ScoringResulting{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 0}, // score{scoreA: 0, scoreB: 4},
	}

	runTest(true, scores, 4, 3, t)
}

func Test_Set_3x4(t *testing.T) {
	scores := []scoring.ScoringResulting{
		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		score{scoreA: 4, scoreB: 2}, score{scoreA: 3, scoreB: 5},

		score{scoreA: 4, scoreB: 0}, score{scoreA: 1, scoreB: 4},

		// score{scoreA: 4, scoreB: 0},
		score{scoreA: 0, scoreB: 4},
	}

	runTest(true, scores, 3, 4, t)
}

func Test_Set_6x4_TooManyGames(t *testing.T) {
	scores := []scoring.ScoringResulting{
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
