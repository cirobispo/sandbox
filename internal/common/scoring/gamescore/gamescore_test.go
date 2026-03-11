package gamescore

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

func runTest(blocks []point.TestBlock, SideA, SideB int, t *testing.T) {
	score := New(turning.STA, false)

	score.AddOnAfterScoreEvent(func(scoreA, scoreB int, done bool) {
		a, b := scoring.Score2GameText(scoreA, scoreB)
		if done {
			t.Logf("Game FINAL status: ( %s x %s ) done: %v \n", a, b, done)
			return
		}

		t.Logf("Game status: ( %s x %s ) done: %v \n", a, b, done)
	})

	score.AddOnAfterScoreEvent(func(scoreA, scoreB int, done bool) {
		if !done && (scoreB >= 3 && scoreA < scoreB) {
			t.Logf("Break point: ( %v )\n", scoreB-scoreA)
		}
	})

	t.Log("Game status: ( 0 x 0 )")
	for i := range blocks {
		block := blocks[i]
		point := point.New(turn.New(turning.STA))
		for j := range block.Items {
			hit := block.Items[j].Value
			point.AddHit(hit)
		}
		score.AddScore(PointToScore(&point, false))
	}

	a, b := score.Result()
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
