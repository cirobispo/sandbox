package matchscore

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/scoring"
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
		return scoring.SSB
	}

	return scoring.SSA
}

func (s score) Type() scoring.ScoringType {
	return scoring.STMatch
}

func runTest(score *MatchScore, results []scoring.Scoring, SideA, SideB int, t *testing.T) {
	score.AddOnAfterScoreEvent(func(scoreA, scoreB int, done bool) {
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

func Test_Set_2x1(t *testing.T) {
	scores := []scoring.Scoring{
		score{scoreA: 6, scoreB: 4}, score{scoreA: 4, scoreB: 6},
		score{scoreA: 7, scoreB: 5}, score{scoreA: 6, scoreB: 4},
	}

	match := New(WithDefault())
	runTest(&match, scores, 2, 1, t)
}

func Test_Set_1x2(t *testing.T) {
	scores := []scoring.Scoring{
		score{scoreA: 6, scoreB: 4}, score{scoreA: 3, scoreB: 6},
		score{scoreA: 5, scoreB: 7}, score{scoreA: 6, scoreB: 4},
	}

	match := New(WithDefault())
	runTest(&match, scores, 1, 2, t)
}
