package timingturn

import (
	"testing"
	"time"

	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type testItem struct {
	turns     int
	startSide turning.TurningSide
}

func runTest(test testItem, t *testing.T) {
	obj := turn.New(WithAnotherTurn(turn.New(turn.WithTurningSide(test.startSide))))
	begin := time.Now()
	for a := test.turns; a > 0; a-- {
		obj.Execute()
	}
	diff := time.Since(begin).Round(time.Millisecond)

	if value := Duration(obj).Round(time.Millisecond); value != diff {
		t.Errorf("turn is \"%v\" turns and should be \"%v\"", value, diff)
	}
}

func Test10CrancksSideA(t *testing.T) {
	runTest(testItem{100, turning.TSA}, t)
}

func Test10CrancksSideB(t *testing.T) {
	runTest(testItem{100, turning.TSB}, t)
}
