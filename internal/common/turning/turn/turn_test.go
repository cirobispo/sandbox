package turn

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/turning"
)

type testItem struct {
	turns     int
	startSide turning.TurningSide
}

func inverse(s turning.TurningSide) turning.TurningSide {
	if s == turning.TSA {
		return turning.TSB
	}
	return turning.TSA
}

func runTest(test testItem, t *testing.T) {
	obj := New(WithTurningSide(test.startSide))
	endSide := test.startSide
	if test.turns%2 != 0 {
		endSide = inverse(endSide)
	}

	for test.turns > 0 {
		obj.Execute()
		test.turns--
	}

	if obj.CurrentSide() != endSide {
		t.Errorf("turn is \"%s\" and should be \"%s\"", obj.CurrentSide(), endSide)
	}
}

func Test10Crancks(t *testing.T) {
	runTest(testItem{10, turning.TSB}, t)
}
