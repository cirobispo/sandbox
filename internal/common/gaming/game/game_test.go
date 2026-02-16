package game

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/countingturn"
	"github.com/cirobispo/sandbox/internal/common/turning/timingturn"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

func runTest(blocks []point.TestBlock, SideA, SideB int, t *testing.T) {
	g := New(countingturn.NewFromTurn(timingturn.NewFromTurn(turn.New(turning.STA))), false)
	g.AddOnAddingPointEvent(func(scoreA, scoreB int, done bool) {
		tA, tB := scoring.Score2GameText(scoreA, scoreB)
		if done {
			t.Logf("Game FINAL status: ( %v x %v )\n", tA, tB)
			t.Logf("Duration: %v\n", timingturn.Duration(g.turn))
			t.Logf("Hits    : %v\n", countingturn.GetCount(g.turn))
			t.Log()
			return
		}

		t.Logf("Game status: ( %v x %v )\n", tA, tB)
		t.Logf("Duration: %v\n", timingturn.Duration(g.turn))
		t.Logf("Hits    : %v\n", countingturn.GetCount(g.turn))
		t.Log()
	})

	for i := range blocks {
		block := blocks[i]
		tn := turn.New(turning.STA)
		p := point.New(tn)

		for j := range block.Items {
			item := block.Items[j]
			t.Logf("%s hits %s, ", tn.LastSide().String(), item.Value.Type())
			p.AddHit(item.Value)
		}

		g.AddPoint(p)
	}

	a, b := g.Score().Result()
	if a != SideA || b != SideB {
		t.Errorf("\n\nGame should be (%d x %d) not (%d x %d)\n", SideA, SideB, a, b)
	}
}

func TestGameToServer_Game40(t *testing.T) {
	blocks := []point.TestBlock{point.AcePoint(), point.AcePoint(), point.WinnerSSPoint(),
		point.WinnerOSPoint(), point.WinnerOSPoint(), point.WinnerOSPoint(), point.WinnerOSPoint(),
		// point.DoubleFault(),
		point.LongRallieOSPoint(2, point.NetOppositeSide(true)),
		point.LongRallieOSPoint(2, point.NetOppositeSide(true)),
		// point.LongRallieOSPoint(2, point.NetSameSide(true)),
	}

	runTest(blocks, 4, 3, t)
}

/**
	func TestGameToServer_Game30(t *testing.T) {
		blocks := []point.TestBlock{point.AcePoint(), point.AcePoint(), point.WinnerSSPoint(),
			point.WinnerOSPoint(), point.WinnerOSPoint(), point.LongRallieOSPoint(point.NetOppositeSide(true)),
		}

		runTest(blocks, 4, 2, t)
	}

	func TestGameToServer_Game15(t *testing.T) {
		blocks := []point.TestBlock{point.AcePoint(), point.AcePoint(), point.WinnerSSPoint(),
			point.WinnerOSPoint(), point.LongRallieOSPoint(point.NetOppositeSide(true)),
		}

		runTest(blocks, 4, 1, t)
	}

	func TestGameToServer_Game00(t *testing.T) {
		blocks := []point.TestBlock{point.AcePoint(), point.AcePoint(), point.WinnerSSPoint(),
			point.LongRallieOSPoint(point.NetOppositeSide(true)),
		}

		runTest(blocks, 4, 0, t)
	}

func TestGameToBrake_4040(t *testing.T) {
	blocks := []point.TestBlock{point.DoubleFault(), point.DoubleFault(), point.WinnerSSPoint(),
		point.LongRallieOSPoint(point.NetOppositeSide(true)), point.WinnerOSPoint(), point.WinnerSSPoint(),
	}
	runTest(blocks, 3, 3, t)
}

/**
func TestGameToBrake_30Game(t *testing.T) {
	blocks := []point.TestBlock{point.DoubleFault(), point.DoubleFault(), point.WinnerSSPoint(),
		point.LongRallieOSPoint(point.NetOppositeSide(true)), point.WinnerOSPoint(), point.WinnerOSPoint(),
	}
	runTest(blocks, 2, 4, t)
}

func TestGameToBrake_15Game(t *testing.T) {
	blocks := []point.TestBlock{point.DoubleFault(), point.LongRallieOSPoint(point.NetOppositeSide(true)),
		point.WinnerOSPoint(), point.DoubleFault(), point.WinnerOSPoint(),
	}
	runTest(blocks, 1, 4, t)
}

func TestGameToBrake_00Game(t *testing.T) {
	blocks := []point.TestBlock{point.DoubleFault(), //point.LongRallieOSPoint(point.NetOppositeSide(true)),
		point.WinnerOSPoint(), point.DoubleFault(), point.WinnerOSPoint(),
	}
	runTest(blocks, 0, 4, t)
}
/**/
