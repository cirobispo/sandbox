package game_test

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/gaming/game"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

func runTest(blocks []point.TestBlock, SideA, SideB int, t *testing.T) {
	g := game.New(turn.New(turning.STA), true)
	g.AddOnAddingPointEvent(func(scoreA, scoreB int, done bool) {
		if done {
			t.Logf("Game FINAL status: ( %d x %d )\n", scoreA, scoreB)
		}
	})

	for i := range blocks {
		block := blocks[i]
		p := point.New(turn.New(turning.STA))
		a, b := g.Score().Result()
		t.Logf("Game status: ( %d x %d )\n", a, b)

		for j := range block.Items {
			item := block.Items[j]
			t.Logf("side %s hits %s, ", p.BallLastSide(), item.Value.Type())
			p.AddHit(item.Value)
		}
		t.Log()

		g.AddPoint(p)
	}

	a, b := g.Score().Result()
	if a != SideA || b != SideB {
		t.Errorf("\n\nGame should be (%d x %d) not (%d x %d)\n", SideA, SideB, a, b)
	}
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

/**/

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
