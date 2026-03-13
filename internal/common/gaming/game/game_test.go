package game

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

func runTest(personToServe turning.TurningSide, blocks []point.TestBlock, SideA, SideB int, t *testing.T) {
	g := New(turn.New(personToServe), false)
	g.AddOnAddingPointEvent(func(scoreA, scoreB int, done bool) {
		tA, tB := scoring.Score2GameText(scoreA, scoreB)
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
		tn := turn.New(turning.TSA)
		p := point.New(tn)

		for j := range block.Items {
			item := block.Items[j]
			t.Logf("%s hits %s, ", tn.CurrentSide().String(), item.Value.Type())
			p.AddHit(item.Value)
		}

		g.AddPoint(p)
	}

	a, b := g.Score().Result()

	if a != SideA || b != SideB {
		t.Errorf("\n\nGame should be (%d x %d) not (%d x %d)\n", SideA, SideB, a, b)
	}
}

func TestTurnA_Game40(t *testing.T) {
	blocks := []point.TestBlock{point.AcePoint(), point.AcePoint(), point.WinnerSSPoint(),
		point.WinnerOSPoint(), point.WinnerOSPoint(), point.WinnerOSPoint(), point.WinnerOSPoint(),
		// point.DoubleFault(),
		point.LongRallieOSPoint(2, point.NetOppositeSide(true)),
		point.LongRallieOSPoint(2, point.NetOppositeSide(true)),
		// point.LongRallieOSPoint(2, point.NetSameSide(true)),
		point.AcePoint(),
	}

	runTest(turning.TSA, blocks, 5, 3, t)
}

func TestTurnB_40Game(t *testing.T) {
	blocks := []point.TestBlock{point.AcePoint(), point.AcePoint(), point.WinnerSSPoint(),
		point.WinnerOSPoint(), point.WinnerOSPoint(), point.WinnerOSPoint(), point.WinnerOSPoint(),
		// point.DoubleFault(),
		point.LongRallieOSPoint(2, point.NetOppositeSide(true)),
		point.LongRallieOSPoint(2, point.NetOppositeSide(true)),
		// point.LongRallieOSPoint(2, point.NetSameSide(true)),
		point.AcePoint(),
	}

	runTest(turning.TSB, blocks, 3, 5, t)
}
