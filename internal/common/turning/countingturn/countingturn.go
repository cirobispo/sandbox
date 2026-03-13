package countingturn

import (
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type CountingTurn interface {
	turning.Turning
	Count(t *turn.Turn) int
}

func WithAnotherTurn(t *turn.Turn) func(t *turn.Turn) {
	return func(t *turn.Turn) {
		turn.AddData(t, "CountedTurning_count", turn.NewMapData(0, func() any { return 0 }))

		t.AddOnAfterChange(func(st turning.TurningSide) {
			value, _ := turn.GetData[int](t, "CountedTurning_count")
			value++
			turn.UpdateData(t, "CountedTurning_count", value)
		})
	}
}

func Count(t *turn.Turn) int {
	result, _ := turn.GetData[int](t, "CountedTurning_count")
	return result

}
