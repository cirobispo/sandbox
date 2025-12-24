package countingturn

import (
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type CountedTurning interface {
	turning.Turning
	Counter(t turn.Turn) int
}

func newFromTurn(t *turn.Turn) *turn.Turn {
	result := t

	result.AddAfterChangeEvent(func(st turning.SideTurn) {
		value := turn.GetData[int](result, "count") + 1
		turn.AddData[int](result, "count", value)
	})

	return result
}

func New(start turning.SideTurn) *turn.Turn {
	t := turn.New(start)
	return newFromTurn(t)
}

func NewFromTurn(t *turn.Turn) *turn.Turn {
	return newFromTurn(t)
}

func GetCount(t *turn.Turn) int {
	return turn.GetData[int](t, "count")
}
