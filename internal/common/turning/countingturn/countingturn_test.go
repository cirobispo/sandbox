package countingturn

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type testItem struct {
	turns     int
	startSide turning.TurningSide
}

func runTest(test testItem, t *testing.T) {
	obj := turn.New(WithAnotherTurn(turn.New(turn.WithTurningSide(test.startSide))))

	for a := test.turns; a > 0; a-- {
		obj.Execute()
	}

	if value := Count(obj); value != test.turns {
		t.Errorf("turn is \"%v\" turns and should be \"%v\"", value, test.turns)
	}
}

func Test10CrancksSideA(t *testing.T) {
	runTest(testItem{10, turning.TSA}, t)
}

func Test10CrancksSideB(t *testing.T) {
	runTest(testItem{10, turning.TSB}, t)
}
