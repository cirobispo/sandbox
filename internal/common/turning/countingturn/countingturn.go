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
	turn.AddData(result, "count", turn.NewMapData(0, func() any { return 0 }))

	result.AddOnAfterChange(func(st turning.SideTurn) {
		value, _ := turn.GetData[int](result, "count")
		value++
		turn.UpdateData(result, "count", value)
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
	result, _ := turn.GetData[int](t, "count")
	return result

}
