package countingturn

import (
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type CountingTurn interface {
	turning.Turning
	Count(t *turn.Turn) int
}

func newFromTurn(t *turn.Turn) *turn.Turn {
	result := t
	turn.AddData(result, "CountedTurning_count", turn.NewMapData(0, func() any { return 0 }))

	result.AddOnAfterChange(func(st turning.TurningSide) {
		value, _ := turn.GetData[int](result, "CountedTurning_count")
		value++
		turn.UpdateData(result, "CountedTurning_count", value)
	})

	return result
}

func New(start turning.TurningSide) *turn.Turn {
	t := turn.New(start)
	return newFromTurn(t)
}

func NewFromTurn(t *turn.Turn) *turn.Turn {
	return newFromTurn(t)
}

func Count(t *turn.Turn) int {
	result, _ := turn.GetData[int](t, "CountedTurning_count")
	return result

}
