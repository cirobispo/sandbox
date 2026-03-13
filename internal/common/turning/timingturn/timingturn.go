package timingturn

import (
	"time"

	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type TimedTurning interface {
	turning.Turning
	Duration(t *turn.Turn) int
}

func newFromTurn(t *turn.Turn) *turn.Turn {
	turn.AddData(t, "timedTurning_start", turn.NewMapData(time.Now(), func() any { return time.Now() }))
	turn.AddData(t, "timedTurning_duration", turn.NewMapData(time.Since(time.Now()), func() any { return time.Since(time.Now()) }))
	result := t

	result.AddOnAfterChange(func(st turning.TurningSide) {
		start, _ := turn.GetData[time.Time](t, "timedTurning_start")
		duration := time.Since(start)

		turn.UpdateData(t, "timedTurning_duration", duration)
	})

	return result
}

func New(start turning.TurningSide) *turn.Turn {
	t := turn.New(start)
	return NewFromTurn(t)
}

func NewFromTurn(t *turn.Turn) *turn.Turn {
	return newFromTurn(t)
}

func Duration(t *turn.Turn) time.Duration {
	duration, _ := turn.GetData[time.Duration](t, "timedTurning_duration")
	return duration
}
