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

func WithAnotherTurn(t *turn.Turn) func(t *turn.Turn) {
	return func(t *turn.Turn) {
		turn.AddData(t, "timedTurning_start", turn.NewMapData(time.Now(), func() any { return time.Now() }))
		turn.AddData(t, "timedTurning_duration", turn.NewMapData(time.Since(time.Now()), func() any { return time.Since(time.Now()) }))

		t.AddOnAfterChange(func(st turning.TurningSide) {
			start, _ := turn.GetData[time.Time](t, "timedTurning_start")
			duration := time.Since(start)

			turn.UpdateData(t, "timedTurning_duration", duration)
		})
	}
}

func Duration(t *turn.Turn) time.Duration {
	duration, _ := turn.GetData[time.Duration](t, "timedTurning_duration")
	return duration
}
